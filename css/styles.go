package css

import (
	"fmt"
	"strings"
)

const (
	backgroundColor = "background-color"
	backgroundImage = "background-image"

	backgroundRepeat   = "background-repeat"
	backgroundPosition = "background-position"
	backgroundSize     = "background-size"

	color   = "color"
	display = "display"

	fontFamily = "font-family"
	fontSize   = "font-size"
	fontWeight = "font-weight"

	width    = "width"
	height   = "height"
	maxWidth = "max-width"

	margin       = "margin"
	marginRight  = "margin-right"
	marginLeft   = "margin-left"
	marginTop    = "margin-top"
	marginBottom = "margin-bottom"

	padding       = "padding"
	paddingRight  = "padding-right"
	paddingLeft   = "padding-left"
	paddingTop    = "padding-top"
	paddingBottom = "padding-bottom"

	justifyContent = "justify-content"
	alignItems     = "align-items"
	textAlign      = "text-align"
	textDecoration = "text-decoration"
	flexWrap       = "flex-wrap"
	lineHeight     = "line-height"

	borderColor    = "border-color"
	borderTop      = "border-top"
	borderBottom   = "border-bottom"
	borderCollapse = "border-collapse"
	borderRadius   = "border-radius"
	borderStyle    = "border-style"
	border         = "border"

	overFlowX     = "overflow-x"
	opacity       = "opacity"
	listStyle     = "list-style"
	verticalAlign = "vertical-align"
	tableLayout   = "table-layout"

	pageBreakBefore = "page-break-before"
)

const (
	Zero = "0"
	None = "none"

	// CSS Units
	Px      = "px"
	Pt      = "pt"
	Percent = "%"
	Em      = "em"
	Cm      = "cm"
	Mm      = "mm"
	Rem     = "rem"

	Bold = "bold"

	Block        = "block"
	Inline       = "inline"
	Auto         = "auto"
	Flex         = "flex"
	Center       = "center"
	Right        = "right"
	Left         = "left"
	SpaceBetween = "space-between"
	Wrap         = "wrap"
	Always       = "always"

	Collapse = "collapse"

	NoRepeat    = "no-repeat"
	Cover       = "cover"
	Solid       = "solid"
	Transparent = "transparent"

	Baseline   = "baseline"
	Sub        = "sub"
	Super      = "super"
	Top        = "top"
	TextTop    = "text-top"
	Middle     = "middle"
	Bottom     = "bottom"
	TextBottom = "text-bottom"
	Initial    = "initial"
	Inherit    = "inherit"

	Fixed = "fixed"
)

var (
	// Font Family Stack
	FontFamilyHelvetica  = Fonts([]string{"Helvetica Neue", "Helvetica", "Arial", "sans-serif"})
	FontFamilyCourierNew = Fonts([]string{"Courier New", "monospace"})
)

type KeyValuePair struct {
	key     string
	value   string
	Include bool
}

type Styles []KeyValuePair

func StyleList(kvs ...KeyValuePair) Styles {
	return kvs
}

func (kvp KeyValuePair) String() string {
	return fmt.Sprintf("%s:%s", kvp.key, kvp.value)
}

func PrintStyles(styles []KeyValuePair) string {
	components := []string{}

	for _, kvp := range styles {
		components = append(components, kvp.String())
	}

	return strings.Join(components, ";")
}

func WithUnits(value interface{}, unit string) string {
	return fmt.Sprintf("%v%s", value, unit)
}

func NoUnits(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func constructKeyValuePair(key, value string, include ...bool) KeyValuePair {
	i := true
	if len(include) > 0 {
		i = include[0]
	}

	return KeyValuePair{key, value, i}
}

func Width(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(width, s, include...)
}

func Height(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(height, s, include...)
}

func MaxWidth(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(maxWidth, s, include...)
}

func Margin(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(margin, s, include...)
}

func MarginVertHoriz(vertical, horiztonal string, include ...bool) KeyValuePair {
	s := []string{vertical, horiztonal}
	return constructKeyValuePair(margin, strings.Join(s, " "), include...)
}

func MarginRight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(marginRight, s, include...)
}

func MarginLeft(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(marginLeft, s, include...)
}

func MarginTop(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(marginTop, s, include...)
}

func MarginBottom(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(marginBottom, s, include...)
}

func BackgroundColor(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundColor, s, include...)
}

func BackgroundImage(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundImage, s, include...)
}

func Color(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(color, s, include...)
}

func RGBA(r, g, b, a float64) string {
	return fmt.Sprintf("rgba(%v,%v,%v,%v)", r, g, b, a)
}

func RGB(r, g, b float64) string {
	return fmt.Sprintf("rgb(%v,%v,%v)", r, g, b)
}

func URL(s string) string {
	return fmt.Sprintf(`url('%s')`, s)
}

func Display(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(display, s, include...)
}

func Fonts(fonts []string) string {
	return strings.Join(fonts, ", ")
}

func FontFamily(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(fontFamily, s, include...)
}

func FontSize(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(fontSize, s, include...)
}

func FontWeight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(fontWeight, s, include...)
}

func JustifyContent(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(justifyContent, s, include...)
}

func MultipleArgs(ss ...string) string {
	return strings.Join(ss, " ")
}

func MultipleValues(ss ...string) string {
	return strings.Join(ss, ",")
}

func Padding(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(padding, s, include...)
}

func PaddingTop(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(paddingTop, s, include...)
}

func PaddingBottom(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(paddingBottom, s, include...)
}

func PaddingLeft(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(paddingLeft, s, include...)
}

func PaddingRight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(paddingRight, s, include...)
}

func PaddingVertHoriz(vertical, horiztonal string, include ...bool) KeyValuePair {
	s := []string{vertical, horiztonal}
	return constructKeyValuePair(padding, strings.Join(s, " "), include...)
}

func FlexWrap(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(flexWrap, s, include...)
}

func LineHeight(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(lineHeight, s, include...)
}

func AlignItems(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(alignItems, s, include...)
}

func BorderCollapse(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderCollapse, s, include...)
}

func BorderColor(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderColor, s, include...)
}

func BorderRadius(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderRadius, s, include...)
}

func BackgroundRepeat(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundRepeat, s, include...)
}

func BackgroundPosition(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundPosition, s, include...)
}

func BackgroundSize(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(backgroundSize, s, include...)
}

func OverflowX(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(overFlowX, s, include...)
}

func Opacity(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(opacity, s, include...)
}

func Border(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(border, s, include...)
}

func BorderTop(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderTop, s, include...)
}

func BorderBottom(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderBottom, s, include...)
}

func BorderStyle(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderStyle, s, include...)
}

func ListStyle(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(listStyle, s, include...)
}

func TextAlign(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(textAlign, s, include...)
}

func TextDecoration(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(textDecoration, s, include...)
}

func PageBreakBefore(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(pageBreakBefore, s, include...)
}

func VerticalAlign(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(verticalAlign, s, include...)
}

func TableLayout(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(tableLayout, s, include...)
}
