package pdf

import (
	"fmt"
	"os"

	"github.com/juanparra/visoria-demo/internal/domain"
	"github.com/jung-kurt/gofpdf"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(player domain.Player) (string, error) {

	err := os.MkdirAll("uploads/pdfs", os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("uploads/pdfs/%s.pdf", player.Name)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Traductor para soportar tildes y eñes (ñ)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1252")
	pdf.ImageOptions("logo.png", 10, 10, 70, 0, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
	// --- ENCABEZADO ---
	pdf.SetY(35)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 5, tr("BOGOTÁ – 09/07/2026"))
	pdf.Ln(6)

	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(0, 5, tr(" MIC FOOTBALL / IDA EASTER CUP 2027"))
	pdf.Ln(10)

	// --- DATOS DEL JUGADOR ---
	colWidth := 40.0
	rowHeight := 5.0

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
	pdf.CellFormat(0, rowHeight, player.BirthDate.Format("02/01/2006"), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(colWidth, rowHeight, tr("Móvil:"), "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, rowHeight, "(+57) "+player.PrimaryPhone, "", 1, "L", false, 0, "")
	pdf.Ln(8)

	// --- REFERENCIA Y TÍTULO ---
	pdf.SetFont("Arial", "BU", 10)
	pdf.CellFormat(0, rowHeight, tr("Ref. Propuesta Jugadores seleccionados"), "", 1, "L", false, 0, "")
	pdf.Ln(4)

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, rowHeight, tr("PROGRAMA DE INTERCAMBIO DEPORTIVO"), "", 1, "L", false, 0, "")
	pdf.Ln(2)

	// --- CUERPO DEL TEXTO ---
	pdf.SetFont("Arial", "", 10)
	cuerpo1 := tr("Consiste en la participación del ALUMNO – DEPORTISTA en el intercambio deportivo, torneo de futbol en ESPAÑA. Sólo se participará de un torneo.\n\n") +
		tr("TORNEO MIC FOOTBALL COSTA BRAVA BARCELONA ESPAÑA\nTORNEO IDA EASTER CUP COMUNIDAD VALENCIANA\n\n") +
		tr("CATEGORIAS:\nINFANTIL nacidos en el año 2013-2014\nCADETES nacidos en el año 2011-2012\nJUVENILES nacidos en el año 2008-2009-2010\n\n") +
		tr("Fechas para el viaje:\nSALIDA el 18/03/2027\nLLEGADA a España el 19/03/2027, ingresarán con cena y se recogerán en el aeropuerto en el transcurso del día máximo hasta las 5:00 pm hora España.\nREGRESO el 28/03/2027 salen con Desayuno y almuerzo, los buses llegaran 3:00 pm para ir de nuevo al aeropuerto, se sugiere que los vuelos sean después de las 8:00 pm.\n\n") +
		tr("INCLUYE: Alimentación – hospedaje – transporte interno en España (aeropuerto-torneo-turismo-entrenos-hoteles) – Indumentaria – Inscripción Torneo – Visitas turísticas, Visorias por parte de los clubes y las academias que tenemos convenios para que el jugador continúe con su primer proceso en Europa según su desempeño, este puede ser de 30-60-90 días.\n\n") +
		tr("En el mes de enero/2027 se realizará la pre-temporada en Bogotá. (fechas por confirmar)")

	pdf.MultiCell(0, 5, cuerpo1, "", "J", false)
	pdf.Ln(6)

	// --- TABLA DE COSTOS Y BECA ---
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 7, fmt.Sprintf(tr("COSTOS: EL JUGADOR OBTUVO BECA DEL %d%%"), player.Scholarship), "", 1, "C", false, 0, "")
	pdf.Ln(2)

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(80, 6, tr("BECA"), "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 6, tr("VALOR € EUROS"), "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	filas := [][]string{
		{"Sin Beca", "€ 2.800"},
		{"Beca al 30%", "€ 1.960"},
		{"Beca al 50%", "€ 1.550"},
		{"Beca al 70%", "€ 990"},
		{"Beca al 100%", "€ 200 administración"},
		{"Acompañante", "€ 1.800"},
	}

	for _, fila := range filas {
		pdf.CellFormat(80, 6, tr(fila[0]), "1", 0, "L", false, 0, "")
		pdf.CellFormat(80, 6, tr(fila[1]), "1", 1, "C", false, 0, "")
	}
	pdf.Ln(6)

	// --- NOTAS FINALES Y CUENTAS ---
	cuerpo2 := tr("NO INCLUYE: Tiquetes Aéreos, Emisión del pasaporte, Seguro de viaje.\n\n") +
		tr("Los pagos se deben realizar a la cuenta de ahorros # 22546881826 de BANCOLOMBIA o en DAVIVIENDA cuenta de ahorros # 0570008380462534 las dos a nombre de Suysan Colmenares Camargo C.C 79739776. SEGÚN LOS VALORES VENTA DE DIVISAS CAMBIOS VANCOUVER (página web cambiosvancouver.com)\n\n") +
		tr("Nota: Los dineros pagados NO tienen devolución si el jugador o el acompañante no viaja al programa, pero sí tendrá devolución en servicios del programa Majestic Intercambio Deportivo y se congelará por máximo 1 año; siempre y cuando los causales sean por lesión con excusa soportada de la EPS igual el acompañante.")

	pdf.MultiCell(0, 5, cuerpo2, "", "J", false)
	pdf.Ln(15)

	// --- FIRMA ---
	// Si quieres insertar la imagen de la firma real, descomenta la siguiente línea y asegúrate de tener 'firma.png' en la raíz.
	// pdf.ImageOptions("firma.png", 10, pdf.GetY(), 40, 0, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
	// pdf.Ln(25) // Espacio para que la imagen no se superponga con el texto

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 5, tr("Cordialmente:"), "", 1, "L", false, 0, "")
	pdf.Ln(10) // Espacio simulando la firma a mano

	pdf.CellFormat(0, 5, tr("SUYSAN COLMENARES C."), "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 5, tr("Coordinador Programa"), "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 5, tr("Móvil (+57) 3202411029"), "", 1, "L", false, 0, "")

	err = pdf.OutputFileAndClose(fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
