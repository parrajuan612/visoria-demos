package domain

import "time"

type Tournament struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	StartDate    time.Time         `json:"startDate"`
	EndDate      time.Time         `json:"endDate"`
	Travel       TravelConfig      `json:"travel"`
	Payments     PaymentConfig     `json:"payments"`
	CategoryText map[string]string `json:"categoryText"`
}

type TravelConfig struct {
	DepartureDate    time.Time `json:"departureDate"`
	ArrivalDate      time.Time `json:"arrivalDate"`
	ArrivalAirport   string    `json:"arrivalAirport"`
	ReturnDate       time.Time `json:"returnDate"`
	DepartureAirport string    `json:"departureAirport"`
}

type PaymentConfig struct {
	Payment1Date time.Time `json:"payment1Date"`
	Payment2Date time.Time `json:"payment2Date"`
	Payment3Date time.Time `json:"payment3Date"`

	Plans map[int]PaymentPlan `json:"plans"`
}

type PaymentPlan struct {
	Total    int `json:"total"`
	Payment1 int `json:"payment1"`
	Payment2 int `json:"payment2"`
	Payment3 int `json:"payment3"`
}
