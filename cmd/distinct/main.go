package main

import (
	"bytes"
	"flag"
	"hash/maphash"
	"log"
	"time"

	"github.com/moretestingtasks/ecwid/pkg/hyperloglog"
	"github.com/moretestingtasks/ecwid/pkg/mmap"
)

func main() {
	var (
		input string
		hash  maphash.Hash
	)

	flag.StringVar(&input, "input", "", "Input filepath")
	flag.Parse()

	st := time.Now()
	log.Printf("Processing '%s'...\n", input)

	// TODO. Close file descriptor!
	ref, data, size, err := mmap.Map(input)
	if err != nil {
		log.Fatalf("Error mapping file '%s' - %s\n", input, err)
	}
	defer func() {
		if ref != nil {
			mmap.Unmap(ref)
		}
	}()

	plus, err := hyperloglog.NewPlus(16)
	if err != nil {
		log.Fatalf("Error creating HLL++ instance - %s\n", err)
	}

	n := 0
	for n < size {
		idx := bytes.IndexByte(data[n:], '\n')
		if idx == -1 {
			log.Fatal("Expected idx > 0, delimeter not found")
		}

		hash.Write(data[n : idx+n])

		plus.Add(hyperloglog.Hash64(&hash))

		hash.Reset()

		n += idx + 1
	}

	log.Printf("Distinct count: %v\n", plus.Count())
	log.Printf("Done in %v\n", time.Now().Sub(st))
}
