package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var BUILTIN_COMMANDS = [...]string{"exit", "echo", "type", "pwd"}

func searchCommandInDirectory(directory string, command string) (string, bool) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return "", true
	}
	for _, file := range files {
		if file.Name() == command {
			return directory, false
		}
	}
	return "", true
}

func handleUnixCommand(builtin string) (string, bool) {
	path := os.Getenv("PATH")
	executableDirs := strings.Split(path, ":")
	for _, dir := range executableDirs {
		directory, err := searchCommandInDirectory(dir, builtin)
		if !err {
			completePath := directory + "/" + builtin
			return completePath, false
		}
	}
	return "", true
}

func handleInvalidCommand(invalid_command string, default_message ...string) {
	err_message := "command not found"
	if len(default_message) > 0 {
		err_message = default_message[0]
	}
	fmt.Printf("%s: %s\n", invalid_command, err_message)
}

func handleExit(args []string) {
	if len(args) < 2 {
		log.Fatal("exit: too less arguments")
	}
	status, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("exit: invalid status code")
	}
	os.Exit(status)
}

func handleEcho(args []string) {
	if len(args) < 2 {
		log.Fatal("echo: too less arguments")
	}
	// if args[1][0] == '"' || args[1][0] == '\'' {
	// 	fmt.Println(strings.Join(args[1:], " ")[1 : len(strings.Join(args[1:], " "))-1])
	// 	return
	// }

	fmt.Println(strings.Join(args, " "))
}

func handleType(builtin string) {

	for _, command := range BUILTIN_COMMANDS {
		if command == builtin {
			fmt.Printf("%s is a shell builtin\n", builtin)
			return
		}
	}

	completePath, err := handleUnixCommand(builtin)
	if !err {
		fmt.Printf("%s is %s\n", builtin, completePath)
		return
	}
	handleInvalidCommand(builtin, "not found")
}

func handleExecutables(command string) {
	args := strings.Split(command, " ")
	terminalCommand := exec.Command(args[0], args[1:]...)
	terminalCommand.Stderr = os.Stderr
	terminalCommand.Stdout = os.Stdout
	err := terminalCommand.Run()
	if err != nil {
		handleInvalidCommand(args[0])
	}
}

func handlePWDCommand() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
}

func handleCDCommand(args []string) {
	defaultDir := "~"
	if len(args) > 1 {
		defaultDir = args[1]
	}

	if defaultDir == "~" {
		homeDir, _ := os.UserHomeDir()
		defaultDir = homeDir
	}

	err := os.Chdir(defaultDir)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", defaultDir)
	}
}

func handleCommand(command string) {
	var args []string
	seenQuote := false
	temp := ""
	for _, ch := range command {
		if ch == '"' || ch == '\'' {
			seenQuote = !seenQuote
			if temp != "" {
				args = append(args, temp)
				temp = ""
			}
		} else {
			if seenQuote {
				temp += string(ch)
			} else if ch == ' ' {
				if temp != "" {
					args = append(args, temp)
					temp = ""
				}
			} else {
				temp += string(ch)
			}
		}
	}
	if temp != "" {
		args = append(args, temp)
	}

	switch args[0] {
	case "exit":
		handleExit(args)
	case "echo":
		handleEcho(args)
	case "type":
		handleType(strings.Join(args[1:], " "))
	case "pwd":
		handlePWDCommand()
	case "cd":
		handleCDCommand(args)
	default:
		handleExecutables(command)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		input = input[:len(input)-1]
		handleCommand(input)
	}
}
