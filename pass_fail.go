package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade : ")
	reader := bufio.NewReader(os.Stdin)
	inuput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inuput = strings.TrimSpace(inuput)
	grade, err := strconv.ParseFloat(inuput, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of ", grade, " is ", status)
}
