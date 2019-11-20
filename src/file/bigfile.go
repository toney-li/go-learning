package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var countGuard sync.Mutex
var count = 0
var keyWord = "jdbc."

func main() {
	start := time.Now()
	var file = "/Users/liyi/Downloads/mysql.txt"
	err := readLineFile(file, handle)
	if err != nil {
		fmt.Printf("Read file=%s failed \n", file)
	}
	end := time.Now()
	fmt.Printf("%s count %d \n", keyWord, count)
	fmt.Printf("Count elapsed time : %f ms", end.Sub(start).Seconds())
}

func readBinaryFile(file string, handle func([]byte)) error {
	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	s := make([]byte, 4096)
	for {
		n, err := f.Read(s)
		// handle(string(s[:n]))
		if n == 0 || (err != nil && err == io.EOF) {
			return nil
		}
	}
}

func handle(lines []string) {
	// countGuard.Lock()
	// countGuard.Unlock()
	for _, line := range lines {
		c := strings.Count(string(line), keyWord)
		count += c
	}
}

func readLineFile(file string, handle func([]string)) error {
	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	maxLine := 1000
	lines := make([]string, maxLine)
	idx := 0

	for {
		line, _, err := buf.ReadLine()
		if idx >= maxLine {
			handle(lines)
			idx = 0
		}
		lines[idx] = string(line)
		idx++
		if err != nil && err == io.EOF {
			break
		}
	}
	if idx > 0 {
		handle(lines[0:idx])
	}
	return nil
}
