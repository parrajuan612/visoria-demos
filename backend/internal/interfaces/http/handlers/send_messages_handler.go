package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/juanparra/visoria-demo/internal/infrastructure/pdf"
	"github.com/juanparra/visoria-demo/internal/infrastructure/whatsapp"
	store "github.com/juanparra/visoria-demo/internal/shared"
)

type PlayerStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type SendResponse struct {
	Total   int            `json:"total"`
	Sent    int            `json:"sent"`
	Failed  int            `json:"failed"`
	Players []PlayerStatus `json:"players"`
}

func SendMessages(w http.ResponseWriter, r *http.Request) {
	players := store.Get()

	if len(players) == 0 {
		http.Error(w, "No hay jugadores guardados para enviar mensajes", http.StatusBadRequest)
		return
	}

	client := whatsapp.NewClient(
		os.Getenv("WHATSAPP_TOKEN"),
		os.Getenv("PHONE_NUMBER_ID"),
	)
	generator := pdf.NewGenerator()

	statuses := []PlayerStatus{}
	sent := 0
	failed := 0

	for _, player := range players {
		// 1. Generar PDF local
		pdfPath, err := generator.Generate(player)
		if err != nil {
			fmt.Printf("Error generando PDF para %s: %v\n", player.Name, err)
			failed++
			statuses = append(statuses, PlayerStatus{Name: player.Name, Status: "failed"})
			continue
		}

		// 2. Subir PDF a Meta
		mediaID, err := client.UploadDocument(pdfPath)
		if err != nil {
			fmt.Printf("Error subiendo PDF de %s a Meta: %v\n", player.Name, err)
			failed++
			statuses = append(statuses, PlayerStatus{Name: player.Name, Status: "failed"})
			continue
		}

		phone := "57" + player.PrimaryPhone

		// 3. Enviar texto complementario PRIMERO
		mensaje := fmt.Sprintf(`Hola %s

Te escribimos desde MAJESTIC INTERCAMBIO DEPORTIVO para comunicarte oficialmente la propuesta económica correspondiente a la beca del %d%% que ha ganado %s, gracias a su excelente desempeño y talento demostrado durante las visorias en las que participó.

Esta beca representa una gran oportunidad para viajar a España 🇪🇸, vivir una experiencia deportiva internacional y mostrar su talento en un escenario de alto nivel, aportando a su proceso de formación y crecimiento como deportista.


Estaremos atentos para resolver cualquier inquietud que tengas y brindarte toda la información necesaria sobre el proceso.


🇨🇴🇪🇸 MAJESTIC INTERCAMBIO DEPORTIVO
con la fe intacta ⚽️💫`, player.GuardianName, player.Scholarship, player.Name)

		err = client.Send(phone, mensaje)
		if err != nil {
			fmt.Printf("Error enviando mensaje de texto a %s: %v\n", player.Name, err)
			failed++
			statuses = append(statuses, PlayerStatus{Name: player.Name, Status: "failed"})
			continue
		}

		// 4. Enviar el Documento por WhatsApp DESPUÉS
		filename := fmt.Sprintf("Beca_%s.pdf", player.Name)

		err = client.SendDocument(phone, mediaID, filename)
		if err != nil {
			fmt.Printf("Error enviando documento a %s: %v\n", player.Name, err)
			failed++
			statuses = append(statuses, PlayerStatus{Name: player.Name, Status: "failed"})
			continue
		}

		// Si ambos se enviaron correctamente
		sent++
		statuses = append(statuses, PlayerStatus{Name: player.Name, Status: "sent"})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SendResponse{
		Total:   len(players),
		Sent:    sent,
		Failed:  failed,
		Players: statuses,
	})
}
