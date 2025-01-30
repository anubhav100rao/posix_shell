package main

import (
	"bufio"
	"fmt"
	"os"
)

func handleInvalidCommand(invalid_command string) {
	fmt.Printf("%s: command not found\n", invalid_command)
}

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')

		// if err != nil {
		// 	log.Fatal(err)
		// }
		input = input[:len(input)-1]
		handleInvalidCommand(input)
	}
}
