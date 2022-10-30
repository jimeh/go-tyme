package dur_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jimeh/go-tyme/dur"
	"gopkg.in/yaml.v3"
)

func ExampleDuration_MarshalJSON() {
	type Connection struct {
		Timeout dur.Duration `json:"timeout"`
	}

	conn := Connection{Timeout: dur.Duration(5 * time.Second)}
	b, _ := json.Marshal(conn)

	fmt.Println(string(b))
	// Output:
	// {"timeout":"5s"}
}

func ExampleDuration_UnmarshalJSON() {
	type Connection struct {
		Timeout dur.Duration `json:"timeout"`
	}

	conn := Connection{}
	_ = json.Unmarshal([]byte(`{"timeout": "10s"}`), &conn)
	fmt.Printf("%+v (%+v)\n", conn.Timeout, time.Duration(conn.Timeout))
	_ = json.Unmarshal([]byte(`{"timeout": 5}`), &conn)
	fmt.Printf("%+v (%+v)\n", conn.Timeout, time.Duration(conn.Timeout))
	_ = json.Unmarshal([]byte(`{"timeout": 0.5}`), &conn)
	fmt.Printf("%+v (%+v)\n", conn.Timeout, time.Duration(conn.Timeout))
	// Output:
	// 10000000000 (10s)
	// 5000000000 (5s)
	// 500000000 (500ms)
}

func ExampleDuration_MarshalYAML() {
	type Connection struct {
		Timeout dur.Duration `yaml:"timeout"`
	}

	conn := Connection{Timeout: dur.Duration(5 * time.Second)}
	b, _ := yaml.Marshal(conn)

	fmt.Println(string(b))
	// Output:
	// timeout: 5s
}

func ExampleDuration_UnmarshalYAML() {
	type Connection struct {
		Timeout dur.Duration `yaml:"timeout"`
	}

	conn := Connection{}
	_ = yaml.Unmarshal([]byte(`timeout: 10s`), &conn)
	fmt.Printf("%+v (%+v)\n", conn.Timeout, time.Duration(conn.Timeout))
	_ = yaml.Unmarshal([]byte(`timeout: 5`), &conn)
	fmt.Printf("%+v (%+v)\n", conn.Timeout, time.Duration(conn.Timeout))
	_ = yaml.Unmarshal([]byte(`timeout: 0.5`), &conn)
	fmt.Printf("%+v (%+v)\n", conn.Timeout, time.Duration(conn.Timeout))
	// Output:
	// 10000000000 (10s)
	// 5000000000 (5s)
	// 500000000 (500ms)
}
