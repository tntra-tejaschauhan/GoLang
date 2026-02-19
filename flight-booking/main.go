package main

import (
	"encoding/json"
	"flight-booking/factory"
	"flight-booking/models"
	"flight-booking/service"
	"flight-booking/strategy"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {

	flight := factory.CreateFlight("domestic", 100)
	go startMetricsServer()
	requests := make(chan service.BookingRequest, 100)

	var wg sync.WaitGroup

	// Start workers
	service.StartWorkerPool(requests, 100, &wg)

	// send booking request
	for i := 0; i < 20000; i++ {

		booking := models.Booking{
			UserName: fmt.Sprintf("User-%d", i),
			Flight:   flight,
			SeatNo:   fmt.Sprintf("A%d", rand.Intn(100)+1),
		}

		requests <- service.BookingRequest{
			Booking: booking,
			Payment: &strategy.UPI{},
		}
	}

	close(requests)
	// WAIT here
	wg.Wait()
	time.Sleep(100 * time.Second)
}
func startMetricsServer() {

	mux := http.NewServeMux()

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {

		data := map[string]interface{}{
			"total_requests": service.GlobalMetrics.TotalRequests,
			"success":        service.GlobalMetrics.Success,
			"failed":         service.GlobalMetrics.Failed,
			"timeout":        service.GlobalMetrics.Timeout,
			"avg_latency":    service.GlobalMetrics.AverageLatency().String(),
		}

		json.NewEncoder(w).Encode(data)
	})

	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {

		html := `
		<html>
		<head>
		<title>Booking Metrics Dashboard</title>
		<meta http-equiv="refresh" content="2">
		</head>
		<body>
		<h1>Flight Booking Metrics</h1>
		<p>Total Requests: %d</p>
		<p>Success: %d</p>
		<p>Failed due to Held by other user: %d</p>
		<p>failed due to payment Timeout: %d</p>
		<p>Average Latency: %s</p>
		</body>
		</html>`

		fmt.Fprintf(w, html,
			service.GlobalMetrics.TotalRequests,
			service.GlobalMetrics.Success,
			service.GlobalMetrics.Failed,
			service.GlobalMetrics.Timeout,
			service.GlobalMetrics.AverageLatency(),
		)
	})

	http.ListenAndServe(":8080", mux)
}
