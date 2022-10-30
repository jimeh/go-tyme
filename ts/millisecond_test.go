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
	minMilli = time.UnixMilli(-9223372036854775808).UTC()
	maxMilli = time.UnixMilli(9223372036854775807).UTC()
)

func millisecondSkipTestCase(t *testing.T, ti time.Time) bool {
	t.Helper()

	return ti.Before(minMilli) || ti.After(maxMilli)
}

func TestMillisecond_MarshalUnmarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			v := Millisecond(tt.t)

			b, err := json.Marshal(v)
			require.NoError(t, err)

			assert.Equal(t, tt.millisecond, string(b))

			if millisecondSkipTestCase(t, tt.t) {
				return
			}

			var got Millisecond
			err = json.Unmarshal(b, &got)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Millisecond)

			assert.Equal(t, want.UTC(), time.Time(got).UTC())
			assert.Equal(t, time.Local, time.Time(got).Location())
		})
	}
}

func TestMillisecond_MarshalUnmarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			v := Millisecond(tt.t)

			b, err := yaml.Marshal(v)
			require.NoError(t, err)

			assert.Equal(t, tt.millisecond+"\n", string(b))

			if millisecondSkipTestCase(t, tt.t) {
				return
			}

			var got Millisecond
			err = yaml.Unmarshal(b, &got)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Millisecond)

			assert.Equal(t, want.UTC(), time.Time(got).UTC())
			assert.Equal(t, time.Local, time.Time(got).Location())
		})
	}
}

func TestMillisecond_MarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			ts := Millisecond(tt.t)

			b, err := json.Marshal(ts)
			require.NoError(t, err)

			assert.Equal(t, tt.millisecond, string(b))
		})
	}
}

func TestMillisecond_MarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			ts := Millisecond(tt.t)

			b, err := yaml.Marshal(ts)
			require.NoError(t, err)

			assert.Equal(t, tt.millisecond+"\n", string(b))
		})
	}
}

func TestMillisecond_UnmarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		if millisecondSkipTestCase(t, tt.t) {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			var ts Millisecond

			err := json.Unmarshal([]byte(tt.millisecond), &ts)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Millisecond)

			assert.Equal(t, want.UTC(), time.Time(ts).UTC())
		})
	}
	for _, tt := range unmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var ts Millisecond

			err := json.Unmarshal([]byte(tt.millisecond), &ts)

			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				want := tt.t.Truncate(time.Millisecond)
				assert.Equal(t, want.UTC(), time.Time(ts).UTC())
			}
		})
	}
}

func TestMillisecond_UnmarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		if millisecondSkipTestCase(t, tt.t) {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			var ts Millisecond

			err := yaml.Unmarshal([]byte(tt.millisecond), &ts)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Millisecond)

			assert.Equal(t, want.UTC(), time.Time(ts).UTC())
		})
	}
	for _, tt := range unmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var ts Millisecond

			err := yaml.Unmarshal([]byte(tt.millisecond), &ts)

			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				want := tt.t.Truncate(time.Millisecond)
				assert.Equal(t, want.UTC(), time.Time(ts).UTC())
			}
		})
	}
}
