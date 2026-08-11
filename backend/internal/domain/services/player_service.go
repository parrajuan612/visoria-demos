package services

import (
	"time"

	tournamentapp "github.com/juanparra/visoria-demo/internal/application/tournament"
	"github.com/juanparra/visoria-demo/internal/domain"
	"github.com/juanparra/visoria-demo/internal/domain/validation"
)

type PlayerService struct {
	validator *validation.PlayerValidator
}

func NewPlayerService(
	validator *validation.PlayerValidator,
) *PlayerService {
	return &PlayerService{
		validator: validator,
	}
}

func (s *PlayerService) PreparePlayer(
	player domain.Player,
	config tournamentapp.TournamentConfig,
) domain.Player {

	startDate, _ := time.Parse("2006-01-02", config.StartDate)
	endDate, _ := time.Parse("2006-01-02", config.EndDate)

	tournament := domain.Tournament{
		Name:         config.Name,
		Description:  config.Description,
		StartDate:    startDate,
		EndDate:      endDate,
		CategoryText: make(map[string]string),
	}

	tournament.Payments.Plans = make(map[int]domain.PaymentPlan)

	for _, category := range config.Categories {
		tournament.CategoryText[category.Name] = category.Description

		for _, year := range category.BirthYears {
			if player.BirthDate.Year() == year {
				player.Category = category.Name
			}
		}
	}

	for _, scholarship := range config.Scholarships {
		total := int(scholarship.Total)

		payment1 := total / 3
		payment2 := total / 3
		payment3 := total - payment1 - payment2

		tournament.Payments.Plans[scholarship.Percentage] = domain.PaymentPlan{
			Total:    total,
			Payment1: payment1,
			Payment2: payment2,
			Payment3: payment3,
		}
	}

	player.Tournament = &tournament

	plan, exists := tournament.Payments.Plans[player.Scholarship]
	if exists {
		player.PaymentPlan = &plan
	}

	errors, warnings := s.validator.Validate(player)

	player.Errors = errors
	player.Warnings = warnings

	switch {
	case len(errors) > 0:
		player.Status = "ERROR"
	case len(warnings) > 0:
		player.Status = "WARNING"
	default:
		player.Status = "VALID"
	}

	return player
}
