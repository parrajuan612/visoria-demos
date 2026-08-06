package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/juanparra/visoria-demo/internal/infrastructure/whatsapp"
)

func TestWhatsApp(w http.ResponseWriter, r *http.Request) {
	client := whatsapp.NewClient(
		os.Getenv("WHATSAPP_TOKEN"),
		os.Getenv("PHONE_NUMBER_ID"),
	)

	// Reemplaza con tu número verificado (+57...)
	err := client.Send(
		"573209323449",
		"🚀 Hola, este mensaje fue enviado automáticamente desde el sistema de Visorías desarrollado en Go.",
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error enviando mensaje: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok", "message": "Mensaje de WhatsApp enviado correctamente"}`))
}
