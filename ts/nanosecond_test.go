package ts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

var (
	minNano = unixNano(-9223372036854775808).UTC()
	maxNano = unixNano(9223372036854775807).UTC()
)

func nanosecondSkipTestCase(t *testing.T, ti time.Time) bool {
	t.Helper()

	return ti.Before(minNano) || ti.After(maxNano)
}

func TestNanosecond_MarshalUnmarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			v := Nanosecond(tt.t)

			b, err := json.Marshal(v)
			require.NoError(t, err)

			assert.Equal(t, tt.nanosecond, string(b))

			if nanosecondSkipTestCase(t, tt.t) {
				return
			}

			var got Nanosecond
			err = json.Unmarshal(b, &got)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Nanosecond)

			assert.Equal(t, want.UTC(), time.Time(got).UTC())
			assert.Equal(t, time.Local, time.Time(got).Location())
		})
	}
}

func TestNanosecond_MarshalUnmarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			v := Nanosecond(tt.t)

			b, err := yaml.Marshal(v)
			require.NoError(t, err)

			assert.Equal(t, tt.nanosecond+"\n", string(b))

			if nanosecondSkipTestCase(t, tt.t) {
				return
			}

			var got Nanosecond
			err = yaml.Unmarshal(b, &got)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Nanosecond)

			assert.Equal(t, want.UTC(), time.Time(got).UTC())
			assert.Equal(t, time.Local, time.Time(got).Location())
		})
	}
}

func TestNanosecond_MarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			ts := Nanosecond(tt.t)

			b, err := json.Marshal(ts)
			require.NoError(t, err)

			assert.Equal(t, tt.nanosecond, string(b))
		})
	}
}

func TestNanosecond_MarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			ts := Nanosecond(tt.t)

			b, err := yaml.Marshal(ts)
			require.NoError(t, err)

			assert.Equal(t, tt.nanosecond+"\n", string(b))
		})
	}
}

func TestNanosecond_UnmarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		if nanosecondSkipTestCase(t, tt.t) {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			var ts Nanosecond

			err := json.Unmarshal([]byte(tt.nanosecond), &ts)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Nanosecond)

			assert.Equal(t, want.UTC(), time.Time(ts).UTC())
		})
	}
	for _, tt := range unmarshalTestCases {
		if tt.nanosecond == "" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			var ts Nanosecond

			err := json.Unmarshal([]byte(tt.nanosecond), &ts)

			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				want := tt.t.Truncate(time.Nanosecond)
				assert.Equal(t, want.UTC(), time.Time(ts).UTC())
			}
		})
	}
}

func TestNanosecond_UnmarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		if nanosecondSkipTestCase(t, tt.t) {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			var ts Nanosecond

			err := yaml.Unmarshal([]byte(tt.nanosecond), &ts)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Nanosecond)

			assert.Equal(t, want.UTC(), time.Time(ts).UTC())
		})
	}
	for _, tt := range unmarshalTestCases {
		if tt.nanosecond == "" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			var ts Nanosecond

			err := yaml.Unmarshal([]byte(tt.nanosecond), &ts)

			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				want := tt.t.Truncate(time.Nanosecond)
				assert.Equal(t, want.UTC(), time.Time(ts).UTC())
			}
		})
	}
}
