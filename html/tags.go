package html

import (
	"fmt"
	"strings"
)

type Attribute struct {
	Name, Val string
}

func Attributes(attrs ...Attribute) []Attribute {
	return attrs
}

type Element struct {
	Tag        string
	Attributes []Attribute
	Elements   []Element
	Text       string
}

func Els(els ...Element) []Element {
	return els
}

func renderElements(els []Element) string {
	var renderedEls []string

	for _, el := range els {
		renderedEls = append(
			renderedEls,
			renderElement(el),
		)
	}

	return strings.Join(renderedEls, "")
}

var Render = renderElement

func renderElement(el Element) string {
	if el.Tag == "text" {
		return el.Text
	}

	return fmt.Sprintf(
		"<%s %s>%s</%s>",
		el.Tag,
		renderAttrs(el.Attributes),
		renderElements(el.Elements),
		el.Tag,
	)
}

func renderAttrs(attrs []Attribute) string {
	var renderedAttrs []string

	for _, attr := range attrs {
		renderedAttrs = append(
			renderedAttrs,
			fmt.Sprintf(`%s="%s"`, attr.Name, attr.Val),
		)
	}

	return strings.Join(renderedAttrs, " ")
}

func Div(attrs []Attribute, elements []Element) Element {
	return Element{
		Tag:        "div",
		Elements:   elements,
		Attributes: attrs,
	}
}

func P(attrs []Attribute, elements []Element) Element {
	return Element{
		Tag:        "p",
		Elements:   elements,
		Attributes: attrs,
	}
}

func Text(s string) Element {
	return Element{
		Tag:  "text",
		Text: s,
	}
}
