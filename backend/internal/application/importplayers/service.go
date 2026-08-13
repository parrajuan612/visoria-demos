package importapp

import (
	"io"

	"github.com/juanparra/visoria-demo/internal/domain"
	"github.com/juanparra/visoria-demo/internal/domain/services"
	"github.com/juanparra/visoria-demo/internal/domain/validation"
	"github.com/juanparra/visoria-demo/internal/infrastructure/excel"
)

type Service struct {
	reader        *excel.Reader
	playerService *services.PlayerService
}

func NewService() *Service {
	validator := validation.NewPlayerValidator()
	tournamentService := services.NewTournamentService()
	categoryService := services.NewCategoryService()

	return &Service{
		reader: excel.NewReader(),
		playerService: services.NewPlayerService(
			tournamentService,
			validator,
			categoryService,
		),
	}
}

func (s *Service) Execute(
	file io.Reader,
) ([]domain.Player, error) {

	rows, err := s.reader.Read(file)
	if err != nil {
		return nil, err
	}

	players := excel.MapPlayers(rows)

	for i := range players {
		players[i] = s.playerService.PreparePlayer(players[i])
	}

	return players, nil
}
