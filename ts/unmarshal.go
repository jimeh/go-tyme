package ts

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

func unmarshalBytes(data []byte) (int64, error) {
	s, err := strconv.Unquote(string(data))
	if err == nil {
		data = []byte(s)
	}

	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		var f float64
		f, err = strconv.ParseFloat(string(data), 64)
		i = int64(f)
	}

	if err != nil {
		return 0, fmt.Errorf("invalid numeric timestamp: %s", string(data))
	}

	return i, nil
}

func unmarshalYAMLNode(node *yaml.Node) (int64, error) {
	var i int64
	var err error
	var invalid bool

	switch node.Tag {
	case "!!int", "!!str":
		i, err = strconv.ParseInt(node.Value, 10, 64)
	case "!!float":
		var f float64
		f, err = strconv.ParseFloat(node.Value, 64)
		i = int64(f)
	default:
		invalid = true
	}

	if err != nil || invalid {
		return 0, &yaml.TypeError{Errors: []string{"invalid numeric timestamp"}}
	}

	return i, nil
}
