package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

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
	outputFile, err := os.Create("output-no-generics.jl")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(inputFile)

	// Create a stack to store the JSON objects
	var stack []string

	// Read each line of the input file
	for scanner.Scan() {
		// Read the line
		line := scanner.Text()

		// Add the line to the stack
		stack = append(stack, line)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning input file:", err)
		return
	}

	// Write the JSON objects in reverse order to the output file
	for i := len(stack) - 1; i >= 0; i-- {
		// Write the JSON object to the output file
		_, err := fmt.Fprintln(outputFile, stack[i])
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
