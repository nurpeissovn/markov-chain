package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	WordSlice []string
	Str       []string
	Result    string
)

func main() {
	input := os.Args[1:]

	if len(input) == 0 {
		Result += "of the"
		read()
		data(2)
		chain(Result)
		fmt.Println(len(Result), Result)
		for i := 0; i < 20; i++ {
			Str = []string{}
			data(len(strings.Fields(Result)) - 1)
			chain(Result)

			// break
			fmt.Println(len(strings.Fields(Result)))

		}
	}
}

func read() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		sentence := scanner.Text()
		words := strings.Fields(sentence)

		WordSlice = append(WordSlice, words...)
	}
}

func data(given int) {
	temp := ""

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

	}
	fmt.Println(Str[0], "tekesris")
}

func chain(prefix string) {
	var tempSlice []string
	temp := []string{}
	rndm := 0

	for i, k := range Str {
		if k == prefix {
			rndm++
			temp = strings.Fields(Str[i+1])
			tempSlice = append(tempSlice, temp[1])

			temp = []string{}
			// fmt.Println(prefix)
		}
	}
	if rndm != 0 {
		Result += " " + tempSlice[rand.Intn(rndm)]
	}
}
