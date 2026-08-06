package importapp

import (
	"fmt"
	"io"

	"github.com/juanparra/visoria-demo/internal/domain"
	"github.com/juanparra/visoria-demo/internal/infrastructure/excel"
)

type Service struct {
	reader *excel.Reader
}

func NewService() *Service {
	return &Service{
		reader: excel.NewReader(),
	}
}

func (s *Service) Execute(file io.Reader) ([]domain.Player, error) {

	rows, err := s.reader.Read(file)
	if err != nil {
		return nil, err
	}

	fmt.Println(rows)

	players := excel.MapPlayers(rows)

	return players, nil
}
