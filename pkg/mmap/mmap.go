package mmap

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const (
	MaxMapSize = 0xFFFFFFFFFFFF
)

func Map(path string) ([]byte, *[MaxMapSize]byte, int, error) {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, nil, 0, err
	}

	info, err := file.Stat()
	if err != nil {
		return nil, nil, 0, fmt.Errorf("mmap stat error: %s", err)
	}

	sz := int(info.Size())

	b, err := syscall.Mmap(int(file.Fd()), 0, sz, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return nil, nil, 0, err
	}

	// Pages in the given range can be aggressively read ahead, and may be freed soon after they are accessed
	if err := madvise(b, syscall.MADV_SEQUENTIAL); err != nil {
		return nil, nil, 0, fmt.Errorf("madvise: %s", err)
	}

	return b, (*[MaxMapSize]byte)(unsafe.Pointer(&b[0])), sz, nil
}

func madvise(b []byte, advice int) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_MADVISE, uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), uintptr(advice))
	if e1 != 0 {
		err = e1
	}
	return
}

func Unmap(ref []byte) error {
	if ref == nil {
		return nil
	}

	return syscall.Munmap(ref)
}
