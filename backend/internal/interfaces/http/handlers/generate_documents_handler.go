package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juanparra/visoria-demo/internal/infrastructure/pdf"
	store "github.com/juanparra/visoria-demo/internal/shared"
)

type GeneratedDocument struct {
	Name string `json:"name"`
	File string `json:"file"`
}

type GenerateDocumentsResponse struct {
	Total     int                 `json:"total"`
	Generated int                 `json:"generated"`
	Skipped   int                 `json:"skipped"`
	Documents []GeneratedDocument `json:"documents"`
}

func GenerateDocuments(w http.ResponseWriter, r *http.Request) {
	players := store.Get()

	if len(players) == 0 {
		http.Error(
			w,
			"No hay jugadores importados",
			http.StatusBadRequest,
		)
		return
	}

	generator := pdf.NewGenerator()

	documents := []GeneratedDocument{}
	skipped := 0

	for _, player := range players {

		// Por ahora solamente generamos documentos
		// para jugadores completamente válidos.
		if player.Status != "VALID" {
			skipped++
			continue
		}

		pdfPath, err := generator.Generate(player)
		if err != nil {
			fmt.Printf(
				"Error generando PDF para %s: %v\n",
				player.Name,
				err,
			)
			skipped++
			continue
		}

		documents = append(documents, GeneratedDocument{
			Name: player.Name,
			File: pdfPath,
		})
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		GenerateDocumentsResponse{
			Total:     len(players),
			Generated: len(documents),
			Skipped:   skipped,
			Documents: documents,
		},
	)
}
