package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("bytecode.txt")

	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	n := 0
	for {
		tmp, err := f.Read(buf)
		if tmp == 0 {
			break
		}
		if err != nil {
			break
		}
		n += tmp
	}
	fmt.Println((string(buf[:n])))
}
