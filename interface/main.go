package main

import (
	"bufio"
	"fmt"
	"io"
)

var CustomWriter io.Writer = ABC{}

type ABC struct{}

func (ABC) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return 0, nil
}

func main() {
	CustomWriter.Write([]byte("sdasa"))
	bufio.NewWriter(CustomWriter)
}
