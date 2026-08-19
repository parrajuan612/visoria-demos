package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/juanparra/visoria-demo/internal/infrastructure/pdf"
	store "github.com/juanparra/visoria-demo/internal/shared"
)

type SendResponse struct {
	Mensaje string `json:"mensaje"`
}

func SendMessages(w http.ResponseWriter, r *http.Request) {
	players := store.Get()

	if len(players) == 0 {
		http.Error(w, "No hay jugadores guardados", http.StatusBadRequest)
		return
	}

	generator := pdf.NewGenerator()

	htmlFile, err := os.Create("panel_envios.html")
	if err != nil {
		http.Error(w, "Error creando panel HTML", http.StatusInternalServerError)
		return
	}
	defer htmlFile.Close()

	htmlFile.WriteString(`
		<html>
		<head><title>Panel de Envíos Manuales (Drive Anti-Fraude)</title>
		<style>
			body { font-family: Arial; padding: 20px; }
			table { border-collapse: collapse; width: 100%; }
			th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
			th { background-color: #f2f2f2; }
			.btn-wa { background-color: #25D366; color: white; padding: 6px 12px; text-decoration: none; border-radius: 4px; display: inline-block; }
			.btn-pdf { background-color: #ff0000; color: white; padding: 6px 12px; text-decoration: none; border-radius: 4px; display: inline-block; }
		</style>
		</head>
		<body>
		<h1>Centro de Control - Envíos Vía Drive</h1>
		<p><b>Instrucciones:</b> Abre el chat, arrastra el PDF y envíalo desde tu número. (Recuerda subir los PDFs a Drive y compartirle la carpeta a la coordinadora).</p>
		<table>
		<tr><th>Jugador</th><th>Acudiente</th><th>PDF Local</th><th>Acción</th></tr>
	`)

	for _, player := range players {
		pdfPath, err := generator.Generate(player)
		if err != nil {
			fmt.Printf("Error generando PDF para %s: %v\n", player.Name, err)
			continue
		}

		// 1. TEXTO PARA LA COORDINADORA (Lo que el papá envía al darle clic al link)
		// No incluimos el % de beca aquí para que el papá no lo edite y la coordinadora revise su Drive.
		textoParaCoordinadora := fmt.Sprintf("Hola Coordinación. Soy el acudiente del jugador %s. Acabo de recibir la propuesta económica de la beca y quiero conocer el siguiente paso del proceso.", player.Name)
		linkCoordinadora := fmt.Sprintf("https://wa.me/573228467206?text=%s", url.QueryEscape(textoParaCoordinadora))

		// 2. TEXTO ORIGINAL QUE TÚ LE ENVÍAS AL PAPÁ (Manteniendo tu estructura)
		mensajeRaw := fmt.Sprintf(`Hola %s,

Te escribimos desde MAJESTIC INTERCAMBIO DEPORTIVO para comunicarte oficialmente la propuesta económica correspondiente a la beca del %d%% que ha ganado %s, gracias a su excelente desempeño y talento demostrado durante las visorias en las que participó.

Esta beca representa una gran oportunidad para viajar a España 🇪🇸, vivir una experiencia deportiva internacional y mostrar su talento en un escenario de alto nivel.

👉 Para continuar el proceso, por favor haz clic en el siguiente enlace para comunicarte con coordinación:
%s`, player.GuardianName, player.Scholarship, player.Name, linkCoordinadora)

		// 3. Codificar el mensaje para la URL de WhatsApp
		mensajeCodificado := url.QueryEscape(mensajeRaw)
		telefono := "57" + player.PrimaryPhone
		linkWAWebPropio := fmt.Sprintf("https://web.whatsapp.com/send?phone=%s&text=%s", telefono, mensajeCodificado)

		// 4. Agregar la fila a la tabla HTML (CORREGIDO EL FILE:// Y LA RUTA RELATIVA)
		fila := fmt.Sprintf(`<tr>
			<td>%s</td>
			<td>%s</td>
			<td><a href="%s" target="_blank" class="btn-pdf">Ver/Descargar PDF</a></td>
			<td><a href="%s" target="_blank" class="btn-wa">Abrir Chat en WA Web</a></td>
		</tr>`, player.Name, player.GuardianName, pdfPath, linkWAWebPropio)

		htmlFile.WriteString(fila)
	}

	htmlFile.WriteString("</table></body></html>")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SendResponse{Mensaje: "Panel de envíos listo."})
}
