package ts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestSecond_MarshalUnmarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			v := Second(tt.t)

			b, err := json.Marshal(v)
			require.NoError(t, err)

			assert.Equal(t, tt.second, string(b))

			var got Second
			err = json.Unmarshal(b, &got)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Second)

			assert.Equal(t, want.UTC(), time.Time(got).UTC())
			assert.Equal(t, time.Local, time.Time(got).Location())
		})
	}
}

func TestSecond_MarshalUnmarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			v := Second(tt.t)

			b, err := yaml.Marshal(v)
			require.NoError(t, err)

			assert.Equal(t, tt.second+"\n", string(b))

			var got Second
			err = yaml.Unmarshal(b, &got)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Second)

			assert.Equal(t, want.UTC(), time.Time(got).UTC())
			assert.Equal(t, time.Local, time.Time(got).Location())
		})
	}
}

func TestSecond_MarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			ts := Second(tt.t)

			b, err := json.Marshal(ts)
			require.NoError(t, err)

			assert.Equal(t, tt.second, string(b))
		})
	}
}

func TestSecond_MarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			ts := Second(tt.t)

			b, err := yaml.Marshal(ts)
			require.NoError(t, err)

			assert.Equal(t, tt.second+"\n", string(b))
		})
	}
}

func TestSecond_UnmarshalJSON(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var ts Second

			err := json.Unmarshal([]byte(tt.second), &ts)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Second)

			assert.Equal(t, want.UTC(), time.Time(ts).UTC())
		})
	}
	for _, tt := range unmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var ts Second

			err := json.Unmarshal([]byte(tt.second), &ts)

			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				want := tt.t.Truncate(time.Second)
				assert.Equal(t, want.UTC(), time.Time(ts).UTC())
			}
		})
	}
}

func TestSecond_UnmarshalYAML(t *testing.T) {
	for _, tt := range marshalUnmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var ts Second

			err := yaml.Unmarshal([]byte(tt.second), &ts)
			require.NoError(t, err)

			want := tt.t.Truncate(time.Second)

			assert.Equal(t, want.UTC(), time.Time(ts).UTC())
		})
	}
	for _, tt := range unmarshalTestCases {
		t.Run(tt.name, func(t *testing.T) {
			var ts Second

			err := yaml.Unmarshal([]byte(tt.second), &ts)

			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				want := tt.t.Truncate(time.Second)
				assert.Equal(t, want.UTC(), time.Time(ts).UTC())
			}
		})
	}
}
