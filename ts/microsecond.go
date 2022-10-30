package ts

import (
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Microsecond is a wrapper around time.Time for marshaling to/from JSON/YAML as
// microsecond-based numeric Unix timestamps.
//
// It marshals to a JSON/YAML number representing the number of microseconds
// since the Unix time epoch.
//
// It unmarshals from a JSON/YAML number representing the number of microseconds
// since the Unix time epoch.
type Microsecond time.Time

// Time returns the time.Time corresponding to the microsecond instant s.
func (ms Microsecond) Time() time.Time {
	return time.Time(ms)
}

// Local returns the local time corresponding to the microsecond instant s.
func (ms Microsecond) Local() Microsecond {
	return Microsecond(time.Time(ms).Local())
}

// GoString implements the fmt.GoStringer interface.
func (ms Microsecond) GoString() string {
	return time.Time(ms).GoString()
}

// IsDST reports whether the microsecond instant s occurs within Daylight Saving
// Time.
func (ms Microsecond) IsDST() bool {
	return time.Time(ms).IsDST()
}

// IsZero returns true if the Microsecond is the zero value.
func (ms Microsecond) IsZero() bool {
	return time.Time(ms).IsZero()
}

// String calls time.Time.String.
func (ms Microsecond) String() string {
	return time.Time(ms).String()
}

// UTC returns a copy of the Microsecond with the location set to UTC.
func (ms Microsecond) UTC() Microsecond {
	return Microsecond(time.Time(ms).UTC())
}

// MarshalJSON implements the json.Marshaler interface.
func (ms Microsecond) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(ms).UnixMicro(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ms *Microsecond) UnmarshalJSON(data []byte) error {
	i, err := unmarshalBytes(data)
	if err != nil {
		return err
	}

	*ms = UnixMicro(i)

	return nil
}

// MarshalJSON implements the yaml.Marshaler interface.
func (ms Microsecond) MarshalYAML() (interface{}, error) {
	return time.Time(ms).UnixMicro(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (ms *Microsecond) UnmarshalYAML(node *yaml.Node) error {
	i, err := unmarshalYAMLNode(node)
	if err != nil {
		return err
	}

	*ms = UnixMicro(i)

	return nil
}
