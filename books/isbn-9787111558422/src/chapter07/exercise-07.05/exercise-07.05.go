package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type LimitedReader struct {
	R io.Reader
	N int64
}

func (r *LimitedReader) Read(p []byte) (int, error) {
	if r.N <= 0 {
		return 0, io.EOF
	}
	
	if int64(len(p)) > r.N {
		p = p[:r.N]
	}

	n, err := r.R.Read(p)
	r.N -= int64(n)

	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func main() {
	r := LimitReader(strings.NewReader("saint-yellow"), 5)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	fmt.Printf("%s\n", b)
}