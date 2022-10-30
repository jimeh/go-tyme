package dur

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestDuration_MarshalUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
	}{
		{name: "zero", d: 0},
		{name: "1ns", d: 1 * time.Nanosecond},
		{name: "2ns", d: 2 * time.Nanosecond},
		{name: "1µs", d: 1 * time.Microsecond},
		{name: "2µs", d: 2 * time.Microsecond},
		{name: "1ms", d: 1 * time.Millisecond},
		{name: "2ms", d: 2 * time.Millisecond},
		{name: "1s", d: 1 * time.Second},
		{name: "2s", d: 2 * time.Second},
		{name: "90s", d: 90 * time.Second},
		{name: "1m", d: 1 * time.Minute},
		{name: "2m", d: 2 * time.Minute},
		{name: "90m", d: 90 * time.Minute},
		{name: "1h", d: 1 * time.Hour},
		{name: "2h", d: 2 * time.Hour},
		{name: "36h", d: 36 * time.Hour},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Duration(tt.d)

			b, err := json.Marshal(d)
			require.NoError(t, err)

			var d2 Duration
			err = json.Unmarshal(b, &d2)
			require.NoError(t, err)

			assert.Equal(t, tt.d, time.Duration(d2))
		})
	}
}

func TestDuration_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{name: "zero", d: 0, want: `"0s"`},
		{name: "1ns", d: 1 * time.Nanosecond, want: `"1ns"`},
		{name: "2ns", d: 2 * time.Nanosecond, want: `"2ns"`},
		{name: "1µs", d: 1 * time.Microsecond, want: `"1µs"`},
		{name: "2µs", d: 2 * time.Microsecond, want: `"2µs"`},
		{name: "1ms", d: 1 * time.Millisecond, want: `"1ms"`},
		{name: "2ms", d: 2 * time.Millisecond, want: `"2ms"`},
		{name: "1s", d: 1 * time.Second, want: `"1s"`},
		{name: "2s", d: 2 * time.Second, want: `"2s"`},
		{name: "90s", d: 90 * time.Second, want: `"1m30s"`},
		{name: "1m", d: 1 * time.Minute, want: `"1m0s"`},
		{name: "2m", d: 2 * time.Minute, want: `"2m0s"`},
		{name: "90m", d: 90 * time.Minute, want: `"1h30m0s"`},
		{name: "1h", d: 1 * time.Hour, want: `"1h0m0s"`},
		{name: "2h", d: 2 * time.Hour, want: `"2h0m0s"`},
		{name: "36h", d: 36 * time.Hour, want: `"36h0m0s"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Duration(tt.d)

			b, err := json.Marshal(d)
			require.NoError(t, err)

			assert.Equal(t, tt.want, string(b))
		})
	}
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    time.Duration
		wantErr string
	}{
		{s: `"1ns"`, want: 1 * time.Nanosecond},
		{s: `"2ns"`, want: 2 * time.Nanosecond},
		{s: `"1µs"`, want: 1 * time.Microsecond},
		{s: `"2µs"`, want: 2 * time.Microsecond},
		{s: `"1ms"`, want: 1 * time.Millisecond},
		{s: `"2ms"`, want: 2 * time.Millisecond},
		{s: `"1s"`, want: 1 * time.Second},
		{s: `"2s"`, want: 2 * time.Second},
		{s: `"90s"`, want: 90 * time.Second},
		{s: `"1m30s"`, want: 90 * time.Second},
		{s: `"1m0s"`, want: 1 * time.Minute},
		{s: `"2m0s"`, want: 2 * time.Minute},
		{s: `"90m"`, want: 90 * time.Minute},
		{s: `"1h30m"`, want: 90 * time.Minute},
		{s: `"1h30m0s"`, want: 90 * time.Minute},
		{s: `"1h0m0s"`, want: 1 * time.Hour},
		{s: `"2h0m0s"`, want: 2 * time.Hour},
		{s: `"36h0m0s"`, want: 36 * time.Hour},
		{s: "0.000000001", want: 1 * time.Nanosecond},
		{s: "0.000001", want: 1 * time.Microsecond},
		{s: "0.001", want: 1 * time.Millisecond},
		{s: "0.1", want: 100 * time.Millisecond},
		{s: "1", want: 1 * time.Second},
		{s: "1.0", want: 1 * time.Second},
		{s: "2.0", want: 2 * time.Second},
		{s: "90", want: 90 * time.Second},
		{s: "90.001", want: (90 * time.Second) + (1 * time.Millisecond)},
		{s: "90.999", want: (90 * time.Second) + (999 * time.Millisecond)},
		{name: "empty", s: "", wantErr: "unexpected end of JSON input"},
		{
			s:       "'2ms'",
			wantErr: "invalid character '\\'' looking for beginning of value",
		},
		{
			s:       "nil",
			wantErr: "invalid character 'i' in literal null (expecting 'u')",
		},
		{s: `"nil"`, wantErr: "time: invalid duration \"nil\""},
		{
			s:       "foo",
			wantErr: "invalid character 'o' in literal false (expecting 'a')",
		},
		{s: `"foo"`, wantErr: "time: invalid duration \"foo\""},
		{s: "null", wantErr: "time: invalid duration <nil>"},
		{s: `"null"`, wantErr: "time: invalid duration \"null\""},
	}
	for _, tt := range tests {
		name := tt.name
		if name == "" {
			name = tt.s
		}

		t.Run(name, func(t *testing.T) {
			var d Duration

			err := json.Unmarshal([]byte(tt.s), &d)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.want, time.Duration(d))
		})
	}
}

func TestDuration_MarshalUnmarshalYAML(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
	}{
		{name: "zero", d: 0},
		{name: "1ns", d: 1 * time.Nanosecond},
		{name: "2ns", d: 2 * time.Nanosecond},
		{name: "1µs", d: 1 * time.Microsecond},
		{name: "2µs", d: 2 * time.Microsecond},
		{name: "1ms", d: 1 * time.Millisecond},
		{name: "2ms", d: 2 * time.Millisecond},
		{name: "1s", d: 1 * time.Second},
		{name: "2s", d: 2 * time.Second},
		{name: "90s", d: 90 * time.Second},
		{name: "1m", d: 1 * time.Minute},
		{name: "2m", d: 2 * time.Minute},
		{name: "90m", d: 90 * time.Minute},
		{name: "1h", d: 1 * time.Hour},
		{name: "2h", d: 2 * time.Hour},
		{name: "36h", d: 36 * time.Hour},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Duration(tt.d)

			b, err := yaml.Marshal(d)
			require.NoError(t, err)

			var d2 Duration
			err = yaml.Unmarshal(b, &d2)
			require.NoError(t, err)

			assert.Equal(t, tt.d, time.Duration(d2))
		})
	}
}

func TestDuration_MarshalYAML(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{name: "zero", d: 0, want: "0s\n"},
		{name: "1ns", d: 1 * time.Nanosecond, want: "1ns\n"},
		{name: "2ns", d: 2 * time.Nanosecond, want: "2ns\n"},
		{name: "1µs", d: 1 * time.Microsecond, want: "1µs\n"},
		{name: "2µs", d: 2 * time.Microsecond, want: "2µs\n"},
		{name: "1ms", d: 1 * time.Millisecond, want: "1ms\n"},
		{name: "2ms", d: 2 * time.Millisecond, want: "2ms\n"},
		{name: "1s", d: 1 * time.Second, want: "1s\n"},
		{name: "2s", d: 2 * time.Second, want: "2s\n"},
		{name: "90s", d: 90 * time.Second, want: "1m30s\n"},
		{name: "1m", d: 1 * time.Minute, want: "1m0s\n"},
		{name: "2m", d: 2 * time.Minute, want: "2m0s\n"},
		{name: "90m", d: 90 * time.Minute, want: "1h30m0s\n"},
		{name: "1h", d: 1 * time.Hour, want: "1h0m0s\n"},
		{name: "2h", d: 2 * time.Hour, want: "2h0m0s\n"},
		{name: "36h", d: 36 * time.Hour, want: "36h0m0s\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Duration(tt.d)

			b, err := yaml.Marshal(d)
			require.NoError(t, err)

			assert.Equal(t, tt.want, string(b))
		})
	}
}

func TestDuration_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    time.Duration
		wantErr string
	}{
		{s: `"1ns"`, want: 1 * time.Nanosecond},
		{s: `"2ns"`, want: 2 * time.Nanosecond},
		{s: `"1µs"`, want: 1 * time.Microsecond},
		{s: `"2µs"`, want: 2 * time.Microsecond},
		{s: `"1ms"`, want: 1 * time.Millisecond},
		{s: `"2ms"`, want: 2 * time.Millisecond},
		{s: `"1s"`, want: 1 * time.Second},
		{s: `"2s"`, want: 2 * time.Second},
		{s: `"90s"`, want: 90 * time.Second},
		{s: `"1m30s"`, want: 90 * time.Second},
		{s: `"1m0s"`, want: 1 * time.Minute},
		{s: `"2m0s"`, want: 2 * time.Minute},
		{s: `"90m"`, want: 90 * time.Minute},
		{s: `"1h30m"`, want: 90 * time.Minute},
		{s: `"1h30m0s"`, want: 90 * time.Minute},
		{s: `"1h0m0s"`, want: 1 * time.Hour},
		{s: `"2h0m0s"`, want: 2 * time.Hour},
		{s: `"36h0m0s"`, want: 36 * time.Hour},
		{s: "'1ns'", want: 1 * time.Nanosecond},
		{s: "'2ns'", want: 2 * time.Nanosecond},
		{s: "'1µs'", want: 1 * time.Microsecond},
		{s: "'2µs'", want: 2 * time.Microsecond},
		{s: "'1ms'", want: 1 * time.Millisecond},
		{s: "'2ms'", want: 2 * time.Millisecond},
		{s: "'1s'", want: 1 * time.Second},
		{s: "'2s'", want: 2 * time.Second},
		{s: "'90s'", want: 90 * time.Second},
		{s: "'1m30s'", want: 90 * time.Second},
		{s: "'1m0s'", want: 1 * time.Minute},
		{s: "'2m0s'", want: 2 * time.Minute},
		{s: "'90m'", want: 90 * time.Minute},
		{s: "'1h30m'", want: 90 * time.Minute},
		{s: "'1h30m0s'", want: 90 * time.Minute},
		{s: "'1h0m0s'", want: 1 * time.Hour},
		{s: "'2h0m0s'", want: 2 * time.Hour},
		{s: "'36h0m0s'", want: 36 * time.Hour},
		{s: "1ns", want: 1 * time.Nanosecond},
		{s: "2ns", want: 2 * time.Nanosecond},
		{s: "1µs", want: 1 * time.Microsecond},
		{s: "2µs", want: 2 * time.Microsecond},
		{s: "1ms", want: 1 * time.Millisecond},
		{s: "2ms", want: 2 * time.Millisecond},
		{s: "1s", want: 1 * time.Second},
		{s: "2s", want: 2 * time.Second},
		{s: "90s", want: 90 * time.Second},
		{s: "1m30s", want: 90 * time.Second},
		{s: "1m0s", want: 1 * time.Minute},
		{s: "2m0s", want: 2 * time.Minute},
		{s: "90m", want: 90 * time.Minute},
		{s: "1h30m", want: 90 * time.Minute},
		{s: "1h30m0s", want: 90 * time.Minute},
		{s: "1h0m0s", want: 1 * time.Hour},
		{s: "2h0m0s", want: 2 * time.Hour},
		{s: "36h0m0s", want: 36 * time.Hour},
		{name: "empty", s: "", want: 0},
		{s: "nil", wantErr: "time: invalid duration \"nil\""},
		{s: "foo", wantErr: "time: invalid duration \"foo\""},
		{s: "null", want: 0},
		{s: "0.000000001", want: 1 * time.Nanosecond},
		{s: "0.000001", want: 1 * time.Microsecond},
		{s: "0.001", want: 1 * time.Millisecond},
		{s: "0.1", want: 100 * time.Millisecond},
		{s: "1", want: 1 * time.Second},
		{s: "1.0", want: 1 * time.Second},
		{s: "2.0", want: 2 * time.Second},
		{s: "90", want: 90 * time.Second},
		{s: "90.001", want: (90 * time.Second) + (1 * time.Millisecond)},
		{s: "90.999", want: (90 * time.Second) + (999 * time.Millisecond)},
		{s: "yes", wantErr: "time: invalid duration \"yes\""},
		{s: "no", wantErr: "time: invalid duration \"no\""},
		{s: "true", wantErr: "yaml: unmarshal errors:\n  invalid duration"},
		{s: "false", wantErr: "yaml: unmarshal errors:\n  invalid duration"},
		{
			s:       "[foo, bar]",
			wantErr: "yaml: unmarshal errors:\n  invalid duration",
		},
		{
			s:       "{foo: bar}",
			wantErr: "yaml: unmarshal errors:\n  invalid duration",
		},
		{
			s:       "2001-12-15T02:59:43.1Z",
			wantErr: "yaml: unmarshal errors:\n  invalid duration",
		},
	}
	for _, tt := range tests {
		name := tt.name
		if name == "" {
			name = tt.s
		}

		t.Run(name, func(t *testing.T) {
			var d Duration

			err := yaml.Unmarshal([]byte(tt.s), &d)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.want, time.Duration(d))
		})
	}
}
