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

func outputPrompt() {
	wd, err := os.Getwd()		
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting working directory: %s\n", err)
		fmt.Print("> ")
		return
	}
	fmt.Printf("%s> ", wd)
}

func main() {
	for {
		outputPrompt()
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
