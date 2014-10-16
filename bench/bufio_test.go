package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"
)

const str = "Go, The standard library."
const Times = 100

// openFile opens a file for writing or creates it if it does not exist.
func openFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed opening %s for writing: %s", name, err)
	}
	return file
}

// This routine benchmarks the buffered io writer wrapper over the file I/O
// writer.
func BenchmarkBufio(b *testing.B) {
	file := openFile(os.DevNull)
	defer file.Close()

	bufferedFile := bufio.NewWriter(file)
	for i := 0; i < b.N; i++ {
		if _, err := bufferedFile.WriteString(str); err != nil {
			log.Fatalf("failed or short write: %s", err)
		}
	}
	// What happens if we forget this line?
	// Flush to disk.
	bufferedFile.Flush()
}

// This routine benchmarks the actual I/O done on disk.
func BenchmarkIo(b *testing.B) {
	file := openFile(os.DevNull)
	defer file.Close()

	for i := 0; i < b.N; i++ {
		// This is an actual kernel system call and doing it so many times
		// is relevant for streaming but not when doing regular file I/O.
		if _, err := io.WriteString(file, str); err != nil {
			log.Fatalf("failed or short write: %s", err)
		}
	}
}
