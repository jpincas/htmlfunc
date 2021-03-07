package attributes

import (
	"fmt"
	"strings"

	"github.com/jpincas/htmlfunc/css"
)

const (
	id           = "id"
	class        = "class"
	title        = "title"
	lang         = "lang"
	charset      = "charset"
	name         = "name"
	content      = "content"
	httpEquiv    = "http-equiv"
	rel          = "rel"
	href         = "href"
	type_        = "type"
	defer_       = "defer"
	src          = "src"
	alt          = "alt"
	as           = "as"
	target       = "target"
	role         = "role"
	width        = "width"
	height       = "height"
	onclick      = "onclick"
	onload       = "onload"
	xmlns        = "xmlns"
	style        = "style"
	sizes        = "sizes"
	async        = "async"
	disabled     = "disabled"
	checked      = "checked"
	property     = "property"
	placeholder  = "placeholder"
	autocomplete = "autocomplete"
	media        = "media"
	integrity    = "integrity"
	crossOrigin  = "crossorigin"

	cellPadding = "cellpadding"
	cellSpacing = "cellspacing"
	border      = "border"
	align       = "align"
	vAlign      = "valign"

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

	colSpan = "colspan"

	method   = "method"
	action   = "action"
	pattern  = "pattern"
	required = "required"
	value    = "value"
)

const (
	Round = "round"
	Blank = "_blank"
)

func Custom(tagName, s string) Attribute {
	return regularAttribute(tagName, s)
}

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

func Style1(style css.KeyValuePair, styles ...css.KeyValuePair) Attribute {
	ss := append(styles, style)
	return Style(ss...)
}

func Class(s string) Attribute {
	return regularAttribute(class, s)
}

// ClassIf is a shortcut for Class().RenderIf()
func ClassIf(s string, b bool) Attribute {
	return regularAttribute(class, s).RenderIf(b)
}

func Title(s string) Attribute {
	return regularAttribute(title, s)
}

func Target(s string) Attribute {
	return regularAttribute(target, s)
}

func Sizes(s string) Attribute {
	return regularAttribute(sizes, s)
}

func Property(s string) Attribute {
	return regularAttribute(property, s)
}

func Placeholder(s string) Attribute {
	return regularAttribute(placeholder, s)
}

func Method(s string) Attribute {
	return regularAttribute(method, s)
}

func Action(s string) Attribute {
	return regularAttribute(action, s)
}

func Pattern(s string) Attribute {
	return regularAttribute(pattern, s)
}

func Value(s string) Attribute {
	return regularAttribute(value, s)
}

func Data(suffix, s string) Attribute {
	return regularAttribute(fmt.Sprintf("data-%s", suffix), s)
}

func Autocomplete(s string) Attribute {
	return regularAttribute(autocomplete, s)
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

func Integrity(s string) Attribute {
	return regularAttribute(integrity, s)
}

func CrossOrigin(s string) Attribute {
	return regularAttribute(crossOrigin, s)
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

func Alt(s string) Attribute {
	return regularAttribute(alt, s)
}

func As(s string) Attribute {
	return regularAttribute(as, s)
}

func Media(s string) Attribute {
	return regularAttribute(media, s)
}

func OnLoad(s string) Attribute {
	return regularAttribute(onload, s)
}

func Defer(b bool) Attribute {
	return booleanAttribute(defer_, b)
}

func Async(b bool) Attribute {
	return booleanAttribute(async, b)
}

func Required(b bool) Attribute {
	return booleanAttribute(required, b)
}

func Disabled(b bool) Attribute {
	return booleanAttribute(disabled, b)
}

func Checked(b bool) Attribute {
	return booleanAttribute(checked, b)
}

func Role(s string) Attribute {
	return regularAttribute(role, s)
}

func Width(s string) Attribute {
	return regularAttribute(width, s)
}

func Height(s string) Attribute {
	return regularAttribute(height, s)
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

func VAlign(s string) Attribute {
	return regularAttribute(vAlign, s)
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

func ColSpan(i int) Attribute {
	return intAttribute(colSpan, i)
}
