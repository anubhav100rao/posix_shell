package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func handleInvalidCommand(invalid_command string) {
	fmt.Printf("%s: command not found\n", invalid_command)
}

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	handleInvalidCommand(input)
}
