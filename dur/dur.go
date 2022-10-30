// Package dur provides a wrapper dur.Duration around time.Duration, with
// sensible JSON/YAML marshaling/unmarshaling support.
//
// Unmarshaling supports standard time.Duration formats string formats such as
// "5s, ""1h30m", all parsed by time.ParseDuration. It also supports integer and
// float values which are interpreted as seconds, rather than nanoseconds, like
// the regular time.Duration does.
//
// Marshaling always outputs a string, using the standard time.Duration format,
// by calling time.Duration(d).String().
package dur
