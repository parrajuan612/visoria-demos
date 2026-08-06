package importapp

import "github.com/juanparra/visoria-demo/internal/domain"

type Summary struct {
	Total  int `json:"total"`
	Valid  int `json:"valid"`
	Errors int `json:"errors"`
}

type ImportResponse struct {
	Summary Summary         `json:"summary"`
	Players []domain.Player `json:"players"`
}
