package html

import (
	"testing"

	"github.com/jpincas/htmlfunc/attributes"
	"github.com/jpincas/htmlfunc/css"
)

func TestOutput(t *testing.T) {
	el := Div(
		attributes.Attrs(
			attributes.Class("myclass"),
		),
		Span(
			attributes.Attrs(
				attributes.Class("two classes"),
				attributes.Class("duplicate-class-declaration"),
			),
			Text("SPAN"),
		),
		Div(
			attributes.Attrs(
				// Duplicate 'class' and 'style' declarations
				attributes.Class("should-not-appear").RenderIf(false),
				attributes.Class("should-appear").RenderIf(true),
				attributes.Style(css.Color("white"), css.FontWeight(css.Bold)),
				attributes.Style(css.FontSize(css.WithUnits(14, css.Px))),
			),
			H1(attributes.Attrs(), Text("H1")),
			H2(attributes.Attrs(), Text("H2")),
		),
		P(
			attributes.Attrs(),
			Text("this shouldn't render"),
		).RenderIf(false),
		P(
			attributes.Attrs(),
			Text("this shouldn't render"),
		).RenderIfWithDefault(false, Text("this should get rendered by default")),
		RawElement([]byte("<notarealtag>Hello, World</notarealtag>")),
	)

	output, _ := el.Output(0)
	expected := `<div class="myclass">
  <span class="two classes duplicate-class-declaration">
    SPAN
  </span>
  <div class="should-appear" style="color:white;font-weight:bold;font-size:14px">
    <h1>
      H1
    </h1>
    <h2>
      H2
    </h2>
  </div>
  this should get rendered by default
  <notarealtag>Hello, World</notarealtag>
</div>
`

	if output != expected {
		t.Errorf("Output was not as expected.  Output: \n%s\n Expected: \n%s\n", output, expected)
	}
}
