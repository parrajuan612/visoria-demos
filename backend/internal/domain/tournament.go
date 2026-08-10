package domain

import "time"

type Tournament struct {
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
}
