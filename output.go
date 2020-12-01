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

func WriteDocWith(docOptions string, el h.Element) []byte {
	return []byte(RenderDocWith(docOptions, el))
}

func RenderNoDoc(el h.Element) string {
	renderedElement, _ := renderElement(el)
	return renderedElement
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
	// Raw trumps everything and is just returned as is
	if el.Raw != "" {
		return el.Raw, true
	}

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
	var renderedAttrs []string

	for _, attr := range attrs {
		r := renderAttr(attr)
		// Don't include blank attrs, to avoid extraneous spaces
		if r != "" {
			renderedAttrs = append(renderedAttrs, r)
		}
	}

	if len(renderedAttrs) == 0 {
		return ""
	}

	return " " + strings.Join(renderedAttrs, " ")
}

func renderAttr(attr h.Attribute) string {
	// An attr without a name is not rendered
	if attr.Name == "" {
		return ""
	}

	if attr.IsBool {
		return attr.Name
	}

	return fmt.Sprintf(`%s="%s"`, attr.Name, attr.Val)
}
