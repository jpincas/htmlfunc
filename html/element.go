package html

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/jpincas/htmlfunc/attributes"
)

// Element

type Element struct {
	// Raw is basically a bypass, allowing for element creation from raw HTML
	Raw string

	Tag           string
	Attributes    attributes.Attributes
	Elements      Elements
	Text          string
	IsSelfClosing bool
}

// Output prints an element to a string with an optional indentation level
func (el Element) Output(tabs int) (string, bool) {
	t := renderNTabs(tabs)

	// Raw trumps everything and is just returned as is
	if el.Raw != "" {
		return fmt.Sprintf("%s%s\n", t, el.Raw), true
	}

	// An element without a tag can be strategically used to render 'nothing'
	if el.Tag == "" {
		return "", false
	}

	if el.Tag == text {
		return fmt.Sprintf("%s%s\n", t, template.HTMLEscapeString(el.Text)), true
	}

	if el.IsSelfClosing {
		return fmt.Sprintf(
			"%s<%s%s>\n",
			t,
			el.Tag,
			el.Attributes.Output(),
		), true
	}

	return fmt.Sprintf(
		"%s<%s%s>\n%s%s</%s>\n",
		t,
		el.Tag,
		el.Attributes.Output(),
		el.Elements.Output(tabs+1),
		t,
		el.Tag,
	), true
}

func (el Element) String() string {
	s, _ := el.Output(0)
	return s
}

func (el Element) Bytes() []byte {
	renderedElement, _ := el.Output(0)
	return []byte(renderedElement)
}

func (el Element) DocBytes() []byte {
	return []byte(el.DocString())
}

func (el Element) DocBytesWithOptions(docOptions string) []byte {
	return []byte(el.DocStringWithOptions(docOptions))
}

func (el Element) DocString() string {
	renderedElement, _ := el.Output(0)
	return fmt.Sprintf("<!DOCTYPE html>\n%s", renderedElement)
}

func (el Element) DocStringWithOptions(docOptions string) string {
	renderedElement, _ := el.Output(0)
	return fmt.Sprintf("<!DOCTYPE html %s>\n%s", docOptions, renderedElement)
}

func (el Element) WriteDoc(w io.Writer) error {
	_, err := w.Write(el.DocBytes())
	return err
}

func (el Element) WriteDocWithOptions(w io.Writer, docOptions string) error {
	_, err := w.Write(el.DocBytesWithOptions(docOptions))
	return err
}

// Elements

type Elements []Element

// Output prints  elements to a string with an optional indentation level
func (els Elements) Output(tabs int) string {
	var renderedEls []string

	for _, el := range els {
		if renderedElement, doRender := el.Output(tabs); doRender {
			renderedEls = append(
				renderedEls,
				renderedElement,
			)
		}
	}

	return strings.Join(renderedEls, "")
}

func (els Elements) String() string {
	return els.Output(0)
}

func renderNTabs(n int) (res string) {
	for i := 0; i < n; i++ {
		res = res + "  "
	}

	return
}

// FlattenElementList flattens a list of lists of elements into a single list
func FlattenElementsList(elementList []Elements) (res Elements) {
	for _, els := range elementList {
		for _, el := range els {
			res = append(res, el)
		}
	}

	return
}

// Helpers

// RawElement is a 'bypass' to create an element directly
func RawElement(b []byte) Element {
	return Element{
		Raw: string(b),
	}
}

// Els is a handy constructor to create a list of elements
func Els(els ...Element) Elements {
	return els
}

// RenderIf will only render the element if it satisfies the predicate
func (el Element) RenderIf(doRender bool) Element {
	if doRender {
		return el
	}

	return Element{}
}

// RenderIf will only render the element if it satisfies the predicate, otherwise it will render a default element
func (el Element) RenderIfWithDefault(doRender bool, d Element) Element {
	if doRender {
		return el
	}

	return d
}

// IsEmpty specifies whether this is a 'blank' element
func (el Element) IsEmpty() bool {
	return el.Text == "" && len(el.Elements) == 0
}

//Els1 is a convenience function to combine a single element then a list (which you can't do with
// the standard spread syntax)
func Els1(el Element, els Elements) Elements {
	return append(Els(el), els...)
}

//Els2 is a convenience function to combine two single elements, then a list (which you can't do with
// the standard spread syntax)
func Els2(el1, el2 Element, els Elements) Elements {
	return append(Els(el1, el2), els...)
}

// Tag Construction

func basicTag(tag string, attrs attributes.Attributes, elements Elements) Element {
	return Element{
		Tag:        tag,
		Elements:   elements,
		Attributes: attrs,
	}
}

func selfClosingTag(tag string, attrs attributes.Attributes) Element {
	return Element{
		Tag:           tag,
		Attributes:    attrs,
		IsSelfClosing: true,
	}
}
