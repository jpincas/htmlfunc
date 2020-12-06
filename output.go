package htmlfunc

import (
	"fmt"

	h "github.com/jpincas/htmlfunc/html"
)

func Write(el h.Element) []byte {
	renderedElement, _ := el.Output(0)
	return []byte(renderedElement)
}

func WriteDoc(el h.Element) []byte {
	return []byte(RenderDoc(el))
}

func WriteDocWith(docOptions string, el h.Element) []byte {
	return []byte(RenderDocWith(docOptions, el))
}

func RenderNoDoc(el h.Element) string {
	renderedElement, _ := el.Output(0)
	return renderedElement
}

func RenderDoc(el h.Element) string {
	renderedElement, _ := el.Output(0)
	return fmt.Sprintf("<!DOCTYPE html>\n%s", renderedElement)
}

func RenderDocWith(docOptions string, el h.Element) string {
	renderedElement, _ := el.Output(0)
	return fmt.Sprintf("<!DOCTYPE html %s>\n%s", docOptions, renderedElement)
}
