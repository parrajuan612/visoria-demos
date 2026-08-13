package handlers

import (
	"encoding/json"
	"net/http"

	imports "github.com/juanparra/visoria-demo/internal/application/import"
	importapp "github.com/juanparra/visoria-demo/internal/application/importplayers"
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

	store.Save(players)

	response := imports.ImportResponse{
		Summary: imports.Summary{
			Total:    len(players),
			Valid:    0,
			Errors:   0,
			Warnings: 0,
		},
		Players: players,
	}

	for _, player := range players {
		switch player.Status {
		case "VALID":
			response.Summary.Valid++
		case "WARNING":
			response.Summary.Warnings++
		case "ERROR":
			response.Summary.Errors++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
