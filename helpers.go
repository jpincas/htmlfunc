package htmlfunc

import (
	h "github.com/jpincas/htmlfunc/html"
)

func ConstructTable(attrs []h.Attribute, headerRow []string, rows [][]interface{}) h.Element {
	headerCells := h.Els()
	for _, columnHeading := range headerRow {
		headerCells = append(headerCells, h.Td(h.Attrs(), h.Els(h.Text(columnHeading))))
	}
	header := h.THead(h.Attrs(), h.Els(h.Tr(h.Attrs(), headerCells)))

	bodyRows := h.Els()
	for _, row := range rows {
		tableCells := h.Els()
		for _, cellValue := range row {
			tableCells = append(tableCells, h.Td(h.Attrs(), h.Els(h.Text(cellValue))))
		}

		bodyRows = append(bodyRows, h.Tr(h.Attrs(), tableCells))
	}
	body := h.TBody(h.Attrs(), bodyRows)

	return h.Table(attrs, h.Els(header, body))
}
