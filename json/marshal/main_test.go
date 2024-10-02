package main

import (
	"fmt"
	"testing"
)

// Example structs for testing
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func TestMarshalAll(t *testing.T) {
	type testCase struct {
		inputs   []any
		expected [][]byte
	}

	tests := []testCase{
		{
			[]any{
				User{"001", "Sir Bedevere the Wise", 25},
				User{"002", "Sir Lancelot the Brave", 30},
				Product{"The Grail", 19.99},
			},
			[][]byte{
				[]byte(`{"id":"001","name":"Sir Bedevere the Wise","age":25}`),
				[]byte(`{"id":"002","name":"Sir Lancelot the Brave","age":30}`),
				[]byte(`{"name":"The Grail","price":19.99}`),
			},
		},
		{
			[]any{
				User{"003", "Sir Galahad the Pure", 19},
				Product{"The Holy Hand Grenade of Antioch", 4.89},
				Product{"The Wooden Rabbit", 1.99},
			},
			[][]byte{
				[]byte(`{"id":"003","name":"Sir Galahad the Pure","age":19}`),
				[]byte(`{"name":"The Holy Hand Grenade of Antioch","price":4.89}`),
				[]byte(`{"name":"The Wooden Rabbit","price":1.99}`),
			},
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				[]any{
					User{"001", "Sir Robin the Not-Quite-So-Brave-As-Sir-Lancelot", 25},
					Product{"A Shrubbery", 99.79},
				},
				[][]byte{
					[]byte(`{"id":"001","name":"Sir Robin the Not-Quite-So-Brave-As-Sir-Lancelot","age":25}`),
					[]byte(`{"name":"A Shrubbery","price":99.79}`),
				},
			},
		}...)
	}

	// Running user tests
	for _, test := range tests {
		output, err := marshalAll(test.inputs)
		if err != nil {
			t.Errorf(`User Test Failed: %v
  unexpected error: %v
`, test.inputs, err)
		}
		if len(output) != len(test.expected) {
			t.Errorf(`User Test Failed: %v
  expected length: %v
  actual length:   %v
`, test.inputs, len(test.expected), len(output))
		}
		for i, jsonOutput := range output {
			if string(jsonOutput) != string(test.expected[i]) {
				t.Errorf(`User Test Failed: %v
  expected: %s
  actual:   %s
`, test.inputs[i], test.expected[i], jsonOutput)
			} else {
				fmt.Printf(`User Test Passed: %v
  expected: %s
  actual:   %s
`, test.inputs[i], test.expected[i], jsonOutput)
			}
		}
		fmt.Println("==============================")
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
