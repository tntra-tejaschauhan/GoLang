package models

import (
	"sync"
)

type SeatStatus int

const (
	Available SeatStatus = iota
	Held
	Booked
)

type Flight struct {
	ID          string
	Name        string
	Seats       int
	Seat_metrix map[string]*Seat
	Price       float64
	mu          sync.Mutex //
}

type Seat struct {
	Number string
	Status SeatStatus
	mu     sync.Mutex
}

func (f *Flight) HoldSeat(seatNo string) bool {
	seat, exists := f.Seat_metrix[seatNo]
	if !exists {
		return false
	}
	seat.mu.Lock()
	defer seat.mu.Unlock()

	if seat.Status != Available {
		return false

	}
	seat.Status = Held

	return true
}

func (f *Flight) ConfirmSeat(seatNo string) bool {
	seat, exists := f.Seat_metrix[seatNo]
	if !exists {
		return false
	}
	seat.mu.Lock()
	defer seat.mu.Unlock()

	if seat.Status != Held {
		return false
	}
	seat.Status = Booked
	return true
}

func (f *Flight) ReleaseSeat(seatNo string) {
	seat, exists := f.Seat_metrix[seatNo]
	if !exists {
		return
	}

	seat.mu.Lock()
	defer seat.mu.Unlock()

	if seat.Status == Held {
		seat.Status = Available
	}
}
