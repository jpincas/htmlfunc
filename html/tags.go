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
)

type Attribute struct {
	Name, Val string
	IsBool    bool
}

type Attributes []Attribute

func (a Attribute) RenderIf(doRender bool) Attribute {
	if doRender {
		return a
	}

	return Attribute{}
}

func Attrs(attrs ...Attribute) Attributes {
	return attrs
}

// Merge helps deal with repition of certain attributes, merging them into one.
// At the moment, we deal with 'class' and 'style'
func (attrs Attributes) Merge() Attributes {
	var classes []string
	var styles []string
	var otherAttrs Attributes

	for _, attr := range atts {
		switch attr.Name {
		case class:
			classes = append(classes, attr.Val)
		case style:
			styles = append(styles, attr.Val)
		default:
			otherAttrs = append(otherAttrs, attr)
		}
	}

	return Attributes{}
}

type Element struct {
	// Raw is basically a bypass, allowing for element creation from raw HTML
	Raw string

	Tag           string
	Attributes    Attributes
	Elements      []Element
	Text          string
	IsSelfClosing bool
}

func (el Element) RenderIf(doRender bool) Element {
	if doRender {
		return el
	}

	return Element{}
}

func (el Element) RenderIfWithDefault(doRender bool, d Element) Element {
	if doRender {
		return el
	}

	return d
}

func (el Element) IsEmpty() bool {
	return el.Text == "" && len(el.Elements) == 0
}

func Els(els ...Element) []Element {
	return els
}

//Els1 is a convenience function to combine a single element then a list (which you can't do with
// the standard spread syntax)
func Els1(el Element, els []Element) []Element {
	return append(Els(el), els...)
}

//Els2 is a convenience function to combine two single elements, then a list (which you can't do with
// the standard spread syntax)
func Els2(el1, el2 Element, els []Element) []Element {
	return append(Els(el1, el2), els...)
}

func basicTag(tag string, attrs Attributes, elements []Element) Element {
	return Element{
		Tag:        tag,
		Elements:   elements,
		Attributes: attrs,
	}
}

func selfClosingTag(tag string, attrs Attributes) Element {
	return Element{
		Tag:           tag,
		Attributes:    attrs,
		IsSelfClosing: true,
	}
}

func RawElement(b []byte) Element {
	return Element{
		Raw: string(b),
	}
}

func Div(attrs Attributes, elements ...Element) Element {
	return basicTag(div, attrs, elements)
}

func P(attrs Attributes, elements ...Element) Element {
	return basicTag(p, attrs, elements)
}

func A(attrs Attributes, elements ...Element) Element {
	return basicTag(a, attrs, elements)
}

func Span(attrs Attributes, elements ...Element) Element {
	return basicTag(span, attrs, elements)
}

func Section(attrs Attributes, elements ...Element) Element {
	return basicTag(section, attrs, elements)
}

func Img(attrs Attributes) Element {
	return selfClosingTag(img, attrs)
}

func SVG(attrs Attributes, elements ...Element) Element {
	return basicTag(svg, attrs, elements)
}

func Path(attrs Attributes) Element {
	return basicTag(path, attrs, []Element{})
}

func Polygon(attrs Attributes) Element {
	return basicTag(polygon, attrs, []Element{})
}

func Rect(attrs Attributes) Element {
	return basicTag(rect, attrs, []Element{})
}

func Line(attrs Attributes) Element {
	return basicTag(line, attrs, []Element{})
}

func Circle(attrs Attributes) Element {
	return basicTag(circle, attrs, []Element{})
}

func Br(_ Attributes) Element {
	return selfClosingTag(br, []Attribute{})
}

func Nav(attrs Attributes, elements ...Element) Element {
	return basicTag(nav, attrs, elements)
}

func Html(attrs Attributes, elements ...Element) Element {
	return basicTag(html, attrs, elements)
}

func Head(attrs Attributes, elements ...Element) Element {
	return basicTag(head, attrs, elements)
}

func Title(attrs Attributes, elements ...Element) Element {
	return basicTag(title, attrs, elements)
}

func Meta(attrs Attributes) Element {
	return selfClosingTag(meta, attrs)
}

func Body(attrs Attributes, elements ...Element) Element {
	return basicTag(body, attrs, elements)
}

func Link(attrs Attributes) Element {
	return selfClosingTag(link, attrs)
}

func Script(attrs Attributes, elements ...Element) Element {
	return basicTag(script, attrs, elements)
}

func Text(i interface{}) Element {
	return Element{
		Tag:  "text",
		Text: fmt.Sprintf("%v", i),
	}
}

// Nothing generates a blank element.  The only reason we have the arguments
// is to make the funciton type signature the same as the other construction
// functions
func Nothing(_ Attributes, _ ...Element) Element {
	return Element{}
}

func Button(attrs Attributes, elements ...Element) Element {
	return basicTag(button, attrs, elements)
}

func Table(attrs Attributes, elements ...Element) Element {
	return basicTag(table, attrs, elements)
}

func THead(attrs Attributes, elements ...Element) Element {
	return basicTag(thead, attrs, elements)
}

func TBody(attrs Attributes, elements ...Element) Element {
	return basicTag(tbody, attrs, elements)
}

func Tr(attrs Attributes, elements ...Element) Element {
	return basicTag(tr, attrs, elements)
}

func Th(attrs Attributes, elements ...Element) Element {
	return basicTag(th, attrs, elements)
}

func Td(attrs Attributes, elements ...Element) Element {
	return basicTag(td, attrs, elements)
}

func H1(attrs Attributes, elements ...Element) Element {
	return basicTag(h1, attrs, elements)
}

func H2(attrs Attributes, elements ...Element) Element {
	return basicTag(h2, attrs, elements)
}

func H3(attrs Attributes, elements ...Element) Element {
	return basicTag(h3, attrs, elements)
}

func H4(attrs Attributes, elements ...Element) Element {
	return basicTag(h4, attrs, elements)
}

func H5(attrs Attributes, elements ...Element) Element {
	return basicTag(h5, attrs, elements)
}

func H6(attrs Attributes, elements ...Element) Element {
	return basicTag(h6, attrs, elements)
}
func Ul(attrs Attributes, elements ...Element) Element {
	return basicTag(ul, attrs, elements)
}
func Ol(attrs Attributes, elements ...Element) Element {
	return basicTag(ol, attrs, elements)
}
func Li(attrs Attributes, elements ...Element) Element {
	return basicTag(li, attrs, elements)
}
