package prompt

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestInt(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		min    int
		max    int
		answer int
	}{
		{
			desc:   "Should read int 5",
			input:  "5\n",
			min:    0,
			max:    10,
			answer: 5,
		},
		{
			desc:   "Should ignore character",
			input:  "p\n5\n",
			min:    0,
			max:    10,
			answer: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := bufio.NewReader(strings.NewReader(tC.input))
			w := new(bytes.Buffer)
			// w := os.Stdout
			prompt := New(w, r)
			v := prompt.Int("Enter a number: ", tC.min, tC.max)
			if tC.answer != v {
				t.Logf("Failing info: %s", w.String())
				t.Fail()
			}
		})
	}
}
