package test

import (
	"os"
	"testing"

	a "github.com/jpincas/htmlfunc/attributes"
	"github.com/jpincas/htmlfunc/css"
	"github.com/jpincas/htmlfunc/html"
	h "github.com/jpincas/htmlfunc/html"
)

// The only test we have is to generate a sample HTML page that uses as much of the library as possible
// Obviously, the test serves to make sure the code compiles and runs, but you can also:
// - visually inspect the HTML to test anything you are working on
// - run it against an HTML validator. e.g., https://github.com/validator/validator/
// - run test, then: /usr/local/bin/vnu-runtime-image/bin/vnu index.html
// - inspect the generated HTML in a browser

func TestPageGen(t *testing.T) {
	content :=
		h.Div(
			a.Attrs(
				a.Id("test-id"),
				a.Class("test-class"),
			),
			h.P(
				a.Attrs(
					a.Id("nested-id"),
					a.Class("nested-class"),
				),
				h.Text("Here is some content"),
			),
		)

	head :=
		h.Head(
			a.Attrs(),

			// Title is required - validator will fail without it
			h.Title(a.Attrs(), html.Text("Test Document")),
			h.Meta(a.Attrs(a.Charset("UTF-8"))),
			h.Meta(a.Attrs(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0"))),
			h.Meta(a.Attrs(a.HttpEquiv("X-UA-Compatible"), a.Content("ie=edge"))),
		)

	table :=
		h.Table(
			a.Attrs(),
			h.Tr(
				a.Attrs(
					a.Class("class1"),
					a.Class("class2"),
				),
				h.Td(
					a.Attributes{
						a.Style(),
					},
				),
			),
		)

	body :=
		h.Body(
			a.Attrs(
				a.Class("body-class"),

				// Repeated 'style' declaration will fail if not dealt with
				a.Style(
					css.Width(css.WithUnits(50, css.Px)),
				),
				a.Style(
					css.Color("white", false),
					css.FontSize(css.WithUnits(150, css.Px), true),
				),
			),
			content,
			table,
		)

	doc :=
		h.Html(
			a.Attrs(a.Lang("en")),
			head,
			body,
		)

	f, err := os.Create("index.html")
	defer f.Close()

	if err != nil {
		t.Fatal("Could not create index.html")
	}

	if _, err := f.Write(doc.DocBytes()); err != nil {
		t.Fatal("Could not write to index.html: ", err)
	}
}
