package attributes

import "github.com/jpincas/htmlfunc/html"

func Class(s string) html.Attribute {
	return html.Attribute{"class", s}
}

func Id(s string) html.Attribute {
	return html.Attribute{"id", s}
}
