package main

import (
	"fmt"
	"strings"
)

func main() {
	var text = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. \n" +
		"Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took \n" +
		"a galley of type and scrambled it to make a type specimen book. \n" +
		"It has survived not only five centuries, but also the leap into electronic typesetting, \n" +
		"remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset \n" +
		"sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum"

	findInText(text, "dummy")
}

func findInText(haystack string, needle string) {
	textSlice := strings.Split(haystack, " ")
	var result []int
	var i int
	var v string

	for i, v = range textSlice {
		if strings.Contains(v, needle) {
			result = append(result, i)
		}
	}

	fmt.Printf("The word \"%s\" appears %d times in the text at positions: %d", needle, i, result)

	students()
}

func students() {
	evaluations := []float32{10.0, 5.5, 2.2, 9.2}
	var sum float32

	for _, v := range evaluations {
		sum += v
	}
	result := sum / float32(len(evaluations))
	fmt.Printf("\ngrade average: %.2f", result)
}
