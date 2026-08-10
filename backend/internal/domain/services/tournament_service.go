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
