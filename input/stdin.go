package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadInput() []rune {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("Usage:")
		fmt.Println("  docker ps | dco")
		fmt.Println("  docker images | dco")
		fmt.Println("  docker-compose ps | dco")
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return output
}
