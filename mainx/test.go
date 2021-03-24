package mainx

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainx() {

	fmt.Print("Enter the string: ")
	in := bufio.NewReader(os.Stdin)
	buff, err := in.ReadString('\n')

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	input := strings.ToLower(strings.Trim(buff, "\n\r"))

	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
