// Package tyme provides wrapper types around time.Time, with sensible JSON/YAML
// marshaling/unmarshaling support.
//
// Unmarshaling supports a wide variety of formats, automatically doing a best
// effort to understand the given input, thanks to using the
// github.com/araddon/dateparse package.
//
// Marshaling always produces a string in RFC 3339 format, by simply formatting
// the Time with the time.RFC3339Nano layout.
package tyme
