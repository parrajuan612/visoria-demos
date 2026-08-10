package tournament

import "time"

type Config struct {
	Name string

	DocumentDate time.Time

	DepartureDate time.Time
	ArrivalDate   time.Time
	ReturnDate    time.Time

	Payment1Amount float64
	Payment1Date   time.Time

	Payment2Amount float64
	Payment2Date   time.Time

	Payment3Amount float64
	Payment3Date   time.Time
}
