package services

import "github.com/juanparra/visoria-demo/internal/domain"

type PlayerService struct {
	categoryService   *CategoryService
	tournamentService *TournamentService
}

func NewPlayerService(
	categoryService *CategoryService,
	tournamentService *TournamentService,
) *PlayerService {
	return &PlayerService{
		categoryService:   categoryService,
		tournamentService: tournamentService,
	}
}

func (s *PlayerService) PreparePlayer(
	player domain.Player,
	tournamentKey string,
) domain.Player {

	player.Category = s.categoryService.GetCategory(player.BirthDate)

	tournament, exists := s.tournamentService.Get(tournamentKey)

	if exists {
		player.Tournament = &tournament
	}

	return player
}
