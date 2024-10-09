package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"net"
)

func main() {
	var count int

	flag.IntVar(&count, "count", 1, "count")
	flag.Parse()

	buf := make([]byte, 4)

	for i := 0; i < count; i++ {
		ip := rand.Uint32()
		binary.LittleEndian.PutUint32(buf, ip)
		fmt.Printf("%s\n", net.IP(buf))
	}
}
