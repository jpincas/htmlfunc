package htmlfunc

import (
	"github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

// ConstructTable is for quickly constructing simple tables with global attributes
func ConstructTable(attrs attributes.Attributes, headerRow []string, rows []h.Elements) h.Element {
	return ComplexTable{
		GlobalAttrs: attrs,
		HeaderRow:   headerRow,
		Rows:        rows,
	}.Render()
}

type ComplexTable struct {
	FirstColumnIsTitle                     bool
	HeaderRow                              []string
	Rows                                   []h.Elements
	GlobalAttrs                            attributes.Attributes
	HeadAttrs, HeadRowAttrs, HeadCellAttrs attributes.Attributes
	LastRowAttrs, LastRowCellAttrs         attributes.Attributes
	BodyAttrs, BodyRowAttrs, BodyCellAttrs attributes.Attributes
}

// ConstructComplexTable is for constructing more complex tables with inline attributes
// at every level.  Useful, for example, for tables in HTML emails
func (complexTable ComplexTable) Render() h.Element {
	headerCells := h.Els()
	for _, columnHeading := range complexTable.HeaderRow {
		headerCells = append(headerCells, h.Th(complexTable.HeadCellAttrs, h.Text(columnHeading)))
	}
	header := h.THead(complexTable.HeadAttrs, h.Tr(complexTable.HeadRowAttrs, headerCells...))

	bodyRows := h.Els()
	for i, row := range complexTable.Rows {
		tableCells := h.Els()
		var cellAttrs, rowAttrs attributes.Attributes

		if i < (len(complexTable.Rows)-1) || len(complexTable.Rows) == 1 {
			cellAttrs = complexTable.BodyCellAttrs
			rowAttrs = complexTable.BodyRowAttrs
		} else {
			cellAttrs = complexTable.LastRowCellAttrs
			rowAttrs = complexTable.LastRowAttrs
		}

		for i, cellValue := range row {
			cellFunc := h.Td
			if i == 0 {
				cellFunc = h.Th
			}

			tableCells = append(tableCells, cellFunc(cellAttrs, cellValue))
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
