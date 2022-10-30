package ts

import "time"

var (
	marshalUnmarshalTestCases = []struct {
		name        string
		t           time.Time
		second      string
		millisecond string
		microsecond string
		nanosecond  string
	}{
		{
			name: "UTC-8",
			t: time.Date(
				2022, 10, 29, 6, 40, 34, 934349003,
				time.FixedZone("UTC-8", -8*60*60),
			),
			second:      "1667054434",
			millisecond: "1667054434934",
			microsecond: "1667054434934349",
			nanosecond:  "1667054434934349003",
		},
		{
			name: "UTC-2",
			t: time.Date(
				2022, 10, 29, 12, 40, 34, 934349003,
				time.FixedZone("UTC-2", -2*60*60),
			),
			second:      "1667054434",
			millisecond: "1667054434934",
			microsecond: "1667054434934349",
			nanosecond:  "1667054434934349003",
		},
		{
			name: "UTC",
			t: time.Date(
				2022, 10, 29, 14, 40, 34, 934349003, time.UTC,
			),
			second:      "1667054434",
			millisecond: "1667054434934",
			microsecond: "1667054434934349",
			nanosecond:  "1667054434934349003",
		},
		{
			name: "UTC+2",
			t: time.Date(
				2022, 10, 29, 16, 40, 34, 934349003,
				time.FixedZone("UTC+2", 2*60*60),
			),
			second:      "1667054434",
			millisecond: "1667054434934",
			microsecond: "1667054434934349",
			nanosecond:  "1667054434934349003",
		},
		{
			name: "UTC+8",
			t: time.Date(
				2022, 10, 29, 22, 40, 34, 934349003,
				time.FixedZone("UTC+8", 8*60*60),
			),
			second:      "1667054434",
			millisecond: "1667054434934",
			microsecond: "1667054434934349",
			nanosecond:  "1667054434934349003",
		},
		{
			name: "epoch",
			t: time.Date(
				1970, 1, 1, 0, 0, 0, 0, time.UTC,
			),
			second:      "0",
			millisecond: "0",
			microsecond: "0",
			nanosecond:  "0",
		},
		{
			name:        "min second",
			t:           time.Date(292277026596, 12, 4, 15, 30, 8, 0, time.UTC),
			second:      "-9223372036854775808",
			millisecond: "0",
			microsecond: "0",
			nanosecond:  "0",
		},
		{
			name:        "max second",
			t:           time.Date(292277026596, 12, 4, 15, 30, 7, 0, time.UTC),
			second:      "9223372036854775807",
			millisecond: "-1000",
			microsecond: "-1000000",
			nanosecond:  "-1000000000",
		},
		{
			name: "min millisecond",
			t: time.Date(
				-292275055, 5, 16, 16, 47, 4, 192000000, time.UTC,
			),
			second:      "-9223372036854776",
			millisecond: "-9223372036854775808",
			microsecond: "0",
			nanosecond:  "0",
		},
		{
			name: "max millisecond",
			t: time.Date(
				292278994, 8, 17, 7, 12, 55, 807000000, time.UTC,
			),
			second:      "9223372036854775",
			millisecond: "9223372036854775807",
			microsecond: "-1000",
			nanosecond:  "-1000000",
		},
		{
			name: "min microseconds",
			t: time.Date(
				-290308, 12, 21, 19, 59, 5, 224192000, time.UTC,
			),
			second:      "-9223372036855",
			millisecond: "-9223372036854776",
			microsecond: "-9223372036854775808",
			nanosecond:  "0",
		},
		{
			name: "max microseconds",
			t: time.Date(
				294247, 1, 10, 4, 0, 54, 775807000, time.UTC,
			),
			second:      "9223372036854",
			millisecond: "9223372036854775",
			microsecond: "9223372036854775807",
			nanosecond:  "-1000",
		},
		{
			name:        "min nanoseconds",
			t:           time.Date(1677, 9, 21, 0, 12, 43, 145224192, time.UTC),
			second:      "-9223372037",
			millisecond: "-9223372036855",
			microsecond: "-9223372036854776",
			nanosecond:  "-9223372036854775808",
		},
		{
			name: "max nanoseconds",
			t: time.Date(
				2262, 4, 11, 23, 47, 16, 854775807, time.UTC,
			),
			second:      "9223372036",
			millisecond: "9223372036854",
			microsecond: "9223372036854775",
			nanosecond:  "9223372036854775807",
		},
		{
			name:        "year 1092",
			t:           time.Date(1092, 3, 23, 3, 52, 8, 734829384, time.UTC),
			second:      "-27699912472",
			millisecond: "-27699912471266",
			microsecond: "-27699912471265171",
			nanosecond:  "9193575676153932616",
		},
		{
			name: "year -1000",
			t: time.Date(
				-1000, 10, 29, 22, 40, 34, 934349003,
				time.FixedZone("UTC+8", 8*60*60),
			),
			second:      "-93698068766",
			millisecond: "-93698068765066",
			microsecond: "-93698068765065651",
			nanosecond:  "-1464348396517892917",
		},
		{
			name: "year 10449",
			t: time.Date(
				10449, 10, 29, 22, 40, 34, 934349003,
				time.FixedZone("UTC+8", 8*60*60),
			),
			second:      "267597528034",
			millisecond: "267597528034934",
			microsecond: "267597528034934349",
			nanosecond:  "-9103633070708925237",
		},
		{
			name: "year 1044938",
			t: time.Date(
				1044938, 10, 29, 22, 40, 34, 934349003,
				time.FixedZone("UTC+8", 8*60*60),
			),
			second:      "32912917195234",
			millisecond: "32912917195234934",
			microsecond: "-3980570952184168883",
			nanosecond:  "3925767737094266059",
		},
	}
	unmarshalTestCases = []struct {
		name        string
		second      string
		millisecond string
		microsecond string
		nanosecond  string
		t           time.Time
		wantErr     string
	}{
		{
			name:        "string",
			second:      `"2019-01-01T00:00:00Z"`,
			millisecond: `"2019-01-01T00:00:00Z"`,
			microsecond: `"2019-01-01T00:00:00Z"`,
			nanosecond:  `"2019-01-01T00:00:00Z"`,
			wantErr:     "invalid numeric timestamp",
		},
		{
			name:        "array",
			second:      `[1, "true", false]`,
			millisecond: `[1, "true", false]`,
			microsecond: `[1, "true", false]`,
			nanosecond:  `[1, "true", false]`,
			wantErr:     "invalid numeric timestamp",
		},
		{
			name:        "object",
			second:      `{"object": "Object"}`,
			millisecond: `{"object": "Object"}`,
			microsecond: `{"object": "Object"}`,
			nanosecond:  `{"object": "Object"}`,
			wantErr:     "invalid numeric timestamp",
		},
		{
			name:        "whitespace",
			second:      "  1667054434  ",
			millisecond: "  1667054434934  ",
			microsecond: "  1667054434934349  ",
			nanosecond:  "  1667054434934349003  ",
			t: time.Date(
				2022, 10, 29, 14, 40, 34, 934349003, time.UTC,
			),
		},
		{
			name:        "float",
			second:      "1667054434.123456789",
			millisecond: "1667054434934.123456",
			microsecond: "1667054434934349.123",
			nanosecond:  "",
			t: time.Date(
				2022, 10, 29, 14, 40, 34, 934349003, time.UTC,
			),
		},
	}
)
