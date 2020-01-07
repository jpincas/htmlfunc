package htmlfunc

import (
	"fmt"
	"strings"

	h "github.com/jpincas/htmlfunc/html"
)

var Render = renderElement

func Write(el h.Element) []byte {
	return []byte(renderElement(el))
}

func WriteDoc(el h.Element) []byte {
	return []byte(RenderDoc(el))
}

func RenderDoc(el h.Element) string {
	return fmt.Sprintf(`<!DOCTYPE html>%s`, renderElement(el))
}

func renderElements(els []h.Element) string {
	var renderedEls []string

	for _, el := range els {
		renderedEls = append(
			renderedEls,
			renderElement(el),
		)
	}

	return strings.Join(renderedEls, "")
}

func renderElement(el h.Element) string {
	if el.Tag == "text" {
		return el.Text
	}

	if el.IsSelfClosing {
		return fmt.Sprintf(
			"<%s%s>",
			el.Tag,
			renderAttrs(el.Attributes),
		)
	}

	return fmt.Sprintf(
		"<%s%s>%s</%s>",
		el.Tag,
		renderAttrs(el.Attributes),
		renderElements(el.Elements),
		el.Tag,
	)
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
