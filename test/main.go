package main

import (
	"fmt"

	"github.com/jpincas/htmlfunc"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

func main() {

	content :=
		h.Div(
			h.Attrs(
				a.Id("test-id"),
				a.Class("test-class"),
			),
			h.Els(
				h.P(
					h.Attrs(
						a.Id("nested-id"),
						a.Class("nested-class"),
					),
					h.Els(
						h.Text("Here is some content"),
					),
				),
			),
		)

	head :=
		h.Head(
			h.Attrs(),
			h.Els(
				h.Meta(h.Attrs(a.Charset("UTF-8"))),
				h.Meta(h.Attrs(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0"))),
				h.Meta(h.Attrs(a.HttpEquiv("X-UA-Compatible"), a.Content("ie=edge"))),
			),
		)

	body :=
		h.Body(
			h.Attrs(a.Class("body-class")),
			h.Els(content),
		)

	doc :=
		h.Html(
			h.Attrs(a.Lang("en")),
			h.Els(head, body),
		)

	fmt.Println(htmlfunc.RenderDoc(doc))
}
