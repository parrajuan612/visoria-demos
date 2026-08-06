package store

import (
	"sync"

	"github.com/juanparra/visoria-demo/internal/domain"
)

var (
	mu      sync.RWMutex
	players []domain.Player
)

func Save(list []domain.Player) {
	mu.Lock()
	defer mu.Unlock()
	players = list
}

func Get() []domain.Player {
	mu.RLock()
	defer mu.RUnlock()
	return players
}
