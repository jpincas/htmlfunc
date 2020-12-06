package test

import (
	"testing"

	"github.com/jpincas/htmlfunc"
	a "github.com/jpincas/htmlfunc/attributes"
	h "github.com/jpincas/htmlfunc/html"
)

func TestRenderDoc(t *testing.T) {
	el := h.Div(
		a.Attrs(
			a.Class("test"),
		),
		h.Text("TEST"),
	)

	output := htmlfunc.RenderDoc(el)
	expected := `<!DOCTYPE html>
<div class="test">
  TEST
</div>
`

	if output != expected {
		t.Errorf("Output was not as expected.  Output: \n%s\n Expected: \n%s\n", output, expected)
	}
}
