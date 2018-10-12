package main

import (
	// "fmt"
	"testing"
)

func Test_main(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("main panic'd\n")
		}
	}()

	main()
}

func Test_Add(t *testing.T) {
	var tests = []struct {
		name   string
		first  int
		second int
		result int
	}{
		{"1 and 1 make 2", 1, 1, 2},
		{"2 and 2 make 4", 2, 2, 4},
		{"42 and 42 make 84", 42, 42, 84},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Add(test.first, test.second)

			if result != test.result {
				t.Errorf("Bad math:\nwant: %d\nhave: %d\n", test.result, result)
			}
		})
	}
}

func Test_TryStuff(t *testing.T) {
	var tests = []struct {
		name  string
		value bool
		fails bool
	}{
		{"Success", true, false},
		{"Fails", false, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := TryStuff(test.value)

			if err != nil && !test.fails {
				t.Errorf("Unexpected error!")
			}

			if result != test.value {
				t.Errorf("Unexpected result:\nwant: %t\nhave: %t\n", test.value, result)
			}
		})
	}
}
