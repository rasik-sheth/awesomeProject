package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PersonInfo struct  {
	fname string
	lname string
}

func main() {

		fmt.Print("Enter the file filename to read: ")

		in := bufio.NewReader(os.Stdin)
		filename, err := in.ReadString('\n')
		filename = strings.Trim(filename, "\n\r")


		if(err == nil){
			file, er := os.Open(filename);

			if(er !=nil){
				log.Fatal("File not found")
			}
			 defer file.Close()

			 personInfoSl := make([]PersonInfo,0,0)

			scanner := bufio.NewScanner(file)
			for(scanner.Scan()){
				oneline := scanner.Text();
				if(oneline!=""){
					arr := strings.Split(oneline, " ");
					p:=PersonInfo{fname: arr[0], lname: arr[1]}
					personInfoSl = append(personInfoSl, p)
				}else{
					fmt.Printf("\n## Blank line will be ignored")
				}
			}

			fmt.Print("\n\nCompleted file reading, printing: ")

			for _,v := range personInfoSl{
				fmt.Printf("\n %s %s", v.fname, v.lname )
			}


		}



}
