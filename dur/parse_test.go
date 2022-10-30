package dur

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		x       interface{}
		want    time.Duration
		wantErr string
	}{
		{x: "1ns", want: 1 * time.Nanosecond},
		{x: "2ns", want: 2 * time.Nanosecond},
		{x: "1µs", want: 1 * time.Microsecond},
		{x: "2µs", want: 2 * time.Microsecond},
		{x: "1ms", want: 1 * time.Millisecond},
		{x: "2ms", want: 2 * time.Millisecond},
		{x: "1s", want: 1 * time.Second},
		{x: "2s", want: 2 * time.Second},
		{x: "90s", want: 90 * time.Second},
		{x: "1m30s", want: 90 * time.Second},
		{x: "1m0s", want: 1 * time.Minute},
		{x: "2m0s", want: 2 * time.Minute},
		{x: "90m", want: 90 * time.Minute},
		{x: "1h30m", want: 90 * time.Minute},
		{x: "1h30m0s", want: 90 * time.Minute},
		{x: "1h0m0s", want: 1 * time.Hour},
		{x: "2h0m0s", want: 2 * time.Hour},
		{x: "36h0m0s", want: 36 * time.Hour},
		{x: 0.000000001, want: 1 * time.Nanosecond},
		{x: 0.000001, want: 1 * time.Microsecond},
		{x: 0.001, want: 1 * time.Millisecond},
		{x: 0.1, want: 100 * time.Millisecond},
		{x: 1, want: 1 * time.Second},
		{x: 1.0, want: 1 * time.Second},
		{x: 2.0, want: 2 * time.Second},
		{x: 90, want: 90 * time.Second},
		{x: 90.001, want: (90 * time.Second) + (1 * time.Millisecond)},
		{x: 90.999, want: (90 * time.Second) + (999 * time.Millisecond)},
		{name: "nil", x: nil, wantErr: "time: invalid duration <nil>"},
		{name: "empty string", x: "", wantErr: "time: invalid duration \"\""},
		{x: "'2ms'", wantErr: "time: invalid duration \"'2ms'\""},
		{x: "nil", wantErr: "time: invalid duration \"nil\""},
		{x: "foo", wantErr: "time: invalid duration \"foo\""},
		{x: "\"foo\"", wantErr: "time: invalid duration \"\\\"foo\\\"\""},
		{x: "null", wantErr: "time: invalid duration \"null\""},
		{x: "\"null\"", wantErr: "time: invalid duration \"\\\"null\\\"\""},
	}
	for _, tt := range tests {
		name := tt.name
		if name == "" {
			name = fmt.Sprintf("%#v", tt.x)
		}

		t.Run(name, func(t *testing.T) {
			got, err := Parse(tt.x)

			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.want, time.Duration(got))
		})
	}
}
