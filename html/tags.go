package html

import (
	"fmt"

	"github.com/jpincas/htmlfunc/attributes"
)

const (
	html    = "html"
	title   = "title"
	body    = "body"
	head    = "head"
	text    = "text"
	div     = "div"
	style   = "style"
	p       = "p"
	meta    = "meta"
	link    = "link"
	script  = "script"
	a       = "a"
	nav     = "nav"
	span    = "span"
	section = "section"
	img     = "img"
	svg     = "svg"
	path    = "path"
	polygon = "polygon"
	rect    = "rect"
	line    = "line"
	circle  = "circle"
	br      = "br"
	button  = "button"
	table   = "table"
	tr      = "tr"
	th      = "th"
	td      = "td"
	tbody   = "tbody"
	thead   = "thead"
	h1      = "h1"
	h2      = "h2"
	h3      = "h3"
	h4      = "h4"
	h5      = "h5"
	h6      = "h6"
	ul      = "ul"
	ol      = "ol"
	li      = "li"
	header  = "header"
	form    = "form"
	input   = "input"
)

func Div(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(div, attrs, elements)
}

func Header(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(header, attrs, elements)
}

func Form(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(form, attrs, elements)
}

func Input(attrs attributes.Attributes) Element {
	return selfClosingTag(input, attrs)
}

func P(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(p, attrs, elements)
}

func A(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(a, attrs, elements)
}

func Span(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(span, attrs, elements)
}

func Style(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(style, attrs, elements)
}

func Section(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(section, attrs, elements)
}

func Img(attrs attributes.Attributes) Element {
	return selfClosingTag(img, attrs)
}

func SVG(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(svg, attrs, elements)
}

func Path(attrs attributes.Attributes) Element {
	return basicTag(path, attrs, Elements{})
}

func Polygon(attrs attributes.Attributes) Element {
	return basicTag(polygon, attrs, Elements{})
}

func Rect(attrs attributes.Attributes) Element {
	return basicTag(rect, attrs, Elements{})
}

func Line(attrs attributes.Attributes) Element {
	return basicTag(line, attrs, Elements{})
}

func Circle(attrs attributes.Attributes) Element {
	return basicTag(circle, attrs, Elements{})
}

func Br(_ attributes.Attributes) Element {
	return selfClosingTag(br, attributes.Attributes{})
}

func Nav(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(nav, attrs, elements)
}

func Html(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(html, attrs, elements)
}

func Head(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(head, attrs, elements)
}

func Title(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(title, attrs, elements)
}

func Meta(attrs attributes.Attributes) Element {
	return selfClosingTag(meta, attrs)
}

func Body(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(body, attrs, elements)
}

func Link(attrs attributes.Attributes) Element {
	return selfClosingTag(link, attrs)
}

func Script(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(script, attrs, elements)
}

func Text(i interface{}) Element {
	return Element{
		Tag:  text,
		Text: fmt.Sprintf("%v", i),
	}
}

// Nothing generates a blank element.  The only reason we have the arguments
// is to make the funciton type signature the same as the other construction
// functions
func Nothing(_ attributes.Attributes, _ ...Element) Element {
	return Element{}
}

func Button(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(button, attrs, elements)
}

func Table(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(table, attrs, elements)
}

func THead(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(thead, attrs, elements)
}

func TBody(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(tbody, attrs, elements)
}

func Tr(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(tr, attrs, elements)
}

func Th(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(th, attrs, elements)
}

func Td(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(td, attrs, elements)
}

func H1(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h1, attrs, elements)
}

func H2(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h2, attrs, elements)
}

func H3(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h3, attrs, elements)
}

func H4(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h4, attrs, elements)
}

func H5(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h5, attrs, elements)
}

func H6(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(h6, attrs, elements)
}
func Ul(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(ul, attrs, elements)
}
func Ol(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(ol, attrs, elements)
}
func Li(attrs attributes.Attributes, elements ...Element) Element {
	return basicTag(li, attrs, elements)
}
