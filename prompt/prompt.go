package prompt

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Prompter struct {
	reader    *bufio.Reader
	writer    io.Writer
	Tolerance int
}

func New(w io.Writer, r *bufio.Reader) *Prompter {
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
		s, err := p.reader.ReadString('\n')
		if err != nil && s == "" {
			if err == io.EOF {
				fmt.Fprintf(p.writer, "EOF Error: %s %s\n", s, err.Error())
				break
			}
			fmt.Fprintf(p.writer, "ReadError: %s\n", err.Error())
			continue
		}
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
