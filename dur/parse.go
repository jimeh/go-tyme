package dur

import (
	"fmt"
	"time"
)

const floatSecond = float64(time.Second)

// Parse parses given interface to a Duration.
//
// If the interface is a string, it will be parsed using time.Parse. If
// the interface is a int or float64, it will be parsed as a number of seconds.
func Parse(x interface{}) (Duration, error) {
	var d Duration
	switch value := x.(type) {
	case string:
		td, err := time.ParseDuration(value)
		if err != nil {
			return 0, err
		}

		d = Duration(td)
	case float64:
		d = Duration(time.Duration(value * floatSecond))
	case int:
		d = Duration(time.Duration(value) * time.Second)
	default:
		return 0, fmt.Errorf("time: invalid duration %+v", x)
	}

	return d, nil
}
