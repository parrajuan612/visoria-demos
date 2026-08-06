package routes

import (
	"net/http"

	"github.com/juanparra/visoria-demo/internal/interfaces/http/handlers"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/players/import", handlers.ImportPlayers)
	mux.HandleFunc("POST /api/v1/test-whatsapp", handlers.TestWhatsApp)
	mux.HandleFunc("POST /api/v1/messages/send", handlers.SendMessages) // Ruta añadida

	return mux
}
