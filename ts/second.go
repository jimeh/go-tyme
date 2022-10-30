package ts

import (
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Second is a wrapper around time.Time for marshaling to/from JSON/YAML as
// second-based numeric Unix timestamps.
//
// It marshals to a JSON/YAML number representing the number of seconds since
// the Unix time epoch.
//
// It unmarshals from a JSON/YAML number representing the number of seconds
// since the Unix time epoch.
type Second time.Time

// Time returns the time.Time corresponding to the second instant s.
func (s Second) Time() time.Time {
	return time.Time(s)
}

// Local returns the local time corresponding to the second instant s.
func (s Second) Local() Second {
	return Second(time.Time(s).Local())
}

// GoString implements the fmt.GoStringer interface.
func (s Second) GoString() string {
	return time.Time(s).GoString()
}

// IsDST reports whether the second instant s occurs within Daylight Saving
// Time.
func (s Second) IsDST() bool {
	return time.Time(s).IsDST()
}

// IsZero returns true if the Second is the zero value.
func (s Second) IsZero() bool {
	return time.Time(s).IsZero()
}

// String calls time.Time.String.
func (s Second) String() string {
	return time.Time(s).String()
}

// UTC returns a copy of the Second with the location set to UTC.
func (s Second) UTC() Second {
	return Second(time.Time(s).UTC())
}

// MarshalJSON implements the json.Marshaler interface.
func (s Second) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(s).Unix(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Second) UnmarshalJSON(data []byte) error {
	i, err := unmarshalBytes(data)
	if err != nil {
		return err
	}

	*s = UnixSecond(i)

	return nil
}

// MarshalJSON implements the yaml.Marshaler interface.
func (s Second) MarshalYAML() (interface{}, error) {
	return time.Time(s).Unix(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (s *Second) UnmarshalYAML(node *yaml.Node) error {
	i, err := unmarshalYAMLNode(node)
	if err != nil {
		return err
	}

	*s = UnixSecond(i)

	return nil
}
