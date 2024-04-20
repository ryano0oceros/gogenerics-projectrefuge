package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Stack is a generic stack data structure
type Stack[T any] []T

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() T {
	n := len(*s)
	if n == 0 {
		panic("stack is empty")
	}
	item := (*s)[n-1]
	*s = (*s)[:n-1]
	return item
}

func main() {
	// Start timing
	start := time.Now()

	// Open the input file
	inputFile, err := os.Open("input.jl")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Open the output file
	outputFile, err := os.Create("output-generics.jl")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(inputFile)

	// Create a stack to store the lines
	var stack Stack[string]

	// Read each line of the input file
	for scanner.Scan() {
		// Read the line
		line := scanner.Text()

		// Push the line onto the stack
		stack.Push(line)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning input file:", err)
		return
	}

	// Write the lines in reverse order to the output file
	for len(stack) > 0 {
		// Pop the top line from the stack
		line := stack.Pop()

		// Write the line to the output file
		_, err := fmt.Fprintln(outputFile, line)
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	fmt.Println("Output file created successfully.")

	// Stop timing and print the elapsed time
	elapsed := time.Since(start)
	fmt.Printf("The code took %s to execute.\n", elapsed)
}
