package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	contents string
	pos      int
}

func main() {

	out, err := io.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output: %s", out)
}
