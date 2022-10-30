package tyme

import "github.com/araddon/dateparse"

var (
	// RetryAmbiguousDateWithSwap is option available in dateparse. This var
	// controls if Time's unmarshalers enables it or not.
	RetryAmbiguousDateWithSwap = false

	// PreferMonthFirst is option available in dateparse. This var
	// controls if Time's unmarshalers enables it or not.
	PreferMonthFirst = false
)

// Parse is a helper function to parse a wide range of string date and time formats using dateparse.ParseAny.
func Parse(s string) (Time, error) {
	t, err := dateparse.ParseAny(
		s,
		dateparse.RetryAmbiguousDateWithSwap(RetryAmbiguousDateWithSwap),
		dateparse.PreferMonthFirst(PreferMonthFirst),
	)
	if err != nil {
		return Time{}, err
	}

	return Time(t), nil
}
