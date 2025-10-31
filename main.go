package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if strings.TrimSpace(str) == "exit" {
			break
		} else {
			fmt.Printf("read string '%s'", str)
		}
	}
}
