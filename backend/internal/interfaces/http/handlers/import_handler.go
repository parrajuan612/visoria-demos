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

	tournamentKey := r.FormValue("tournament")

	service := importapp.NewService()

	players, err := service.Execute(file, tournamentKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	store.Save(players)

	response := imports.ImportResponse{
		Summary: imports.Summary{
			Total:  len(players),
			Valid:  len(players),
			Errors: 0,
		},
		Players: players,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
