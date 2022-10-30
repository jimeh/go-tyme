package tyme

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

var (
	loc  = time.FixedZone("UTC+8", 8*60*60)
	utc  = time.Date(2022, 10, 29, 14, 40, 34, 934349003, time.UTC)
	utc8 = utc.In(loc)
)

var timeMarshalUnmarshalTestCases = []struct {
	name string
	t    time.Time
}{
	{
		name: "UTC nanosecond precision",
		t:    utc.Round(time.Nanosecond),
	},
	{
		name: "UTC microsecond precision",
		t:    utc.Round(time.Microsecond),
	},
	{
		name: "UTC millisecond precision",
		t:    utc.Round(time.Millisecond),
	},
	{
		name: "UTC second precision",
		t:    utc.Round(time.Second),
	},
	{
		name: "UTC minute precision",
		t:    utc.Round(time.Second),
	},
	{
		name: "UTC+8 nanosecond precision",
		t:    utc8.Round(time.Nanosecond),
	},
	{
		name: "UTC+8 microsecond precision",
		t:    utc8.Round(time.Microsecond),
	},
	{
		name: "UTC+8 millisecond precision",
		t:    utc8.Round(time.Millisecond),
	},
	{
		name: "UTC+8 second precision",
		t:    utc8.Round(time.Second),
	},
	{
		name: "UTC+8 minute precision",
		t:    utc8.Round(time.Second),
	},
}

var timeMarshalTestCases = []struct {
	name string
	t    time.Time
	want string
}{
	{
		name: "UTC nanosecond precision",
		t:    utc.Round(time.Nanosecond),
		want: `2022-10-29T14:40:34.934349003Z`,
	},
	{
		name: "UTC microsecond precision",
		t:    utc.Round(time.Microsecond),
		want: `2022-10-29T14:40:34.934349Z`,
	},
	{
		name: "UTC millisecond precision",
		t:    utc.Round(time.Millisecond),
		want: `2022-10-29T14:40:34.934Z`,
	},
	{
		name: "UTC second precision",
		t:    utc.Round(time.Second),
		want: `2022-10-29T14:40:35Z`,
	},
	{
		name: "UTC minute precision",
		t:    utc.Round(time.Minute),
		want: `2022-10-29T14:41:00Z`,
	},
	{
		name: "UTC+8 nanosecond precision",
		t:    utc8.Round(time.Nanosecond),
		want: `2022-10-29T22:40:34.934349003+08:00`,
	},
	{
		name: "UTC+8 microsecond precision",
		t:    utc8.Round(time.Microsecond),
		want: `2022-10-29T22:40:34.934349+08:00`,
	},
	{
		name: "UTC+8 millisecond precision",
		t:    utc8.Round(time.Millisecond),
		want: `2022-10-29T22:40:34.934+08:00`,
	},
	{
		name: "UTC+8 second precision",
		t:    utc8.Round(time.Second),
		want: `2022-10-29T22:40:35+08:00`,
	},
	{
		name: "UTC+8 minute precision",
		t:    utc8.Round(time.Minute),
		want: `2022-10-29T22:41:00+08:00`,
	},
}

var timeUnmarshalTestCases = []struct {
	name string
	s    string
	want time.Time
}{
	{s: `1667054434934349003`, want: utc.Round(time.Nanosecond)},
	{s: `1667054434934349`, want: utc.Round(time.Microsecond)},
	{s: `1667054434934`, want: utc.Round(time.Millisecond)},
	{s: `1667054435`, want: utc.Round(time.Second)},
	{s: `20221029144035`, want: utc.Round(time.Second)},
	{s: `221029 14:40:35`, want: utc.Round(time.Second)},
	{s: `October 29th, 2022, 14:40:35`, want: utc.Round(time.Second)},
	{s: `2022-10-29 14:40:35`, want: utc.Round(time.Second)},
	{s: `2022-10-29T14:40:34.934349003Z`, want: utc.Round(time.Nanosecond)},
	{s: `2022-10-29T14:40:34.934349Z`, want: utc.Round(time.Microsecond)},
	{s: `2022-10-29T14:40:34.934Z`, want: utc.Round(time.Millisecond)},
	{s: `2022-10-29T14:40:35Z`, want: utc.Round(time.Second)},
	{s: `2022-10-29T14:41:00Z`, want: utc.Round(time.Minute)},
	{s: `2022-10-29T22:40:34.934349003+08:00`, want: utc8.Round(time.Nanosecond)},
	{s: `2022-10-29T22:40:34.934349+08:00`, want: utc8.Round(time.Microsecond)},
	{s: `2022-10-29T22:40:34.934+08:00`, want: utc8.Round(time.Millisecond)},
	{s: `2022-10-29T22:40:35+08:00`, want: utc8.Round(time.Second)},
	{s: `2022-10-29T22:41:00+08:00`, want: utc8.Round(time.Minute)},
}

func TestTime_MarshalUnmarshalJSON(t *testing.T) {
	for _, tt := range timeMarshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t1 := Time(tt.t)

			b, err := json.Marshal(t1)
			require.NoError(t, err)

			var t2 Time
			err = json.Unmarshal(b, &t2)
			require.NoError(t, err)

			assert.WithinDuration(t, tt.t, time.Time(t2), time.Nanosecond)
		})
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	for _, tt := range timeMarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t1 := Time(tt.t)

			b, err := json.Marshal(t1)
			require.NoError(t, err)

			assert.Equal(t, "\""+tt.want+"\"", string(b))
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	for _, tt := range timeUnmarshalTestCases {
		t.Run(tt.s, func(t *testing.T) {
			var got Time

			err := json.Unmarshal([]byte("\""+tt.s+"\""), &got)
			require.NoError(t, err)

			assert.WithinDuration(t, tt.want, time.Time(got), time.Nanosecond)
		})
	}
}

func TestTime_MarshalUnmarshalYAML(t *testing.T) {
	for _, tt := range timeMarshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t1 := Time(tt.t)

			b, err := yaml.Marshal(t1)
			require.NoError(t, err)

			var t2 Time
			err = yaml.Unmarshal(b, &t2)
			require.NoError(t, err)

			assert.WithinDuration(t, tt.t, time.Time(t2), time.Nanosecond)
		})
	}
}

func TestTime_MarshalYAML(t *testing.T) {
	for _, tt := range timeMarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t1 := Time(tt.t)

			b, err := yaml.Marshal(t1)
			require.NoError(t, err)

			assert.Equal(t, tt.want+"\n", string(b))
		})
	}
}

func TestTime_UnmarshalYAML(t *testing.T) {
	for _, tt := range timeUnmarshalTestCases {
		t.Run(tt.s, func(t *testing.T) {
			var got Time

			err := yaml.Unmarshal([]byte("\""+tt.s+"\""), &got)
			require.NoError(t, err)

			assert.WithinDuration(t, tt.want, time.Time(got), time.Nanosecond)
		})
	}
}
