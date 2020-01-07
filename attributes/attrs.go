package attributes

import (
	"strconv"

	"github.com/jpincas/htmlfunc/html"
)

const (
	id        = "id"
	class     = "class"
	lang      = "lang"
	charset   = "charset"
	name      = "name"
	content   = "content"
	httpEquiv = "http-equiv"
	rel       = "rel"
	href      = "href"
	type_     = "type"
	defer_    = "defer"
	src       = "src"
	role      = "role"
	width     = "width"
	height    = "height"
	onclick   = "onclick"
)

func regularAttribute(k, v string) html.Attribute {
	return html.Attribute{
		Name:   k,
		Val:    v,
		IsBool: false,
	}
}

func booleanAttribute(k string) html.Attribute {
	return html.Attribute{
		Name:   k,
		IsBool: true,
	}
}

func intAttribute(k string, v int) html.Attribute {
	return html.Attribute{
		Name:   k,
		Val:    strconv.Itoa(v),
		IsBool: false,
	}
}

func Class(s string) html.Attribute {
	return regularAttribute(class, s)
}

func Id(s string) html.Attribute {
	return regularAttribute(id, s)
}

func Lang(s string) html.Attribute {
	return regularAttribute(lang, s)
}

func Charset(s string) html.Attribute {
	return regularAttribute(charset, s)
}

func Name(s string) html.Attribute {
	return regularAttribute(name, s)
}

func Content(s string) html.Attribute {
	return regularAttribute(content, s)
}

func HttpEquiv(s string) html.Attribute {
	return regularAttribute(httpEquiv, s)
}

func Rel(s string) html.Attribute {
	return regularAttribute(rel, s)
}

func Href(s string) html.Attribute {
	return regularAttribute(href, s)
}

func Type(s string) html.Attribute {
	return regularAttribute(type_, s)
}

func Src(s string) html.Attribute {
	return regularAttribute(src, s)
}

func Defer() html.Attribute {
	return booleanAttribute(defer_)
}

func Role(s string) html.Attribute {
	return regularAttribute(role, s)
}

func Width(i int) html.Attribute {
	return intAttribute(width, i)
}

func Height(i int) html.Attribute {
	return intAttribute(height, i)
}

func OnClick(s string) html.Attribute {
	return regularAttribute(onclick, s)
}
