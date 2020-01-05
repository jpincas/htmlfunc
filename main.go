package main

import (
	"fmt"

	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

func main() {

	result := h.Div(
		h.Attributes(
			a.Id("test-id"),
			a.Class("test-class"),
		),
		h.Els(
			h.P(
				h.Attributes(
					a.Id("nested-id"),
					a.Class("nested-class"),
				),
				h.Els(
					h.Text("Here is some content"),
				),
			),
		),
	)

	fmt.Println(h.Render(result))
}
