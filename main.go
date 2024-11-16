package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	WordSlice []string
	Str       []string
	Result    []string
)

func engine(num, wordlen int) {
	i := 0
	for {
		Str = []string{}
		makeSentence(wordlen)
		chain(strings.Join(Result[i:], " "))
		i++
		if len(Result) == num {
			printResult()
			break
		}

	}
}

func printResult() {
	for i, k := range Result {
		fmt.Print(k)
		if i < len(Result)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func read() {
	scanner := bufio.NewScanner(os.Stdin)
	file, _ := os.Stdin.Stat()
	if (file.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Error: no input text")
		os.Exit(0)
	}

	for scanner.Scan() {
		sentence := scanner.Text()
		words := strings.Fields(sentence)

		WordSlice = append(WordSlice, words...)

	}
}

func makeSentence(given int) {
	temp := ""

	for i := 0; i < len(WordSlice); i++ {
		if i+given <= len(WordSlice) {
			temp = strings.Join(WordSlice[i:i+given], " ")

			Str = append(Str, temp)
			temp = ""
		}
	}
}

func chain(prefix string) {
	var tempSlice []string
	temper := []string{}
	san, rndm := 0, 0
	if len(strings.Fields(Str[0])) == len(WordSlice) || prefix == Str[len(Str)-1] {
		printResult()
		os.Exit(0)
	}
	for i, k := range Str {
		if k == prefix {
			san++
			if i < len(Str)-1 {
				rndm++

				temper = strings.Fields(Str[i+1])
				tempSlice = append(tempSlice, temper[len(temper)-1])
				temper = []string{}

			}

		}
	}

	if rndm != 0 {
		Result = append(Result, tempSlice[rand.Intn(rndm)])
	} else if san == 0 {
		fmt.Println("Words not found")
		os.Exit(0)
	}
}

func help() {
	fmt.Print("Markov Chain text generator.", "\n", "\n")
	fmt.Print("Usage:", "\n", "markovchain [-w <N>] [-p <S>] [-l <N>]", "\n", "markovchain --help", "\n", "\n")
	fmt.Println("Options:", "\n", "--help   Show this screen.", "\n", "-w N   Number of maximum words", "\n", "-p S   Starting prefix", "\n", "-l N   Prefix length")
}

func main() {
	input := os.Args[1:]
	read()
	if len(WordSlice) == 1 {
		fmt.Println(WordSlice[0])
		return
	} else if len(input) == 0 {
		Result = append(Result, WordSlice[:2]...)
		engine(100, 2)

	} else if input[0] == "-w" && len(input) == 2 {
		Result = append(Result, WordSlice[:2]...)

		if num, err := strconv.Atoi(input[1]); err == nil && num < 10001 && num >= 0 {
			engine(num, 2)
		} else if err != nil || num > 10000 || num < 0 {
			fmt.Println("Error: provide valid number.")
			help()
		}
	} else if len(input) == 4 && input[0] == "-w" && input[2] == "-p" && input[3] != "" && len(strings.Fields(input[3])) > 1 {
		if num, err := strconv.Atoi(input[1]); err == nil && num < 10001 && num >= 0 {
			Result = append(Result, strings.Fields(input[3])...)
			engine(num, len(strings.Fields(input[3])))

		} else if err != nil || num > 10000 || num < 0 {
			fmt.Println("Error: provide valid number.")
			help()

		}
	} else if len(input) == 6 && input[0] == "-w" && input[2] == "-p" && input[4] == "-l" {
		num, err := strconv.Atoi(input[1])
		san, errors := strconv.Atoi(input[5])

		if errors != nil || err != nil || san > 5 || num > 10000 || num < 0 || san < 2 || len(strings.Fields(input[3])) < san {
			fmt.Println("Error: provide valid numbers")
		} else if len(strings.Fields(input[3])) >= san {
			slice := strings.Fields(input[3])
			Result = append(Result, slice[:san]...)
			engine(num, san)
		}

	} else {
		help()
	}
}
