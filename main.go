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

	parts := strings.Split(trimmed, " ")

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if parts[0] == "cd" && len(parts) == 1 {
		os.Chdir("/")
	}
	var err error
	switch {
	case parts[0] == "exit":
		os.Exit(0)
	case parts[0] == "cd" && len(parts) == 1:
		err = os.Chdir("/")
	case parts[0] == "cd":
		err = os.Chdir(parts[1])
	default:
		err = cmd.Run()
	}

	return err
}

func main() {
	for {
		fmt.Print("> ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
