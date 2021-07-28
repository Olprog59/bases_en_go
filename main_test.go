package main

import (
	"bufio"
	"os"
	"testing"
)

const FILE_PATH = "test.txt"
const LINES_COUNT = 100000
const TEXT = "hello\n"

func BenchmarkWriteFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Create(FILE_PATH)
		if err != nil {
			panic(err)
		}

		for i := 0; i < LINES_COUNT; i++ {
			f.WriteString(TEXT)
		}

		f.Close()
	}
}

func BenchmarkWriteFileBuffered(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Create(FILE_PATH)
		if err != nil {
			panic(err)
		}

		w := bufio.NewWriter(f)

		for i := 0; i < LINES_COUNT; i++ {
			w.WriteString(TEXT)
		}

		w.Flush()
		f.Close()
	}
}

//BenchmarkWriteFile-8                   7         158044482 ns/op
//BenchmarkWriteFileBuffered-8        1188           1043201 ns/op
