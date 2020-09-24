package htmlfunc

import (
	"fmt"
	"strings"

	h "github.com/jpincas/htmlfunc/html"
)

var Render = renderElement

func Write(el h.Element) []byte {
	renderedElement, _ := renderElement(el)
	return []byte(renderedElement)
}

func WriteDoc(el h.Element) []byte {
	return []byte(RenderDoc(el))
}

func RenderDoc(el h.Element) string {
	renderedElement, _ := renderElement(el)
	return fmt.Sprintf(`<!DOCTYPE html>%s`, renderedElement)
}

func RenderDocWith(docOptions string, el h.Element) string {
	renderedElement, _ := renderElement(el)
	return fmt.Sprintf(`<!DOCTYPE html %s>%s`, docOptions, renderedElement)
}

func renderElements(els []h.Element) string {
	var renderedEls []string

	for _, el := range els {
		if renderedElement, doRender := renderElement(el); doRender {
			renderedEls = append(
				renderedEls,
				renderedElement,
			)
		}
	}

	return strings.Join(renderedEls, "")
}

func renderElement(el h.Element) (string, bool) {
	// An element without a tag can be strategically used to render 'nothing'
	if el.Tag == "" {
		return "", false
	}

	if el.Tag == "text" {
		return el.Text, true
	}

	if el.IsSelfClosing {
		return fmt.Sprintf(
			"<%s%s>",
			el.Tag,
			renderAttrs(el.Attributes),
		), true
	}

	return fmt.Sprintf(
		"<%s%s>%s</%s>",
		el.Tag,
		renderAttrs(el.Attributes),
		renderElements(el.Elements),
		el.Tag,
	), true
}

func renderAttrs(attrs []h.Attribute) string {
	if len(attrs) == 0 {
		return ""
	}

	var renderedAttrs []string

	for _, attr := range attrs {
		renderedAttrs = append(
			renderedAttrs,
			renderAttr(attr),
		)
	}

	return " " + strings.Join(renderedAttrs, " ")
}

func renderAttr(attr h.Attribute) string {
	if attr.IsBool {
		return attr.Name
	}

	return fmt.Sprintf(`%s="%s"`, attr.Name, attr.Val)
}
