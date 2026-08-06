package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/juanparra/visoria-demo/internal/interfaces/http/routes"
)

func main() {
	godotenv.Load()
	server := &http.Server{
		Addr:    ":8880",
		Handler: enableCORS(routes.RegisterRoutes()),
	}

	log.Println("Servidor iniciado en http://localhost:8880")

	log.Fatal(server.ListenAndServe())
}
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
