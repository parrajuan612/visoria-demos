package store

import (
	"sync"

	tournamentapp "github.com/juanparra/visoria-demo/internal/application/tournament"
)

var (
	configMu sync.RWMutex
	configs  []tournamentapp.TournamentConfig
)

func SaveTournamentConfig(config tournamentapp.TournamentConfig) {
	configMu.Lock()
	defer configMu.Unlock()

	configs = append(configs, config)
}

func GetTournamentConfigs() []tournamentapp.TournamentConfig {
	configMu.RLock()
	defer configMu.RUnlock()

	return append([]tournamentapp.TournamentConfig(nil), configs...)
}
