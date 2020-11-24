package attributes

import (
	"strconv"
	"strings"

	"github.com/jpincas/htmlfunc/css"
	"github.com/jpincas/htmlfunc/html"
)

const (
	id        = "id"
	class     = "class"
	title     = "title"
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
	xmlns     = "xmlns"
	style     = "style"

	cellPadding = "cellpadding"
	cellSpacing = "cellspacing"
	border      = "border"
	align       = "align"
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

func Style(styles ...css.KeyValuePair) html.Attribute {

	includedStyles := []css.KeyValuePair{}
	for _, style := range styles {
		if style.Include {
			includedStyles = append(includedStyles, style)
		}
	}

	if len(includedStyles) == 0 {
		return html.Attribute{}
	}

	return regularAttribute(style, css.PrintStyles(includedStyles))
}

func Class(s string) html.Attribute {
	return regularAttribute(class, s)
}

func Title(s string) html.Attribute {
	return regularAttribute(title, s)
}

// ClassesIf takes a list of classes to apply according to a corresponding list
// of booleans.  If more classes than appliers are provided, then extra
// classes are automatically applied, which is a convenient way to provide
// unconditional classes to the function.
func ClassesIf(classes []string, appliers []bool) html.Attribute {
	var classesToApply []string

	for i, class := range classes {
		if i < len(appliers) {
			if appliers[i] {
				classesToApply = append(classesToApply, class)
			}
		} else {
			classesToApply = append(classesToApply, class)
		}
	}

	return Class(strings.Join(classesToApply, " "))
}

func Classes(classes []string) html.Attribute {
	return Class(strings.Join(classes, " "))
}

func Id(s string) html.Attribute {
	return regularAttribute(id, s)
}

func Xmlns(s string) html.Attribute {
	return regularAttribute(xmlns, s)
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

func CellPadding(pixels int) html.Attribute {
	return intAttribute(cellPadding, pixels)
}

func CellSpacing(pixels int) html.Attribute {
	return intAttribute(cellSpacing, pixels)
}

func Border(pixels int) html.Attribute {
	return intAttribute(border, pixels)
}

func Align(s string) html.Attribute {
	return regularAttribute(align, s)
}
