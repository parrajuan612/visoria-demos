package excel

import (
	"strconv"
	"time"

	"github.com/juanparra/visoria-demo/internal/domain"
)

func get(row []string, index int) string {
	if index >= len(row) {
		return ""
	}
	return row[index]
}

func MapPlayers(rows [][]string) []domain.Player {

	var players []domain.Player

	for i, row := range rows {

		if i == 0 {
			continue
		}

		scholarship, _ := strconv.Atoi(get(row, 2))

		birthDate, _ := time.Parse("02/01/2006", get(row, 3))

		player := domain.Player{
			Name:           get(row, 0),
			Club:           get(row, 1),
			Scholarship:    scholarship,
			BirthDate:      birthDate,
			GuardianName:   get(row, 4),
			PrimaryPhone:   get(row, 5),
			SecondaryPhone: get(row, 6),
		}

		players = append(players, player)
	}

	return players
}
