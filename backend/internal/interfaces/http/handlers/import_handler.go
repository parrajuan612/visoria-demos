package handlers

import (
	"encoding/json"
	"net/http"

	importapp "github.com/juanparra/visoria-demo/internal/application/import"
	store "github.com/juanparra/visoria-demo/internal/shared"
)

func ImportPlayers(w http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	service := importapp.NewService()

	players, err := service.Execute(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Guarda los jugadores en memoria antes de responder
	store.Save(players)

	response := importapp.ImportResponse{
		Summary: importapp.Summary{
			Total:  len(players),
			Valid:  len(players),
			Errors: 0,
		},
		Players: players,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
