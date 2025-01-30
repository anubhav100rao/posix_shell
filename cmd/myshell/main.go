package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handleInvalidCommand(invalid_command string) {
	fmt.Printf("%s: command not found\n", invalid_command)
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

func handleCommand(command string) {
	args := strings.Split(command, " ")
	switch args[0] {
	case "exit":
		handleExit(args)
	default:
		handleInvalidCommand(command)
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
