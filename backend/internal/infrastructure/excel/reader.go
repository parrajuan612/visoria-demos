package excel

import (
	"bytes"
	"io"

	"github.com/xuri/excelize/v2"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Read(file io.Reader) ([][]string, error) {

	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	workbook, err := excelize.OpenReader(bytes.NewReader(buffer))
	if err != nil {
		return nil, err
	}

	sheet := workbook.GetSheetName(0)

	rows, err := workbook.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
