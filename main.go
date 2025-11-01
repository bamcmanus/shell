package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)


const defaultPrompt = "> "

func execInput(input, homeDir string) error {
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
		err = os.Chdir(homeDir)
	case parts[0] == "cd":
		err = os.Chdir(parts[1])
	default:
		err = cmd.Run()
	}

	return err
}

func outputPrompt(username, homeDir string) {
	wd, err := os.Getwd()		
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting working directory; using default prompt: %s\n", err)
		fmt.Print(defaultPrompt)
		return
	} else {
		fmt.Printf("Got working dir: %s\n", wd)
	}
	relativeWorkDir := strings.Replace(wd, homeDir, "~", 1)
	if relativeWorkDir == "/" {
		relativeWorkDir = "~"
	}

	fmt.Printf("%s|%s > ", username, relativeWorkDir)
}

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting user; using default prompt: %s\n", err)
		fmt.Print(defaultPrompt)
		return
	}

	for {
		outputPrompt(user.Username, user.HomeDir)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input, user.HomeDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
