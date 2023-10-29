package main

import (
	"bufio"
	"fmt"
	"os"
)

type InputBuffer struct {
	buffer       *string
	bufferLength int
	inputLength  int
}

func newInputBuffer() *InputBuffer {
	return &InputBuffer{}
}

func printPrompt() {
	fmt.Print("db > ")
}

func readInput(inputBuffer *InputBuffer) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	inputBuffer.buffer = &input
	inputBuffer.inputLength = len(input)
}

func main() {
	inputBuffer := newInputBuffer()

	for {
		printPrompt()
		readInput(inputBuffer)
		if *inputBuffer.buffer == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
			break
		} else {
			fmt.Printf("Unrecognized command '%s'.\n", *inputBuffer.buffer)
		}
	}
}
