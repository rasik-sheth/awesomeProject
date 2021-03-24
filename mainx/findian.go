package mainx

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func findian() {

	fmt.Println("\nEnter text to check findian:")

	//reading a line as scan is only getting tokens
	var stdin *bufio.Reader
	var line []rune
	var c rune
	var err error

	stdin = bufio.NewReader(os.Stdin)
	for {
		c, _, err = stdin.ReadRune()
		if err == io.EOF || c == '\n' {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading standard input\n")
			os.Exit(1)
		}
		line = append(line, c)
	}

	sline := string(line[:len(line)])

	//validate content
	if sline != "" {
		if len(sline) >= 3 && strings.HasPrefix(strings.ToLower(sline), "i") &&
			strings.HasSuffix(strings.ToLower(sline), "n") &&
			strings.Contains(strings.ToLower(sline), "a") {
			fmt.Printf("Found!")
		} else {
			fmt.Printf("Not Found!")
		}
	} else {
		fmt.Printf("** Not a valid content **")
	}
}
