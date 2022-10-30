package tyme

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

var timeRFC3339UnmarshalTestCases = []struct {
	name string
	s    string
	want time.Time
}{
	{
		name: "UTC nanosecond precision",
		s:    `2022-10-29T14:40:34.934349003Z`,
		want: utc.Round(time.Nanosecond),
	},
	{
		name: "UTC microsecond precision",
		s:    `2022-10-29T14:40:34.934349Z`,
		want: utc.Round(time.Microsecond),
	},
	{
		name: "UTC millisecond precision",
		s:    `2022-10-29T14:40:34.934Z`,
		want: utc.Round(time.Millisecond),
	},
	{
		name: "UTC second precision",
		s:    `2022-10-29T14:40:35Z`,
		want: utc.Round(time.Second),
	},
	{
		name: "UTC minute precision",
		s:    `2022-10-29T14:41:00Z`,
		want: utc.Round(time.Minute),
	},
	{
		name: "UTC+8 nanosecond precision",
		s:    `2022-10-29T22:40:34.934349003+08:00`,
		want: utc8.Round(time.Nanosecond),
	},
	{
		name: "UTC+8 microsecond precision",
		s:    `2022-10-29T22:40:34.934349+08:00`,
		want: utc8.Round(time.Microsecond),
	},
	{
		name: "UTC+8 millisecond precision",
		s:    `2022-10-29T22:40:34.934+08:00`,
		want: utc8.Round(time.Millisecond),
	},
	{
		name: "UTC+8 second precision",
		s:    `2022-10-29T22:40:35+08:00`,
		want: utc8.Round(time.Second),
	},
	{
		name: "UTC+8 minute precision",
		s:    `2022-10-29T22:41:00+08:00`,
		want: utc8.Round(time.Minute),
	},
}

func TestTimeRFC3339_MarshalUnmarshalJSON(t *testing.T) {
	for _, tt := range timeMarshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t1 := TimeRFC3339(tt.t)

			b, err := json.Marshal(t1)
			require.NoError(t, err)

			var t2 TimeRFC3339
			err = json.Unmarshal(b, &t2)
			require.NoError(t, err)

			assert.WithinDuration(t, tt.t, time.Time(t2), time.Nanosecond)
		})
	}
}

func TestTimeRFC3339_MarshalJSON(t *testing.T) {
	for _, tt := range timeMarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t1 := TimeRFC3339(tt.t)

			b, err := json.Marshal(t1)
			require.NoError(t, err)

			assert.Equal(t, "\""+tt.want+"\"", string(b))
		})
	}
}

func TestTimeRFC3339_UnmarshalJSON(t *testing.T) {
	for _, tt := range timeRFC3339UnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var got TimeRFC3339

			err := json.Unmarshal([]byte("\""+tt.s+"\""), &got)
			require.NoError(t, err)

			assert.WithinDuration(t, tt.want, time.Time(got), time.Nanosecond)
		})
	}
}

func TestTimeRFC3339_MarshalUnmarshalYAML(t *testing.T) {
	for _, tt := range timeMarshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t1 := TimeRFC3339(tt.t)

			b, err := yaml.Marshal(t1)
			require.NoError(t, err)

			var t2 TimeRFC3339
			err = yaml.Unmarshal(b, &t2)
			require.NoError(t, err)

			assert.WithinDuration(t, tt.t, time.Time(t2), time.Nanosecond)
		})
	}
}

func TestTimeRFC3339_MarshalYAML(t *testing.T) {
	for _, tt := range timeMarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t1 := TimeRFC3339(tt.t)

			b, err := yaml.Marshal(t1)
			require.NoError(t, err)

			assert.Equal(t, tt.want+"\n", string(b))
		})
	}
}

func TestTimeRFC3339_UnmarshalYAML(t *testing.T) {
	for _, tt := range timeRFC3339UnmarshalTestCases {
		t.Run(tt.name+" (YAML string type)", func(t *testing.T) {
			var got TimeRFC3339

			err := yaml.Unmarshal([]byte("\""+tt.s+"\""), &got)
			require.NoError(t, err)

			assert.WithinDuration(t, tt.want, time.Time(got), time.Nanosecond)
		})

		t.Run(tt.name+" (YAML timestamp type)", func(t *testing.T) {
			var got TimeRFC3339

			err := yaml.Unmarshal([]byte(tt.s), &got)
			require.NoError(t, err)

			assert.WithinDuration(t, tt.want, time.Time(got), time.Nanosecond)
		})
	}
}
