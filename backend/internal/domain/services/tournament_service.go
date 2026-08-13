package services

import (
	"time"

	tournamentapp "github.com/juanparra/visoria-demo/internal/application/tournament"
	"github.com/juanparra/visoria-demo/internal/domain"
	store "github.com/juanparra/visoria-demo/internal/shared"
)

type TournamentService struct{}

func NewTournamentService() *TournamentService {
	return &TournamentService{}
}

// Actualizamos la firma para que retorne un []domain.Tournament
func (s *TournamentService) GetByBirthDate(
	birthDate time.Time,
) ([]domain.Tournament, bool) {

	configs := store.GetTournamentConfigs()
	birthYear := birthDate.Year()

	// Creamos un slice para almacenar todos los torneos que coincidan
	var matchedTournaments []domain.Tournament

	for _, config := range configs {
		for _, year := range config.BirthYears {
			if year == birthYear {
				// Si coincide, lo agregamos al slice
				matchedTournaments = append(matchedTournaments, configToDomain(config))
				// Rompemos el ciclo interno de años para no agregar el mismo torneo dos veces
				// en caso de que el JSON tuviera un año duplicado por error.
				break
			}
		}
	}

	// Si encontramos al menos un torneo, retornamos el slice completo y true
	if len(matchedTournaments) > 0 {
		return matchedTournaments, true
	}

	// Si no hubo coincidencias, retornamos nil y false
	return nil, false
}

func configToDomain(
	config tournamentapp.TournamentConfig,
) domain.Tournament {

	startDate, _ := time.Parse("2006-01-02", config.StartDate)
	endDate, _ := time.Parse("2006-01-02", config.EndDate)

	departureDate, _ := time.Parse(
		"2006-01-02",
		config.Travel.DepartureDate,
	)

	arrivalDate, _ := time.Parse(
		"2006-01-02",
		config.Travel.ArrivalDate,
	)

	returnDate, _ := time.Parse(
		"2006-01-02",
		config.Travel.ReturnDate,
	)

	payment1Date, _ := time.Parse(
		"2006-01-02",
		config.Payments.Payment1Date,
	)

	payment2Date, _ := time.Parse(
		"2006-01-02",
		config.Payments.Payment2Date,
	)

	payment3Date, _ := time.Parse(
		"2006-01-02",
		config.Payments.Payment3Date,
	)

	return domain.Tournament{
		Name:        config.Name,
		Description: config.Description,
		StartDate:   startDate,
		EndDate:     endDate,

		Travel: domain.TravelConfig{
			DepartureDate:    departureDate,
			ArrivalDate:      arrivalDate,
			ArrivalAirport:   config.Travel.ArrivalAirport,
			ReturnDate:       returnDate,
			DepartureAirport: config.Travel.DepartureAirport,
		},

		Payments: domain.PaymentConfig{
			Payment1Date: payment1Date,
			Payment2Date: payment2Date,
			Payment3Date: payment3Date,
		},
	}
}
