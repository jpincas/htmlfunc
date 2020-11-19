package css

import (
	"fmt"
	"strings"
)

const (
	backgroundColor = "background-color"
	color           = "color"
	display         = "display"

	fontFamily = "font-family"
	fontSize   = "font-size"

	width    = "width"
	maxWidth = "max-width"

	marginRight = "margin-right"
	marginLeft  = "margin-left"
)

const (
	// CSS Units
	Px      = "px"
	Percent = "%"

	Block = "block"
	Auto  = "auto"
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

func MaxWidth(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(maxWidth, s, include...)
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

func Color(s string, include ...bool) KeyValuePair {
	return constructKeyValuePair(color, s, include...)
}

func RGBA(r, g, b, a float64) string {
	return fmt.Sprintf("rgba(%v,%v,%v,%v)", r, g, b, a)
}

func RGB(r, g, b float64) string {
	return fmt.Sprintf("rgb(%v,%v,%v)", r, g, b)
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
