package tournamentapp

type TournamentConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BirthYears  []int  `json:"birthYears"`

	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`

	Travel TravelConfig `json:"travel"`

	Payments PaymentConfig `json:"payments"`
}

type TravelConfig struct {
	DepartureDate    string `json:"departureDate"`
	ArrivalDate      string `json:"arrivalDate"`
	ArrivalAirport   string `json:"arrivalAirport"`
	ReturnDate       string `json:"returnDate"`
	DepartureAirport string `json:"departureAirport"`
}

type PaymentConfig struct {
	Payment1Date string `json:"payment1Date"`
	Payment2Date string `json:"payment2Date"`
	Payment3Date string `json:"payment3Date"`
}
