package ts

import (
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Millisecond is a wrapper around time.Time for marshaling to/from JSON/YAML as
// millisecond-based numeric Unix timestamps.
//
// It marshals to a JSON/YAML number representing the number of milliseconds
// since the Unix time epoch.
//
// It unmarshals from a JSON/YAML number representing the number of milliseconds
// since the Unix time epoch.
type Millisecond time.Time

// Time returns the time.Time corresponding to the millisecond instant s.
func (ms Millisecond) Time() time.Time {
	return time.Time(ms)
}

// Local returns the local time corresponding to the millisecond instant s.
func (ms Millisecond) Local() Millisecond {
	return Millisecond(time.Time(ms).Local())
}

// GoString implements the fmt.GoStringer interface.
func (ms Millisecond) GoString() string {
	return time.Time(ms).GoString()
}

// IsDST reports whether the millisecond instant s occurs within Daylight Saving
// Time.
func (ms Millisecond) IsDST() bool {
	return time.Time(ms).IsDST()
}

// IsZero returns true if the Millisecond is the zero value.
func (ms Millisecond) IsZero() bool {
	return time.Time(ms).IsZero()
}

// String calls time.Time.String.
func (ms Millisecond) String() string {
	return time.Time(ms).String()
}

// UTC returns a copy of the Millisecond with the location set to UTC.
func (ms Millisecond) UTC() Millisecond {
	return Millisecond(time.Time(ms).UTC())
}

// MarshalJSON implements the json.Marshaler interface.
func (ms Millisecond) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(ms).UnixMilli(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ms *Millisecond) UnmarshalJSON(data []byte) error {
	i, err := unmarshalBytes(data)
	if err != nil {
		return err
	}

	*ms = UnixMilli(i)

	return nil
}

// MarshalJSON implements the yaml.Marshaler interface.
func (ms Millisecond) MarshalYAML() (interface{}, error) {
	return time.Time(ms).UnixMilli(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (ms *Millisecond) UnmarshalYAML(node *yaml.Node) error {
	i, err := unmarshalYAMLNode(node)
	if err != nil {
		return err
	}

	*ms = UnixMilli(i)

	return nil
}
