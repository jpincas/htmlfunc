package htmlfunc

import (
	"fmt"
	"strings"

	"github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

var Render = renderElement

func Write(el h.Element) []byte {
	renderedElement, _ := renderElement(el, 0)
	return []byte(renderedElement)
}

func WriteDoc(el h.Element) []byte {
	return []byte(RenderDoc(el))
}

func WriteDocWith(docOptions string, el h.Element) []byte {
	return []byte(RenderDocWith(docOptions, el))
}

func RenderNoDoc(el h.Element) string {
	renderedElement, _ := renderElement(el, 0)
	return renderedElement
}

func RenderDoc(el h.Element) string {
	renderedElement, _ := renderElement(el, 0)
	return fmt.Sprintf("<!DOCTYPE html>\n%s", renderedElement)
}

func RenderDocWith(docOptions string, el h.Element) string {
	renderedElement, _ := renderElement(el, 0)
	return fmt.Sprintf("<!DOCTYPE html %s>\n%s", docOptions, renderedElement)
}

func renderElements(els []h.Element, tabs int) string {
	var renderedEls []string

	for _, el := range els {
		if renderedElement, doRender := renderElement(el, tabs); doRender {
			renderedEls = append(
				renderedEls,
				renderedElement,
			)
		}
	}

	return strings.Join(renderedEls, "")
}

func renderNTabs(n int) (res string) {
	for i := 0; i < n; i++ {
		res = res + "  "
	}

	return
}

func renderElement(el h.Element, tabs int) (string, bool) {
	// Raw trumps everything and is just returned as is
	if el.Raw != "" {
		return el.Raw, true
	}

	// An element without a tag can be strategically used to render 'nothing'
	if el.Tag == "" {
		return "", false
	}

	t := renderNTabs(tabs)

	if el.Tag == "text" {
		return fmt.Sprintf("%s%s\n", t, el.Text), true
	}

	if el.IsSelfClosing {
		return fmt.Sprintf(
			"%s<%s%s>\n",
			t,
			el.Tag,
			renderAttrs(el.Attributes),
		), true
	}

	return fmt.Sprintf(
		"%s<%s%s>\n%s%s</%s>\n",
		t,
		el.Tag,
		renderAttrs(el.Attributes),
		renderElements(el.Elements, tabs+1),
		t,
		el.Tag,
	), true
}

func renderAttrs(attrs attributes.Attributes) string {
	attrs = attrs.MergeDuplicates()

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

func renderAttr(attr attributes.Attribute) string {
	// An attr without a name is not rendered
	if attr.Name == "" {
		return ""
	}

	if attr.IsBool {
		return attr.Name
	}

	return fmt.Sprintf(`%s="%s"`, attr.Name, attr.Val)
}
