package models

import "time"

type Reservation struct {
	BookId      int
	MemberId    int
	BookingTime time.Time
}
