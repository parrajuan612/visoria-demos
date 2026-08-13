package services

import (
	"github.com/juanparra/visoria-demo/internal/domain"
	domainPayment "github.com/juanparra/visoria-demo/internal/domain/payment"
	"github.com/juanparra/visoria-demo/internal/domain/validation"
)

type PlayerService struct {
	tournamentService *TournamentService
	validator         *validation.PlayerValidator
	categoryService   *CategoryService
}

func NewPlayerService(
	tournamentService *TournamentService,
	validator *validation.PlayerValidator,
	categoryService *CategoryService,
) *PlayerService {
	return &PlayerService{
		tournamentService: tournamentService,
		validator:         validator,
		categoryService:   categoryService,
	}
}

func (s *PlayerService) PreparePlayer(
	player domain.Player,
) domain.Player {

	// NOTA: Si en el futuro actualizas GetByBirthDate para que retorne varios torneos
	// (ej: tournaments, exists := ...), tendrás que ajustar esta línea.
	tournaments, exists := s.tournamentService.GetByBirthDate(player.BirthDate)

	if !exists {
		player.Status = "ERROR"
		player.Errors = []string{
			"No existe un torneo configurado para el año de nacimiento del jugador",
		}
		return player
	}

	// Ahora la asignación es directa porque ambos son arreglos
	player.Tournaments = tournaments
	// --------------------------------------

	player.Category = s.categoryService.GetCategory(player.BirthDate)

	plan, exists := domainPayment.GetPlan(player.Scholarship)

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
