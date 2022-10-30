package ts

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnixSecond(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		i, err := strconv.ParseInt(tt.second, 10, 64)
		if err != nil {
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			got := UnixSecond(i)

			assert.IsType(t, Second(time.Time{}), got)

			want := tt.t.Truncate(time.Second)
			assert.Equal(t, want.UTC(), time.Time(got).UTC())
		})
	}
}

func TestUnixMilli(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		i, err := strconv.ParseInt(tt.millisecond, 10, 64)
		if err != nil || millisecondSkipTestCase(t, tt.t) {
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			got := UnixMilli(i)

			assert.IsType(t, Millisecond(time.Time{}), got)

			want := tt.t.Truncate(time.Millisecond)
			assert.Equal(t, want.UTC(), time.Time(got).UTC())
		})
	}
}

func TestUnixMicro(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		i, err := strconv.ParseInt(tt.microsecond, 10, 64)
		if err != nil || microsecondSkipTestCase(t, tt.t) {
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			got := UnixMicro(i)

			assert.IsType(t, Microsecond(time.Time{}), got)

			want := tt.t.Truncate(time.Microsecond)
			assert.Equal(t, want.UTC(), time.Time(got).UTC())
		})
	}
}

func TestUnixNano(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		i, err := strconv.ParseInt(tt.nanosecond, 10, 64)
		if err != nil || nanosecondSkipTestCase(t, tt.t) {
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			got := UnixNano(i)

			assert.IsType(t, Nanosecond(time.Time{}), got)

			want := tt.t.Truncate(time.Nanosecond)
			assert.Equal(t, want.UTC(), time.Time(got).UTC())
		})
	}
}
