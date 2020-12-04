package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune

	if info.Mode() & os.ModeCharDevice != 0 {
		print([]rune("Pipe me!"))
		os.Exit(0)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	
	print(output)

}

func print(text []rune) {
	for i, c := range text {
		r, g, b := rgb(i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, c)
	}
}

func rgb(i int) (int, int, int) {
	f := 0.1
	return int(math.Sin(f * float64(i)) * 127 + 128 ),
				 int(math.Sin(f * float64(i) + 2) * 127 + 128 ),
				 int(math.Sin(f * float64(i) + 4) * 127 + 128 )
}