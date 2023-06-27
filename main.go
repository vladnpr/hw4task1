package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	findInText(readText(), "dummy")
}

func readText() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var text []string

	for scanner.Scan() {
		textLine := scanner.Text()
		if textLine == "stop" {
			break
		}
		text = append(text, textLine)
	}

	return text
}

func findInText(haystack []string, needle string) {
	var result []int
	var i int
	var v string

	for i, v = range haystack {
		if strings.Contains(v, needle) {
			result = append(result, i)
		}
	}

	fmt.Printf("The word \"%s\" appears %d times in the text at rows: %d", needle, i, result)

	students()
}

func students() {
	evaluations := []float32{10.0, 5.5, 2.2, 9.2}
	var sum float32

	for _, val := range evaluations {
		sum += val
	}
	result := sum / float32(len(evaluations))
	fmt.Printf("\ngrade average: %.2f", result)
}
