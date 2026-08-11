package excel

import (
	"strconv"
	"strings"
	"time"

	"github.com/juanparra/visoria-demo/internal/domain"
)

func get(row []string, index int) string {
	if index >= len(row) {
		return ""
	}

	return strings.TrimSpace(row[index])
}

func parseBirthDate(value string) (time.Time, error) {
	formats := []string{
		"01-02-06",
		"01-02-2006",
		"02/01/2006",
		"02-01-2006",
		"2006-01-02",
	}

	for _, format := range formats {
		if date, err := time.Parse(format, value); err == nil {
			return date, nil
		}
	}

	return time.Time{}, nil
}

func MapPlayers(rows [][]string) []domain.Player {

	var players []domain.Player

	for i, row := range rows {

		if i == 0 {
			continue
		}

		scholarship, _ := strconv.Atoi(get(row, 2))

		birthDate, _ := parseBirthDate(get(row, 3))

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
