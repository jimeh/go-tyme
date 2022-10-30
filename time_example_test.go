package tyme_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jimeh/go-tyme"
	"gopkg.in/yaml.v3"
)

func ExampleTime_MarshalJSON() {
	type Order struct {
		Date tyme.Time `json:"date"`
	}
	t := time.Date(2006, 1, 2, 15, 4, 5, 999000000, time.UTC)
	order := Order{Date: tyme.Time(t)}
	b, _ := json.Marshal(order)

	fmt.Println(string(b))
	// Output:
	// {"date":"2006-01-02T15:04:05.999Z"}
}

func ExampleTime_MarshalYAML() {
	type Order struct {
		Date tyme.Time `yaml:"date"`
	}
	t := time.Date(2006, 1, 2, 15, 4, 5, 999000000, time.UTC)
	order := Order{Date: tyme.Time(t)}
	b, _ := yaml.Marshal(order)

	fmt.Println(string(b))
	// Output:
	// date: 2006-01-02T15:04:05.999Z
}
