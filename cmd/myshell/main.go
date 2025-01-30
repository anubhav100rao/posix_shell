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
	fmt.Println(strings.Join(args[1:], " "))
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

func handleCommand(command string) {
	args := strings.Split(command, " ")
	switch args[0] {
	case "exit":
		handleExit(args)
	case "echo":
		handleEcho(args)
	case "type":
		handleType(strings.Join(args[1:], " "))
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
