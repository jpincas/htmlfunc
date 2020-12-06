package attributes

import (
	"strings"

	"github.com/jpincas/htmlfunc/css"
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

	viewBox        = "viewBox" //yes, camel case is correct
	stroke         = "stroke"
	strokeWidth    = "stroke-width"
	strokeLineCap  = "stroke-linecap"
	strokeLineJoin = "stroke-linejoin"
	fill           = "fill"
	d              = "d"
	points         = "points"

	x  = "x"
	rx = "rx"
	y  = "y"
	ry = "ry"
	x1 = "x1"
	y1 = "y1"
	x2 = "x2"
	y2 = "y2"

	cx = "cx"
	cy = "cy"
	r  = "r"
)

const (
	Round = "round"
)

func Style(styles ...css.KeyValuePair) Attribute {

	includedStyles := []css.KeyValuePair{}
	for _, style := range styles {
		if style.Include {
			includedStyles = append(includedStyles, style)
		}
	}

	if len(includedStyles) == 0 {
		return Attribute{}
	}

	return regularAttribute(style, css.PrintStyles(includedStyles))
}

func Class(s string) Attribute {
	return regularAttribute(class, s)
}

func Title(s string) Attribute {
	return regularAttribute(title, s)
}

// ClassesIf takes a list of classes to apply according to a corresponding list
// of booleans.  If more classes than appliers are provided, then extra
// classes are automatically applied, which is a convenient way to provide
// unconditional classes to the function.
func ClassesIf(classes []string, appliers []bool) Attribute {
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

func Classes(classes []string) Attribute {
	return Class(strings.Join(classes, " "))
}

func Id(s string) Attribute {
	return regularAttribute(id, s)
}

func Xmlns(s string) Attribute {
	return regularAttribute(xmlns, s)
}

func Lang(s string) Attribute {
	return regularAttribute(lang, s)
}

func Charset(s string) Attribute {
	return regularAttribute(charset, s)
}

func Name(s string) Attribute {
	return regularAttribute(name, s)
}

func Content(s string) Attribute {
	return regularAttribute(content, s)
}

func HttpEquiv(s string) Attribute {
	return regularAttribute(httpEquiv, s)
}

func Rel(s string) Attribute {
	return regularAttribute(rel, s)
}

func Href(s string) Attribute {
	return regularAttribute(href, s)
}

func Type(s string) Attribute {
	return regularAttribute(type_, s)
}

func Src(s string) Attribute {
	return regularAttribute(src, s)
}

func Defer() Attribute {
	return booleanAttribute(defer_)
}

func Role(s string) Attribute {
	return regularAttribute(role, s)
}

func Width(i int) Attribute {
	return intAttribute(width, i)
}

func Height(i int) Attribute {
	return intAttribute(height, i)
}

func OnClick(s string) Attribute {
	return regularAttribute(onclick, s)
}

func CellPadding(pixels int) Attribute {
	return intAttribute(cellPadding, pixels)
}

func CellSpacing(pixels int) Attribute {
	return intAttribute(cellSpacing, pixels)
}

func Border(pixels int) Attribute {
	return intAttribute(border, pixels)
}

func Align(s string) Attribute {
	return regularAttribute(align, s)
}

func ViewBox(s string) Attribute {
	return regularAttribute(viewBox, s)
}

func Stroke(s string) Attribute {
	return regularAttribute(stroke, s)
}

func Fill(s string) Attribute {
	return regularAttribute(fill, s)
}

func StrokeLineCap(s string) Attribute {
	return regularAttribute(strokeLineCap, s)
}

func StrokeLineJoin(s string) Attribute {
	return regularAttribute(strokeLineJoin, s)
}

func StrokeWidth(i int) Attribute {
	return intAttribute(strokeWidth, i)
}

func D(s string) Attribute {
	return regularAttribute(d, s)
}

func Points(s string) Attribute {
	return regularAttribute(points, s)
}

func X(i int) Attribute {
	return intAttribute(x, i)
}

func RX(i int) Attribute {
	return intAttribute(rx, i)
}

func X1(f float64) Attribute {
	return floatAttribute(x1, f)
}

func X2(f float64) Attribute {
	return floatAttribute(x2, f)
}

func Y(i int) Attribute {
	return intAttribute(y, i)
}

func RY(i int) Attribute {
	return intAttribute(ry, i)
}

func Y1(f float64) Attribute {
	return floatAttribute(y1, f)
}

func Y2(f float64) Attribute {
	return floatAttribute(y2, f)
}

func CX(i int) Attribute {
	return intAttribute(cx, i)
}

func CY(i int) Attribute {
	return intAttribute(cy, i)
}

func R(i int) Attribute {
	return intAttribute(r, i)
}
