package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var WordSlice []string
var Str []string

func main() {

	input := os.Args[1:]

	if len(input) == 0 {
		read()
		// chain()
	}

}

func read() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		sentence := scanner.Text()
		words := strings.Fields(sentence)

		WordSlice = append(WordSlice, words...)
	}
	temp := ""
	given := 3

	// Map := make(map[string][]string)
	for i := 0; i < len(WordSlice); i++ {

		for j := i; j < given+i; j++ {
			if j < len(WordSlice) {
				temp += strings.ToLower(WordSlice[j])

				if j < given+i-1 && j < len(WordSlice)-1 {
					temp += " "

				}
			}
		}
		Str = append(Str, temp)
		temp = ""
		i += given - 1

	}

	fmt.Println(Str[0])
	fmt.Println(Str[1])
	fmt.Println(Str[3])
	fmt.Println(Str[3])
	fmt.Println(Str[len(Str)-2])
	fmt.Println(Str[len(Str)-1])
	for i := range Str {
		if Str[i] == "my younger and" {

			word := strings.Fields(Str[i+1])
			fmt.Println(Str[i], word[0])

		}
	}

}

func chain() {
	// san := 0
	// str, str2 := "", ""
	// for i := 0; i < len(Str); i++ {
	// 	// index := strings.Index(Str, "in my")
	// 	for j := 0; j < len(Str); j++ {
	// 		if j != len(Str)-3 {
	// 			if Str[i:i+2] == Str[j:j+2] {
	// 				fmt.Println(Str[j : j+3])
	// 			}
	// 		}
	// 	}

	// }

}
