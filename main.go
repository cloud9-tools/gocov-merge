package main

import (
	"os"
	"encoding/json"
	"io"
	"log"

	"github.com/axw/gocov"
)

type file struct {
	Packages []*gocov.Package
}

func main() {
	dec := json.NewDecoder(os.Stdin)
	var in, out file
	var eof bool
	for !eof {
		in = file{}
		err := dec.Decode(&in)
		if err == io.EOF {
			err = nil
			eof = true
		}
		if err != nil {
			log.Fatalf("failed to parse JSON: %v", err)
		}
		out.Packages = append(out.Packages, in.Packages...)
	}
	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(out)
	if err != nil {
		log.Fatalf("failed to write output: %v", err)
	}
}
