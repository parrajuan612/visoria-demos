package tournamentapp

type TournamentConfig struct {
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	StartDate    string              `json:"startDate"`
	EndDate      string              `json:"endDate"`
	Travel       TravelConfig        `json:"travel"`
	Categories   []CategoryConfig    `json:"categories"`
	Scholarships []ScholarshipConfig `json:"scholarships"`
	Payments     PaymentConfig       `json:"payments"`
}

type TravelConfig struct {
	DepartureDate    string `json:"departureDate"`
	ArrivalDate      string `json:"arrivalDate"`
	ArrivalAirport   string `json:"arrivalAirport"`
	ReturnDate       string `json:"returnDate"`
	DepartureAirport string `json:"departureAirport"`
}

type CategoryConfig struct {
	Name        string `json:"name"`
	BirthYears  []int  `json:"birthYears"`
	Description string `json:"description"`
}

type ScholarshipConfig struct {
	Percentage int     `json:"percentage"`
	Total      float64 `json:"total"`
}

type PaymentConfig struct {
	Payment1Date string `json:"payment1Date"`
	Payment2Date string `json:"payment2Date"`
	Payment3Date string `json:"payment3Date"`
}
