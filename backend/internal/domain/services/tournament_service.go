package services

import (
	"time"

	"github.com/juanparra/visoria-demo/internal/domain"
)

type TournamentService struct {
	tournaments map[string]domain.Tournament
}

func NewTournamentService() *TournamentService {
	return &TournamentService{
		tournaments: map[string]domain.Tournament{
			"mic": {
				Name:        "MIC FOOTBALL COSTA BRAVA BARCELONA ESPAÑA",
				Description: "Torneo MIC Football Costa Brava Barcelona España",
			},

			"easter": {
				Name:        "IDA EASTER CUP COMUNIDAD VALENCIANA",
				Description: "Torneo IDA Easter Cup Comunidad Valenciana",
			},

			"mic_easter": {
				Name:        "MIC FOOTBALL / IDA EASTER CUP 2027",
				Description: "Programa de intercambio deportivo",

				// 1. Integración de fechas generales del torneo
				StartDate: parseDate("18/03/2027"),
				EndDate:   parseDate("28/03/2027"),

				Travel: domain.TravelConfig{
					DepartureDate:    parseDate("18/03/2027"),
					ArrivalDate:      parseDate("19/03/2027"),
					ArrivalAirport:   "Barcelona",
					ReturnDate:       parseDate("28/03/2027"),
					DepartureAirport: "Barcelona",
				},

				Payments: domain.PaymentConfig{
					Payment1Date: parseDate("10/01/2027"),
					Payment2Date: parseDate("10/02/2027"),
					Payment3Date: parseDate("10/03/2027"),

					Plans: map[int]domain.PaymentPlan{
						0: {
							Total:    2800,
							Payment1: 930,
							Payment2: 930,
							Payment3: 940,
						},
						30: {
							Total:    1960,
							Payment1: 650,
							Payment2: 650,
							Payment3: 660,
						},
						50: {
							Total:    1550,
							Payment1: 500,
							Payment2: 500,
							Payment3: 550,
						},
						70: {
							Total:    990,
							Payment1: 495,
							Payment2: 495,
							Payment3: 0,
						},
						100: {
							Total:    200,
							Payment1: 200,
							Payment2: 0,
							Payment3: 0,
						},
					},
				},

				CategoryText: map[string]string{
					"JUVENIL":  "JUVENILES nacidos en el año 2008-2009-2010",
					"CADETE":   "CADETES nacidos en el año 2011-2012",
					"INFANTIL": "INFANTIL nacidos en el año 2013-2014",
					"ALEVÍN":   "ALEVÍN nacidos en el año 2015-2016",
					"BENJAMÍN": "BENJAMÍN nacidos en el año 2017-2018",
				},
			},

			"villarreal": {
				Name:        "VILLARREAL YELLOW CUP",
				Description: "Torneo Villarreal Yellow Cup",
			},

			"custom": {
				Name:        "TORNEO PERSONALIZADO",
				Description: "Programa de intercambio deportivo",
			},
		},
	}
}

func (s *TournamentService) Get(key string) (domain.Tournament, bool) {
	tournament, exists := s.tournaments[key]

	if !exists {
		return domain.Tournament{}, false
	}

	return tournament, true
}

func (s *TournamentService) SetDates(
	key string,
	startDate time.Time,
	endDate time.Time,
) bool {
	tournament, exists := s.tournaments[key]

	if !exists {
		return false
	}

	tournament.StartDate = startDate
	tournament.EndDate = endDate

	s.tournaments[key] = tournament

	return true
}

func parseDate(value string) time.Time {
	date, _ := time.Parse("02/01/2006", value)
	return date
}
