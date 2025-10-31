package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	trimmed := strings.TrimSuffix(input, "\n")

	cmd := exec.Command(trimmed)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	fmt.Printf("Command: \n%v\n", cmd)

	return cmd.Run()
}

func main() {
	for {
		fmt.Print("> ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if strings.TrimSpace(input) == "exit" {
			break
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
