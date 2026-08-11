package validation

import (
	"strings"
	"time"

	"github.com/juanparra/visoria-demo/internal/domain"
)

type PlayerValidator struct{}

func NewPlayerValidator() *PlayerValidator {
	return &PlayerValidator{}
}

func (v *PlayerValidator) Validate(player domain.Player) ([]string, []string) {

	var errors []string
	var warnings []string

	if strings.TrimSpace(player.Name) == "" {
		errors = append(errors, "El nombre del jugador es obligatorio")
	}

	if strings.TrimSpace(player.Club) == "" {
		errors = append(errors, "El club es obligatorio")
	}

	if player.BirthDate.IsZero() {
		errors = append(errors, "La fecha de nacimiento es obligatoria")
	}

	if strings.TrimSpace(player.GuardianName) == "" {
		errors = append(errors, "El nombre del acudiente es obligatorio")
	}

	if strings.TrimSpace(player.PrimaryPhone) == "" {
		errors = append(errors, "El teléfono principal es obligatorio")
	}

	if player.Category == "" || player.Category == "NO DEFINIDA" {
		warnings = append(warnings, "El jugador no tiene una categoría definida")
	}

	return errors, warnings
}

func isValidDate(date time.Time) bool {
	return !date.IsZero()
}
