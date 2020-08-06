package main

import (
	"fmt"
	"github.com/devemio/docker-color-output/input"
)

func main() {
	chars := input.ReadFakeInput()

	for j := 0; j < len(chars); j++ {
		fmt.Printf("%c", chars[j])
	}
}
