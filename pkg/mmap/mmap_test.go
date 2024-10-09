package mmap

import (
	"bytes"
	"io/ioutil"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	src, err := ioutil.ReadFile("./mmap.go")
	if err != nil {
		t.Fatal(err)
	}

	fn := func(w *sync.WaitGroup) {
		ref, data, size, err := Map("./mmap.go")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			w.Done()
			if ref != nil {
				Unmap(ref)
			}
		}()

		if size != len(src) {
			t.Fatalf("Expected size: %d, obtained: %d\n", len(src), size)
		}

		n := 0
		for n < size {
			idx := bytes.IndexByte(data[n:], '\n')
			if idx == -1 {
				t.Fatal("Expected idx > 0, delimeter not found")
			}

			chunk := data[n : idx+n]

			if bytes.Compare(src[n:idx+n], chunk) != 0 {
				t.Fatalf("Chunk doesn't equal to src\nExpected:%s\nObtained:%s\n", string(src[n:idx+n]), string(chunk))
			}

			n += idx + 1
		}
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		fn(wg)
	}

	wg.Wait()
}
