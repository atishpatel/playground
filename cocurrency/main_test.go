package main

import "testing"

func TestCocurrency(t *testing.T) {

}

func TestGoFunc(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
	}{
		{
			desc:  "Go function should handle test1",
			input: "test1",
		},
		{
			desc:  "Go function should handle test2",
			input: "test2",
		},
		{
			desc:  "Go function should fail at test3",
			input: "test3",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			Go(func() {

				if tC.input == "test3" {
					panic("failed test 3")
				}
			})
		})
	}
}
