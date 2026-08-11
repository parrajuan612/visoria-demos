package handlers

import (
	"encoding/json"
	"net/http"

	store "github.com/juanparra/visoria-demo/internal/shared"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	players := store.Get()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"total":   len(players),
		"players": players,
	})
}
