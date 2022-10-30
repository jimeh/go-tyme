package tyme

import (
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// TimeRFC3339 is a wrapper around time.Time that implements JSON/YAML
// marshaling/unmarshaling interfaces. As opposed to Time, only RFC 3339
// formatted input is accepted when unmarshaling.
//
// It marshals to a string in RFC 3339 format, with sub-second precision added
// if present.
//
// It will only unmarshal from a string in RFC 3339 format. Any other format
// will cause an unmarshaling error.
type TimeRFC3339 time.Time

// RFC3339NanoJSON is the time.RFC3339Nano format with double quotes around it.
const RFC3339NanoJSON = `"` + time.RFC3339Nano + `"`

// MarshalJSON implements the json.Marshaler interface, and formats the time as
// a JSON string in RFC 3339 format, with sub-second precision added if
// present.
func (t TimeRFC3339) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(RFC3339NanoJSON)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface, and parses a
// JSON string in RFC 3339 format, with sub-second precision added if
// present.
func (t *TimeRFC3339) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	nt, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return err
	}

	*t = TimeRFC3339(nt)

	return err
}

// MarshalYAML implements the yaml.Marshaler interface, and formats the time as
// a YAML timestamp type in RFC 3339 string format, with sub-second precision
// added if present.
func (t TimeRFC3339) MarshalYAML() (interface{}, error) {
	return time.Time(t), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface, and parses a YAML
// timestamp or string formatted according to RFC 3339, with sub-second
// precision added if present.
func (t *TimeRFC3339) UnmarshalYAML(node *yaml.Node) error {
	var nt time.Time
	var err error

	switch node.Tag {
	case "!!timestamp":
		err = node.Decode(&nt)
	case "!!str":
		nt, err = time.Parse(time.RFC3339Nano, node.Value)
	default:
		return &yaml.TypeError{Errors: []string{"invalid time format"}}
	}
	if err != nil {
		return err
	}

	*t = TimeRFC3339(nt)

	return nil
}
