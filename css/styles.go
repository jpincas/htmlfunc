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

	margin         = "margin"
	marginRight    = "margin-right"
	marginLeft     = "margin-left"
	padding        = "padding"
	justifyContent = "justify-content"
	alignItems     = "align-items"
	textAlign      = "text-align"
	flexWrap       = "flex-wrap"
	lineHeight     = "line-height"

	borderColor    = "border-color"
	borderTop      = "border-top"
	borderBottom   = "border-bottom"
	borderCollapse = "border-collapse"
	overFlowX      = "overflow-x"
	opacity        = "opacity"
	listStyle      = "list-style"
)

const (
	Zero = "0"
	None = "none"

	// CSS Units
	Px      = "px"
	Percent = "%"
	Em      = "em"

	Bold = "bold"

	Block        = "block"
	Inline       = "inline"
	Auto         = "auto"
	Flex         = "flex"
	Center       = "center"
	SpaceBetween = "space-between"
	Wrap         = "wrap"

	Collapse = "collapse"

	NoRepeat = "no-repeat"
	Cover    = "cover"
)

var (
	// Font Family Stack
	FontFamilyHelvetica = Fonts([]string{"Helvetica Neue", "Helvetica", "Arial", "sans-serif"})
)

type KeyValuePair struct {
	key     string
	value   string
	Include bool
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

func MultipleArgs(ss []string) string {
	return strings.Join(ss, " ")
}

func MultipleValues(ss []string) string {
	return strings.Join(ss, ",")
}

func Padding(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(padding, s, include...)
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

func BorderTop(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderTop, s, include...)
}

func BorderBottom(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(borderBottom, s, include...)
}

func ListStyle(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(listStyle, s, include...)
}

func TextAlign(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(textAlign, s, include...)
}
