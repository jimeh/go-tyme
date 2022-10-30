package dur

import (
	"encoding/json"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Duration is a wrapper around time.Duration that implements JSON and YAML
// marshaler and unmarshaler interfaces.
//
// When unmarshaling, string values in JSON and YAML are parsed using
// time.ParseDuration. Numeric values are parsed as number of seconds.
//
// When marshaling, the duration is formatted as a string using time.Duration's
// String method.
type Duration time.Duration

// MarshalJSON implements the json.Marshaler interface, returning the duration
// as a string in the format "1h2m3s", same as time.Duration.String().
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

// UmarshalJSON implements the json.Unmarshaler interface. Supports string,
// numeric JSON types, converting them to a Duration using ParseDuration.
func (d *Duration) UnmarshalJSON(b []byte) error {
	var x interface{}
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	pd, err := Parse(x)
	if err != nil {
		return err
	}

	*d = pd

	return nil
}

// MarshalJSON implements the yaml.Marshaler interface, returning the duration
// as a string in the format 1h2m3s, same as time.Duration.String().
func (d Duration) MarshalYAML() (interface{}, error) {
	return time.Duration(d).String(), nil
}

// UmarshalYAML implements the yaml.Unmarshaler interface. Supports string, int
// and float YAML types, converting them to a Duration using ParseDuration.
func (d *Duration) UnmarshalYAML(node *yaml.Node) error {
	var x interface{}
	var err error

	switch node.Tag {
	case "!!str":
		x = node.Value
	case "!!int":
		x, err = strconv.Atoi(node.Value)
	case "!!float":
		x, err = strconv.ParseFloat(node.Value, 64)
	default:
		return &yaml.TypeError{Errors: []string{"invalid duration"}}
	}
	if err != nil {
		return err
	}

	pd, err := Parse(x)
	if err != nil {
		return err
	}

	*d = pd

	return nil
}
