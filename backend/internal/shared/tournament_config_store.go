package store

import (
	"sync"

	tournamentapp "github.com/juanparra/visoria-demo/internal/application/tournament"
)

var (
	configMu sync.RWMutex
	config   tournamentapp.TournamentConfig
)

func SaveTournamentConfig(newConfig tournamentapp.TournamentConfig) {
	configMu.Lock()
	defer configMu.Unlock()

	config = newConfig
}

func GetTournamentConfig() tournamentapp.TournamentConfig {
	configMu.RLock()
	defer configMu.RUnlock()

	return config
}
