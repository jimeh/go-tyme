// Package ts provides wrapper types around time.Time, with support for
// JSON/YAML marshaling to/from numeric Unix timestamps of various precisions.
//
// Unmarshaling supports both numeric and string values. Marshaling always
// produces a integer value.
package ts

import (
	"time"

	"github.com/jimeh/go-tyme/dur"
)

// Timestamp is a type constraint that matches against time.Time, Second,
// Millisecond, Microsecond, and Nanosecond.
type Timestamp interface {
	time.Time | Second | Millisecond | Microsecond | Nanosecond
}

// Duration is a type constraint that matches against time.Duration and
// dur.Duration.
type Duration interface {
	time.Duration | dur.Duration
}

// Add returns a new Timestamp with given Duration added to it, using
// time.Time.Add.
func Add[T Timestamp, D Duration](ts T, d D) T {
	var t time.Time
	t = time.Time(ts)
	t = t.Add(time.Duration(d))

	return T(t)
}

// Sub returns the dur.Duration between two Timestamps, using time.Time.Sub.
func Sub[T, U Timestamp](t T, u U) dur.Duration {
	return dur.Duration(time.Time(t).Sub(time.Time(u)))
}

// After reports whether the Timestamp instant t is after u, using
// time.Time.After.
func After[T, U Timestamp](t T, u U) bool {
	return time.Time(t).After(time.Time(u))
}

// Before reports whether the Timestamp instant t is before u, using
// time.Time.Before.
func Before[T, U Timestamp](t T, u U) bool {
	return time.Time(t).Before(time.Time(u))
}

// Equal reports whether t and u represent the same Timestamp instant, using
// time.Time.Equal.
func Equal[T, U Timestamp](t T, u U) bool {
	return time.Time(t).Equal(time.Time(u))
}

// Round returns the result of rounding t to the nearest multiple of d, using
// time.Time.Round.
func Round[T Timestamp, D Duration](t T, d D) T {
	return T(time.Time(t).Round(time.Duration(d)))
}

// Truncate returns the result of trucating t down to the nearest multiple of d,
// using time.Time.Truncate.
func Truncate[T Timestamp, D Duration](t T, d D) T {
	return T(time.Time(t).Truncate(time.Duration(d)))
}
