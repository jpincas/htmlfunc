package main

import (
	"fmt"

	"github.com/jpincas/htmlfunc"
	a "github.com/jpincas/htmlfunc/attributes"
	"github.com/jpincas/htmlfunc/css"
	h "github.com/jpincas/htmlfunc/html"
)

func main() {

	content :=
		h.Div(
			h.Attrs(
				a.Id("test-id"),
				a.Class("test-class"),
			),
			h.P(
				h.Attrs(
					a.Id("nested-id"),
					a.Class("nested-class"),
				),
				h.Text("Here is some content"),
			),
		)

	head :=
		h.Head(
			h.Attrs(),
			h.Meta(h.Attrs(a.Charset("UTF-8"))),
			h.Meta(h.Attrs(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0"))),
			h.Meta(h.Attrs(a.HttpEquiv("X-UA-Compatible"), a.Content("ie=edge"))),
		)

	body :=
		h.Body(
			h.Attrs(
				a.Class("body-class"),
				a.Style(
					css.Width(50, css.Px),
				),
                                a.Style(
                                        css.Width(90, css.Px, false),
                                        css.Width(150, css.Px, true),
                                ),
			),
			content,
		)

	doc :=
		h.Html(
			h.Attrs(a.Lang("en")),
			head,
			body,
		)

        output := htmlfunc.RenderDoc(doc)

	fmt.Println(output)
}
