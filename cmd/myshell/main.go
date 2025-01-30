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
	fmt.Fprint(os.Stdout, "$ ")

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		input = input[:len(input)-1]
		handleInvalidCommand(input)
	}
}
