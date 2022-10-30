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
	minMicro = time.UnixMicro(-9223372036854775808).UTC()
	maxMicro = time.UnixMicro(9223372036854775807).UTC()
)

func microsecondSkipTestCase(t *testing.T, ti time.Time) bool {
	t.Helper()

	return ti.Before(minMicro) || ti.After(maxMicro)
}

func TestMicrosecond_MarshalUnmarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			v := Microsecond(tt.t)

			b, err := json.Marshal(v)
			require.NoError(t, err)

			assert.Equal(t, tt.microsecond, string(b))

			if microsecondSkipTestCase(t, tt.t) {
				return
			}

			var got Microsecond
			err = json.Unmarshal(b, &got)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Microsecond)

			assert.Equal(t, want.UTC(), time.Time(got).UTC())
			assert.Equal(t, time.Local, time.Time(got).Location())
		})
	}
}

func TestMicrosecond_MarshalUnmarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			v := Microsecond(tt.t)

			b, err := yaml.Marshal(v)
			require.NoError(t, err)

			assert.Equal(t, tt.microsecond+"\n", string(b))

			if microsecondSkipTestCase(t, tt.t) {
				return
			}

			var got Microsecond
			err = yaml.Unmarshal(b, &got)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Microsecond)

			assert.Equal(t, want.UTC(), time.Time(got).UTC())
			assert.Equal(t, time.Local, time.Time(got).Location())
		})
	}
}

func TestMicrosecond_MarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			ts := Microsecond(tt.t)

			b, err := json.Marshal(ts)
			require.NoError(t, err)

			assert.Equal(t, tt.microsecond, string(b))
		})
	}
}

func TestMicrosecond_MarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			ts := Microsecond(tt.t)

			b, err := yaml.Marshal(ts)
			require.NoError(t, err)

			assert.Equal(t, tt.microsecond+"\n", string(b))
		})
	}
}

func TestMicrosecond_UnmarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		if microsecondSkipTestCase(t, tt.t) {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			var ts Microsecond

			err := json.Unmarshal([]byte(tt.microsecond), &ts)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Microsecond)

			assert.Equal(t, want.UTC(), time.Time(ts).UTC())
		})
	}
	for _, tt := range unmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var ts Microsecond

			err := json.Unmarshal([]byte(tt.microsecond), &ts)

			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				want := tt.t.Truncate(time.Microsecond)
				assert.Equal(t, want.UTC(), time.Time(ts).UTC())
			}
		})
	}
}

func TestMicrosecond_UnmarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		if microsecondSkipTestCase(t, tt.t) {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			var ts Microsecond

			err := yaml.Unmarshal([]byte(tt.microsecond), &ts)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Microsecond)

			assert.Equal(t, want.UTC(), time.Time(ts).UTC())
		})
	}
	for _, tt := range unmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var ts Microsecond

			err := yaml.Unmarshal([]byte(tt.microsecond), &ts)

			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				want := tt.t.Truncate(time.Microsecond)
				assert.Equal(t, want.UTC(), time.Time(ts).UTC())
			}
		})
	}
}
