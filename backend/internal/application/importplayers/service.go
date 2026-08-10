package importapp

import (
	"io"

	"github.com/juanparra/visoria-demo/internal/domain"
	"github.com/juanparra/visoria-demo/internal/domain/services"
	"github.com/juanparra/visoria-demo/internal/infrastructure/excel"
)

type Service struct {
	reader        *excel.Reader
	playerService *services.PlayerService
}

func NewService() *Service {
	categoryService := services.NewCategoryService()
	tournamentService := services.NewTournamentService()

	return &Service{
		reader: excel.NewReader(),
		playerService: services.NewPlayerService(
			categoryService,
			tournamentService,
		),
	}
}

func (s *Service) Execute(
	file io.Reader,
	tournamentKey string,
) ([]domain.Player, error) {

	rows, err := s.reader.Read(file)
	if err != nil {
		return nil, err
	}

	players := excel.MapPlayers(rows)

	for i := range players {
		players[i] = s.playerService.PreparePlayer(
			players[i],
			tournamentKey,
		)
	}

	return players, nil
}
