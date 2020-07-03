package prompt

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Prompter struct {
	reader    *bufio.Scanner
	writer    io.Writer
	Tolerance int
}

func New(w io.Writer, r *bufio.Scanner) *Prompter {
	return &Prompter{
		reader:    r,
		writer:    w,
		Tolerance: 20,
	}
}

func (p *Prompter) Int(request string, min, max int) int {
	v := 0
	for count := 0; count < p.Tolerance; count++ {
		fmt.Fprintf(p.writer, "%s", request)
		ok := p.reader.Scan()
		if !ok {
			if p.reader.Err() != nil {
				fmt.Fprintf(p.writer, "ReadError: %s\n", p.reader.Err())
			}
			break
		}
		s := p.reader.Text()
		i, err := strconv.ParseInt(strings.TrimSuffix(s, "\n"), 10, 64)
		if err != nil {
			fmt.Fprintf(p.writer, "ParseError: %s\n", err.Error())
			continue
		}
		v = int(i)
		if v < min {
			fmt.Fprintf(p.writer, "Value must be larger than %d.\n", min)
			continue
		}
		if v > max {
			fmt.Fprintf(p.writer, "Value must be smaller than %d.\n", max)
			continue
		}
		break
	}
	return v
}
