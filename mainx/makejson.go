package mainx

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func mainy() {

	fmt.Print("Enter your name: ")

	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')
	name = strings.Trim(name, "\n\r")

	fmt.Print("Enter your address: ")
	addrss := bufio.NewReader(os.Stdin)
	addr, err := addrss.ReadString('\n')
	addr = strings.Trim(addr, "\n\r")

	var nameMap map[string]string
	nameMap = make(map[string]string)
	nameMap[name] = addr

	arr, err := json.Marshal(nameMap)
	if err == nil {
		fmt.Print(string(arr))
	}

}
