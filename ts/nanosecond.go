package ts

import (
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Nanosecond is a wrapper around time.Time for marshaling to/from JSON/YAML as
// nanosecond-based numeric Unix timestamps.
//
// It marshals to a JSON/YAML number representing the number of nanoseconds
// since the Unix time epoch.
//
// It unmarshals from a JSON/YAML number representing the number of nanoseconds
// since the Unix time epoch.
type Nanosecond time.Time

// Time returns the time.Time corresponding to the nanosecond instant s.
func (ns Nanosecond) Time() time.Time {
	return time.Time(ns)
}

// Local returns the local time corresponding to the nanosecond instant s.
func (ns Nanosecond) Local() Nanosecond {
	return Nanosecond(time.Time(ns).Local())
}

// GoString implements the fmt.GoStringer interface.
func (ns Nanosecond) GoString() string {
	return time.Time(ns).GoString()
}

// IsDST reports whether the nanosecond instant s occurs within Daylight Saving
// Time.
func (ns Nanosecond) IsDST() bool {
	return time.Time(ns).IsDST()
}

// IsZero returns true if the Nanosecond is the zero value.
func (ns Nanosecond) IsZero() bool {
	return time.Time(ns).IsZero()
}

// String calls time.Time.String.
func (ns Nanosecond) String() string {
	return time.Time(ns).String()
}

// UTC returns a copy of the Nanosecond with the location set to UTC.
func (ns Nanosecond) UTC() Nanosecond {
	return Nanosecond(time.Time(ns).UTC())
}

// MarshalJSON implements the json.Marshaler interface.
func (ns Nanosecond) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(ns).UnixNano(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ns *Nanosecond) UnmarshalJSON(data []byte) error {
	i, err := unmarshalBytes(data)
	if err != nil {
		return err
	}

	*ns = UnixNano(i)

	return nil
}

// MarshalJSON implements the yaml.Marshaler interface.
func (ns Nanosecond) MarshalYAML() (interface{}, error) {
	return time.Time(ns).UnixNano(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (ns *Nanosecond) UnmarshalYAML(node *yaml.Node) error {
	i, err := unmarshalYAMLNode(node)
	if err != nil {
		return err
	}

	*ns = UnixNano(i)

	return nil
}

func unixNano(ts int64) time.Time {
	return time.Unix(ts/1e9, ts%1e9)
}
