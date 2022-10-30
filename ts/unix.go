package ts

import "time"

// UnixSecond parses a given int64 as a Unix timestamp with second accuracy.
func UnixSecond(ts int64) Second {
	return Second(time.Unix(ts, 0))
}

// UnixMilli parses a given int64 as a Unix timestamp with millisecond accuracy.
func UnixMilli(ts int64) Millisecond {
	return Millisecond(time.UnixMilli(ts))
}

// UnixMicro parses a given int64 as a Unix timestamp with microsecond accuracy.
func UnixMicro(ts int64) Microsecond {
	return Microsecond(time.UnixMicro(ts))
}

// UnixNano parses a given int64 as a Unix timestamp with nanosecond accuracy.
func UnixNano(ts int64) Nanosecond {
	return Nanosecond(unixNano(ts))
}
