package pdf

import (
	"fmt"
	"os"
	"time"

	"github.com/juanparra/visoria-demo/internal/domain"
	"github.com/jung-kurt/gofpdf"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func formatDate(d time.Time) string {
	if d.IsZero() {
		return "19/03/2027"
	}
	return d.Format("02/01/2006")
}

func (g *Generator) Generate(player domain.Player) (string, error) {

	// 1. Validaciones
	if len(player.Tournaments) == 0 {
		return "", fmt.Errorf("el jugador %s no tiene torneos asignados", player.Name)
	}
	if player.PaymentPlan == nil {
		return "", fmt.Errorf("el jugador %s no tiene plan de pagos asignado", player.Name)
	}

	err := os.MkdirAll("uploads/pdfs", os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("uploads/pdfs/%s.pdf", player.Name)

	pdf := gofpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1252")

	// --- MARCA DE AGUA (FONDO DIFUMINADO EN TODAS LAS PÁGINAS) ---
	pdf.SetHeaderFunc(func() {
		pdf.SetAlpha(0.1, "Normal")
		pdf.ImageOptions("Isologo_6@2x (1).png", 30, 75, 150, 0, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")
		pdf.SetAlpha(1.0, "Normal")
	})

	pdf.AddPage()

	// Desactivamos el salto de página automático para forzar que todo quepa en una hoja
	pdf.SetAutoPageBreak(false, 0)

	// Logo principal superior izquierdo
	pdf.ImageOptions("logo.png", 10, 10, 70, 0, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")

	// --- ENCABEZADO ---
	pdf.SetY(32) // Ligeramente más arriba
	pdf.SetFont("Arial", "", 10)
	fechaActual := time.Now().Format("02/01/2006")
	pdf.Cell(0, 4, tr(fmt.Sprintf("BOGOTÁ – %s", fechaActual)))
	pdf.Ln(5)

	pdf.SetFont("Arial", "B", 10)
	nombresTorneos := "Soccer Melgar del 07/08/2026 al 09/08/2026"
	// for i, t := range player.Tournaments {
	//  if i > 0 {
	//      nombresTorneos += " / "
	//  }
	//  nombresTorneos += t.Name
	// }
	pdf.MultiCell(0, 4, tr(nombresTorneos), "", "L", false)
	pdf.Ln(4)

	// --- DATOS DEL JUGADOR ---
	colWidth := 40.0
	rowHeight := 4.0 // Altura reducida para ahorrar espacio

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(colWidth, rowHeight, tr("Acudiente:"), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, rowHeight, tr(player.GuardianName), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(colWidth, rowHeight, tr("Jugador:"), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, rowHeight, tr(player.Name), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(colWidth, rowHeight, tr("Club:"), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, rowHeight, tr(player.Club), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(colWidth, rowHeight, tr("Fecha de Nacimiento:"), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, rowHeight, formatDate(player.BirthDate), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(colWidth, rowHeight, tr("Móvil:"), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, rowHeight, "(+57) "+player.PrimaryPhone, "", 1, "L", false, 0, "")
	pdf.Ln(4) // Espaciado reducido

	// --- REFERENCIA Y TÍTULO ---
	pdf.SetFont("Arial", "BU", 10)
	pdf.CellFormat(0, rowHeight, tr("Ref. Propuesta Jugadores seleccionados"), "", 1, "L", false, 0, "")
	pdf.Ln(2)

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, rowHeight, tr("PROGRAMA DE INTERCAMBIO DEPORTIVO"), "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// --- CUERPO DEL TEXTO ---
	textoCategoria := player.Category
	if val, ok := player.Tournaments[0].CategoryText[player.Category]; ok {
		textoCategoria = val
	}

	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 4, tr("Consiste en la participación del ALUMNO – DEPORTISTA en el intercambio deportivo, torneo de futbol en ESPAÑA."), "", "J", false)

	pdf.SetFont("Arial", "U", 10)
	pdf.CellFormat(0, 4, tr("solamente participará de un torneo"), "", 1, "L", false, 0, "")

	// Modificación: Aplicar negrilla al nombre de los torneos
	pdf.SetFont("Arial", "B", 10)
	for _, t := range player.Tournaments {
		pdf.CellFormat(0, 4, tr(fmt.Sprintf("- %s", t.Name)), "", 1, "L", false, 0, "")
	}
	pdf.SetFont("Arial", "", 10) // Restauramos fuente normal
	pdf.Ln(2)

	// CATEGORIAS:
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 4, tr("CATEGORIAS:"), "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 4, tr(textoCategoria), "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// Fechas para el viaje:
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 4, tr("Fechas para el viaje:"), "", 1, "L", false, 0, "")

	tViaje := player.Tournaments[0]

	pdf.SetFont("Arial", "B", 10)
	pdf.Write(4, tr("SALIDA el "))
	pdf.SetFont("Arial", "", 10)
	pdf.Write(4, tr(formatDate(tViaje.Travel.DepartureDate)+"\n"))

	pdf.SetFont("Arial", "B", 10)
	pdf.Write(4, tr("LLEGADA a España al aeropuerto de BARCELONA el 19/03/2027"))
	pdf.SetFont("Arial", "", 10)
	pdf.Write(4, tr(", ingresarán con cena y se recogerán en el aeropuerto en el transcurso del día máximo hasta las 5:00 pm hora España.\n"))

	pdf.SetFont("Arial", "B", 10)
	pdf.Write(4, tr("REGRESO el "))
	pdf.SetFont("Arial", "", 10)
	pdf.Write(4, tr(fmt.Sprintf("%s salen con Desayuno y almuerzo, los buses llegaran 3:00 pm para ir de nuevo al aeropuerto, se sugiere que los vuelos sean después de las 8:00 pm.\n", formatDate(tViaje.Travel.ReturnDate))))

	// Modificación: Separar "INCLUYE:" en negrilla del resto del texto
	pdf.SetFont("Arial", "B", 10)
	pdf.Write(4, tr("INCLUYE: "))
	pdf.SetFont("Arial", "", 10)
	pdf.Write(4, tr("Alimentación – hospedaje – transporte interno en España (aeropuerto-torneo-turismo-entrenos-hoteles) – Indumentaria – Inscripción Torneo – Visitas turísticas, Visorias por parte de los clubes y las academias que tenemos convenios para que el jugador continúe con su primer proceso en Europa según su desempeño, este puede ser de 30-60-90 días.\n"))

	// Modificación: Corrección del bug "%!(EXTRA int=2026)" que aparecía en la pre-temporada
	pdf.MultiCell(0, 4, tr("En el mes de enero/2027 se realizará la pre-temporada en Bogotá. (fechas por confirmar)"), "", "J", false)
	pdf.Ln(3)

	// --- TABLA DE COSTOS Y BECA ---
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(0, 5, fmt.Sprintf(tr("COSTOS: EL JUGADOR OBTUVO BECA DEL %d%%"), player.Scholarship), "", 1, "C", false, 0, "")
	pdf.Ln(1)

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(80, 5, tr("BECA"), "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 5, tr("VALOR € EUROS"), "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	filas := [][]string{
		{"Sin Beca", "€ 2.800"},
		{"Beca al 30%", "€ 1.960"},
		{"Beca al 50%", "€ 1.550"},
		{"Beca al 70%", "€ 990"},
		{"Beca al 100%", "€ 200 administración"},
		{"Acompañante", "€ 1.800"},
	}

	for i, fila := range filas {
		fill := false
		percent := -1
		switch i {
		case 0:
			percent = 0
		case 1:
			percent = 30
		case 2:
			percent = 50
		case 3:
			percent = 70
		case 4:
			percent = 100
		}

		if percent == player.Scholarship {
			pdf.SetFillColor(255, 255, 0) // Amarillo
			fill = true
		}

		pdf.CellFormat(80, 5, tr(fila[0]), "1", 0, "L", fill, 0, "")
		pdf.CellFormat(80, 5, tr(fila[1]), "1", 1, "C", fill, 0, "")
	}
	pdf.Ln(2)

	cuerpo2 := tr("NO INCLUYE: Tiquetes Aéreos, Emisión del pasaporte, Seguro de viaje.\n") +
		tr("Los pagos se deben realizar a la cuenta de ahorros # 22546881826 de BANCOLOMBIA o en DAVIVIENDA cuenta de ahorros # 0570008380462534 las dos a nombre de Suysan Colmenares Camargo C.C 79739776. SEGÚN LOS VALORES VENTA DE DIVISAS CAMBIOS VANCOUVER (página web cambiosvancouver.com)")

	pdf.MultiCell(0, 4, cuerpo2, "", "J", false)
	pdf.Ln(3)

	// --- TABLA RESPONSIVE DE PAGOS ---
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 4, tr("Programación Pagos:"), "", 1, "L", false, 0, "")

	numPagos := 3
	if player.PaymentPlan.Payment3 == 0 {
		numPagos = 2
	}
	if player.PaymentPlan.Payment2 == 0 {
		numPagos = 1
	}

	anchoTotal := 180.0
	anchoColumna := anchoTotal / float64(numPagos+1)

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(anchoColumna, 5, tr("PROGRAMA"), "1", 0, "C", false, 0, "")
	pdf.CellFormat(anchoColumna, 5, tr(fmt.Sprintf("1er pago %d€", player.PaymentPlan.Payment1)), "1", 0, "C", false, 0, "")

	if numPagos >= 2 {
		pdf.CellFormat(anchoColumna, 5, tr(fmt.Sprintf("2do pago %d€", player.PaymentPlan.Payment2)), "1", 0, "C", false, 0, "")
	}
	if numPagos == 3 {
		pdf.CellFormat(anchoColumna, 5, tr(fmt.Sprintf("3er pago %d€", player.PaymentPlan.Payment3)), "1", 0, "C", false, 0, "")
	}
	pdf.Ln(5)

	pdf.SetFont("Arial", "", 10)
	tPago := player.Tournaments[0]

	pdf.CellFormat(anchoColumna, 5, tr("TORNEO ESPAÑA"), "1", 0, "C", false, 0, "")
	pdf.CellFormat(anchoColumna, 5, formatDate(tPago.Payments.Payment1Date), "1", 0, "C", false, 0, "")

	if numPagos >= 2 {
		pdf.CellFormat(anchoColumna, 5, formatDate(tPago.Payments.Payment2Date), "1", 0, "C", false, 0, "")
	}
	if numPagos == 3 {
		pdf.CellFormat(anchoColumna, 5, formatDate(tPago.Payments.Payment3Date), "1", 0, "C", false, 0, "")
	}
	pdf.Ln(6)

	cuerpo3 := tr("Nota: Los dineros pagados NO tienen devolución si el jugador o el acompañante no viaja al programa, pero sí tendrá devolución en servicios del programa Majestic Intercambio Deportivo y se congelará por máximo 1 año; siempre y cuando los causales sean por lesión con excusa soportada de la EPS igual el acompañante.")
	pdf.SetFont("Arial", "", 9) // Bajamos a tamaño 9 para la nota y ahorrar un poco de espacio
	pdf.MultiCell(0, 4, cuerpo3, "", "J", false)
	pdf.Ln(4)

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 4, tr("Cordialmente:"), "", 1, "L", false, 0, "")
	pdf.Ln(1)

	xFirma := pdf.GetX()
	yFirma := pdf.GetY()

	pdf.ImageOptions("firma.png", xFirma, yFirma-3, 40, 0, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")

	// Modificación: Se aumentó a +24 el salto de Y para que la firma no se solape con el nombre
	pdf.SetY(yFirma + 24)

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 4, tr("SUYSAN COLMENARES C."), "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 4, tr("Coordinador Programa"), "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 4, tr("Móvil (+57) 3202411029"), "", 1, "L", false, 0, "")

	// --- REDES SOCIALES Y CONTACTO (PIE DE PÁGINA) ---
	// Se sube ligeramente el margen Y a -18mm para asegurar que cabe sin romper la página
	pdf.SetY(-18)
	pdf.SetFont("Times", "", 12)
	pdf.CellFormat(0, 4, tr("Instagram @majestic_intercambio"), "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 4, tr("Facebook Majestic Intercambio"), "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 4, tr("Móvil +57 3202411029"), "", 1, "C", false, 0, "")

	err = pdf.OutputFileAndClose(fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
