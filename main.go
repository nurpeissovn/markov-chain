package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	WordSlice  []string
	Str        []string
	Result     []string
	TotalWords int
)

func main() {
	input := os.Args[1:]
	read()

	if len(input) == 0 {
		Result = append(Result, WordSlice[:2]...)
		for {
			Str = []string{}
			makeSentence(len(Result))
			chain(strings.Join(Result, " "), 100)

			if len(Result) == 100 {
				break
			}

		}
		// fmt.Println(Result)

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

func makeSentence(given int) {
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

}

func chain(prefix string, shek int) {
	var tempSlice []string
	temp := []string{}
	rndm := 0
	// fmt.Println(TotalWords, len(WordSlice))
	if len(strings.Fields(Str[0])) == len(WordSlice) {
		os.Exit(1)
	}
	for i, k := range Str {
		if k == prefix {
			fmt.Println(i, len(strings.Fields(Str[0])), Str[0])

			rndm++
			if len(strings.Fields(Str[0])) < shek {
				temp = strings.Fields(Str[i+1])
				tempSlice = append(tempSlice, temp[len(temp)-1])
				fmt.Println(temp, len(temp))
				temp = []string{}

			}

		}

	}

	if rndm != 0 {
		Result = append(Result, tempSlice[rand.Intn(rndm)])
	}
}
