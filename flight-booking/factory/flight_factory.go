package factory

import (
	"flight-booking/models"
	"fmt"
)

func CreateFlight(flightType string, seats int) *models.Flight {
	switch flightType {
	case "domestic":

		d_flight := &models.Flight{
			ID:          "D101",
			Name:        "IndiGo",
			Seats:       seats,
			Seat_metrix: make(map[string]*models.Seat),
			Price:       5000,
		}
		for i := 1; i <= seats; i++ {
			seatNumber := fmt.Sprintf("A%d", i)
			d_flight.Seat_metrix[seatNumber] = &models.Seat{
				Number: seatNumber,
				Status: models.Available,
			}
		}
		return d_flight

	case "international":

		i_flight := &models.Flight{
			ID:          "I201",
			Name:        "Air India",
			Seats:       seats,
			Seat_metrix: make(map[string]*models.Seat),
			Price:       25000,
		}
		for i := 1; i <= seats; i++ {
			seatNumber := fmt.Sprintf("A%d", i)
			i_flight.Seat_metrix[seatNumber] = &models.Seat{
				Number: seatNumber,
				Status: models.Available,
			}
		}
		return i_flight

	default:
		return nil
	}
}
