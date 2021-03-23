package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	var slicePersons []person

	var fileName string
	fmt.Println("Enter name of the file to read: ")
	fmt.Scan(&fileName)
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		sline := string(line)
		slineSplit := strings.Split(sline, " ")
		if len(slineSplit[0]) > 20 || len(slineSplit[1]) > 20 {
			fmt.Println("Length of first or last name is more than 20 characters")
			os.Exit(1)
		}
		slicePersons = append(slicePersons, person{slineSplit[0], slineSplit[1]})
	}

	for i, v := range slicePersons {
		fmt.Printf("Name of person %s:\n", strconv.Itoa(i+1))
		fmt.Printf("first name: %s\n", v.fname)
		fmt.Printf("last name: %s\n", v.lname)
	}

}
