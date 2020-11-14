package css

import (
	"fmt"
	"strings"
)

const (
	width = "width"
)

const (
	// CSS Units
	Px = "px"
)

type KeyValuePair struct {
	key   string
	value string
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

func Width(w int, unit string) KeyValuePair {
	return KeyValuePair{width, fmt.Sprintf("%v%s", w, unit)}
}
