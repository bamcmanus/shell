package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read: %v", err)
		}
		if strings.TrimSpace(str) == "exit" {
			break
		} else {
			fmt.Printf("read string '%s'", str)
		}
	}
}
