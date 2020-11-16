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
	key      string
	value    string
        Include  bool
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

func constructKeyValuePair (key string , value interface{}, unit string, include ...bool) KeyValuePair {
    i := true
    if len(include) > 0 {
        i = include[0]
    }

    return KeyValuePair{key, fmt.Sprintf("%v%s", value, unit), i}
}

func Width(w int, unit string, include ...bool) KeyValuePair {
        return constructKeyValuePair(width, w, unit, include...)
}
