package html

import "fmt"

const (
	html    = "html"
	title   = "title"
	body    = "body"
	head    = "head"
	div     = "div"
	p       = "p"
	meta    = "meta"
	link    = "link"
	script  = "script"
	a       = "a"
	nav     = "nav"
	span    = "span"
	section = "section"
	img     = "img"
	button  = "button"
	table   = "table"
	tr      = "tr"
	th      = "th"
	td      = "td"
	tbody   = "tbody"
	thead   = "thead"
)

type Attribute struct {
	Name, Val string
	IsBool    bool
}

func Attrs(attrs ...Attribute) []Attribute {
	return attrs
}

type Element struct {
	Tag           string
	Attributes    []Attribute
	Elements      []Element
	Text          string
	IsSelfClosing bool
}

func Els(els ...Element) []Element {
	return els
}

func basicTag(tag string, attrs []Attribute, elements []Element) Element {
	return Element{
		Tag:        tag,
		Elements:   elements,
		Attributes: attrs,
	}
}

func selfClosingTag(tag string, attrs []Attribute) Element {
	return Element{
		Tag:           tag,
		Attributes:    attrs,
		IsSelfClosing: true,
	}
}

func Div(attrs []Attribute, elements []Element) Element {
	return basicTag(div, attrs, elements)
}

func P(attrs []Attribute, elements []Element) Element {
	return basicTag(p, attrs, elements)
}

func A(attrs []Attribute, elements []Element) Element {
	return basicTag(a, attrs, elements)
}

func Span(attrs []Attribute, elements []Element) Element {
	return basicTag(span, attrs, elements)
}

func Section(attrs []Attribute, elements []Element) Element {
	return basicTag(section, attrs, elements)
}

func Img(attrs []Attribute) Element {
	return selfClosingTag(img, attrs)
}

func Nav(attrs []Attribute, elements []Element) Element {
	return basicTag(nav, attrs, elements)
}

func Html(attrs []Attribute, elements []Element) Element {
	return basicTag(html, attrs, elements)
}

func Head(attrs []Attribute, elements []Element) Element {
	return basicTag(head, attrs, elements)
}

func Title(attrs []Attribute, elements []Element) Element {
	return basicTag(title, attrs, elements)
}

func Meta(attrs []Attribute) Element {
	return selfClosingTag(meta, attrs)
}

func Body(attrs []Attribute, elements []Element) Element {
	return basicTag(body, attrs, elements)
}

func Link(attrs []Attribute) Element {
	return selfClosingTag(link, attrs)
}

func Script(attrs []Attribute, elements []Element) Element {
	return basicTag(script, attrs, elements)
}

func Text(i interface{}) Element {
	return Element{
		Tag:  "text",
		Text: fmt.Sprintf("%v", i),
	}
}

func Button(attrs []Attribute, elements []Element) Element {
	return basicTag(button, attrs, elements)
}

func Table(attrs []Attribute, elements []Element) Element {
	return basicTag(table, attrs, elements)
}

func THead(attrs []Attribute, elements []Element) Element {
	return basicTag(thead, attrs, elements)
}

func TBody(attrs []Attribute, elements []Element) Element {
	return basicTag(tbody, attrs, elements)
}

func Tr(attrs []Attribute, elements []Element) Element {
	return basicTag(tr, attrs, elements)
}

func Th(attrs []Attribute, elements []Element) Element {
	return basicTag(th, attrs, elements)
}

func Td(attrs []Attribute, elements []Element) Element {
	return basicTag(td, attrs, elements)
}
