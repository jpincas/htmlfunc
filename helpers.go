package htmlfunc

import (
	h "github.com/jpincas/htmlfunc/html"
)

// ConstructTable is for quickly constructing simple tables with global attributes
func ConstructTable(attrs []h.Attribute, headerRow []string, rows [][]h.Element) h.Element {
	return ConstructComplexTable(ComplexTable{
		GlobalAttrs: attrs,
		HeaderRow:   headerRow,
		Rows:        rows,
	})
}

type ComplexTable struct {
	HeaderRow                              []string
	Rows                                   [][]h.Element
	GlobalAttrs                            []h.Attribute
	HeadAttrs, HeadRowAttrs, HeadCellAttrs []h.Attribute
	LastRowAttrs, LastRowCellAttrs         []h.Attribute
	BodyAttrs, BodyRowAttrs, BodyCellAttrs []h.Attribute
}

// ConstructComplexTable is for constructing more complex tables with inline attributes
// at every level.  Useful, for example, for tables in HTML emails
func ConstructComplexTable(complexTable ComplexTable) h.Element {
	headerCells := h.Els()
	for _, columnHeading := range complexTable.HeaderRow {
		headerCells = append(headerCells, h.Th(complexTable.HeadCellAttrs, h.Text(columnHeading)))
	}
	header := h.THead(complexTable.HeadAttrs, h.Tr(complexTable.HeadRowAttrs, headerCells...))

	bodyRows := h.Els()
	for i, row := range complexTable.Rows {
		tableCells := h.Els()
		var cellAttrs, rowAttrs []h.Attribute

		if i < (len(complexTable.Rows) - 1) {
			cellAttrs = complexTable.BodyCellAttrs
			rowAttrs = complexTable.BodyRowAttrs
		} else {
			cellAttrs = complexTable.LastRowCellAttrs
			rowAttrs = complexTable.LastRowAttrs
		}

		for _, cellValue := range row {
			tableCells = append(tableCells, h.Td(cellAttrs, cellValue))
		}

		bodyRows = append(bodyRows, h.Tr(rowAttrs, tableCells...))
	}

	body := h.TBody(complexTable.BodyAttrs, bodyRows...)

	return h.Table(
		complexTable.GlobalAttrs,
		header,
		body,
	)
}

// Join separates HTML elements with the specifid separator
func Join(els []h.Element, sep h.Element) (res []h.Element) {
	for _, el := range els {
		res = append(res, el, sep)
	}

	// remove the final separator
	if len(res) > 0 {
		res = res[:len(res)-1]
	}

	return
}

// JoinIf separates HTML elements with the specifid separator, if they are non blank
func JoinIf(els []h.Element, sep h.Element) (res []h.Element) {
	for _, el := range els {
		if !el.IsEmpty() {
			res = append(res, el, sep)
		}
	}

	// remove the final separator
	if len(res) > 0 {
		res = res[:len(res)-1]
	}

	return
}
