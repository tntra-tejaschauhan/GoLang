package service

import (
	"flight-booking/models"
	"flight-booking/strategy"
	"fmt"
	"sync"
	"time"
)

type BookingRequest struct {
	Booking models.Booking
	Payment strategy.PaymentStrategy
}

func StartWorkerPool(requests chan BookingRequest, workerCount int, wg *sync.WaitGroup) {

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, requests, wg)
	}
}

func worker(id int, requests chan BookingRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	for req := range requests {
		start := time.Now() // ⏱️ START TIMER
		status := ""
		fmt.Println("worker", id, "processing booking for", req.Booking.UserName)
		seatNo := req.Booking.SeatNo

		// 1️⃣ Hold Seat
		if !req.Booking.Flight.HoldSeat(seatNo) {

			status = "failed"
			duration := time.Since(start)
			GlobalMetrics.Record(duration, status)

			fmt.Println("seat not available for", req.Booking.UserName)
			continue
		}

		// 2️⃣ Create channel for payment result
		paymentDone := make(chan bool)

		go func() {
			success := req.Payment.Pay(req.Booking.Flight.Price)
			paymentDone <- success
		}()

		// 3️⃣ Wait for payment OR timeout
		select {
		// after succesful payment the seat will confirm, otherwise failed
		case success := <-paymentDone:
			if success {
				status = "success"
				req.Booking.Flight.ConfirmSeat(seatNo)
				fmt.Println("Ticket confirmed for", req.Booking.UserName)
			} else {
				status = "failed"
				req.Booking.Flight.ReleaseSeat(seatNo)
				fmt.Println("Payment failed for", req.Booking.UserName)
			}
			// it start counting immediately when the select statement executed.

			// So timeline:

			// select starts running

			// Two things are now being waited on:

			// paymentDone channel
			// 8-second timer

			// Whichever happens first wins:

			// Payment finishes → first case runs
			// 8 seconds pass → timeout case runs

		case <-time.After(8 * time.Second):

			req.Booking.Flight.ReleaseSeat(seatNo)
			status = "timeout"
			fmt.Println("Payment timeout! Booking failed for", req.Booking.UserName)
		}
		duration := time.Since(start)
		// 📊 Record metrics
		GlobalMetrics.Record(duration, status)
	}

}
