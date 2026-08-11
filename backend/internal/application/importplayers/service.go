package importapp

import (
	"io"

	tournamentapp "github.com/juanparra/visoria-demo/internal/application/tournament"
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

	return &Service{
		reader: excel.NewReader(),
		playerService: services.NewPlayerService(
			validator,
		),
	}
}

func (s *Service) Execute(
	file io.Reader,
	config tournamentapp.TournamentConfig,
) ([]domain.Player, error) {

	rows, err := s.reader.Read(file)
	if err != nil {
		return nil, err
	}

	players := excel.MapPlayers(rows)

	for i := range players {
		players[i] = s.playerService.PreparePlayer(
			players[i],
			config,
		)
	}

	return players, nil
}
