package handlers

import (
	"encoding/json"
	"net/http"

	tournamentapp "github.com/juanparra/visoria-demo/internal/application/tournament"
	store "github.com/juanparra/visoria-demo/internal/shared"
)

func SaveTournamentConfig(w http.ResponseWriter, r *http.Request) {
	var config tournamentapp.TournamentConfig

	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, "Configuración inválida", http.StatusBadRequest)
		return
	}

	store.SaveTournamentConfig(config)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Configuración guardada correctamente",
		"config":  config,
	})
}

func GetTournamentConfig(w http.ResponseWriter, r *http.Request) {
	configs := store.GetTournamentConfigs()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(configs)
}
