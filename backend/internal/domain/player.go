package domain

import "time"

type Player struct {
	Name           string       `json:"name"`
	Club           string       `json:"club"`
	Scholarship    int          `json:"scholarship"`
	BirthDate      time.Time    `json:"birthDate"`
	GuardianName   string       `json:"guardianName"`
	PrimaryPhone   string       `json:"primaryPhone"`
	SecondaryPhone string       `json:"secondaryPhone"`
	Category       string       `json:"category"`
	Tournaments    []Tournament `json:"tournaments,omitempty"` // <--- CAMBIO AQUÍ (Array de torneos)
	PaymentPlan    *PaymentPlan `json:"paymentPlan,omitempty"`
	Status         string       `json:"status"`
	Errors         []string     `json:"errors"`
	Warnings       []string     `json:"warnings"`
}
