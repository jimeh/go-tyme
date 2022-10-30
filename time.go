package tyme

import (
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Time is a wrapper around time.Time that implements JSON and YAML
// marshaler and unmarshaler interfaces.
//
// It marshals to a string in RFC 3339 format, with sub-second
// precision added if present.
//
// It unmarshals from a wide range of string date and time formats, by using the
// dateparse package.
type Time time.Time

// MarshalJSON implements the json.Marshaler interface, and formats the time as
// a JSON string in RFC 3339 format, with sub-second precision added if
// present.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(RFC3339NanoJSON)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface, and parses a wide
// range of string date and time formats, by using the dateparse package.
func (t *Time) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	nt, err := Parse(s)
	if err != nil {
		return err
	}

	*t = nt

	return err
}

// MarshalYAML implements the yaml.Marshaler interface, and formats the time as
// a YAML timestamp type in RFC 3339 string format, with sub-second precision
// added if present.
func (t Time) MarshalYAML() (interface{}, error) {
	return time.Time(t), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface, and parses a wide
// range of date and time formats, by using the dateparse package.
func (t *Time) UnmarshalYAML(node *yaml.Node) error {
	var nt Time
	var err error

	switch node.Tag {
	case "!!timestamp":
		var tt time.Time
		err = node.Decode(&tt)
		nt = Time(tt)
	case "!!str":
		nt, err = Parse(node.Value)
	default:
		return &yaml.TypeError{Errors: []string{"invalid time format"}}
	}
	if err != nil {
		return err
	}

	*t = nt

	return nil
}
