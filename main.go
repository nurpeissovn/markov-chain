package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	WordSlice []string
	Sentence  []string
	Result    map[string][]string
	Prefix    string
	Input     []string
)

func printResult(prefix []string, wordCount uint) {
	rand.Seed(time.Now().UnixNano())
	i := 0
	Sentence = append(Sentence, prefix...)
	for {
		temp := strings.Join(Sentence[i:], " ")
		i++
		if _, exists := Result[temp]; !exists || uint(len(Sentence)) == wordCount {
			break
		}

		Sentence = append(Sentence, (Result[temp][rand.Intn(len(Result[temp]))]))

	}

	for i, k := range Sentence {
		if wordCount == 1 && len(WordSlice) <= 2 {
			fmt.Print(WordSlice[0])
			break
		}
		fmt.Print(k)
		if i < len(Sentence)-1 {
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

func chain(wordlen int) {
	Result = make(map[string][]string)
	for i := range WordSlice {
		if i < len(WordSlice)-wordlen {
			tempSentence := strings.Join(WordSlice[i:i+wordlen], " ")

			Result[tempSentence] = append(Result[tempSentence], WordSlice[i+wordlen])

		} else if len(WordSlice) == len(Input) {
			printResult(WordSlice, 2)
			os.Exit(0)
		} else if len(strings.Fields(Prefix)) == len(WordSlice) {
			printResult(WordSlice, uint(len(strings.Fields(Prefix))))
			os.Exit(0)
		}

	}
	if _, exists := Result[Prefix]; !exists {
		fmt.Println("words not found in the text!")
		os.Exit(0)
	}
}

func help() {
	fmt.Print("Markov Chain text generator.", "\n", "\n")
	fmt.Print("Usage:", "\n", "markovchain [-w <N>] [-p <S>] [-l <N>]", "\n", "markovchain --help", "\n", "\n")
	fmt.Println("Options:", "\n", "--help   Show this screen.", "\n", "-w N   Number of maximum words", "\n", "-p S   Starting prefix", "\n", "-l N   Prefix length")
}

func main() {
	Input = os.Args[1:]
	var wordCount uint
	var preLen uint

	if len(Input) != 0 && Input[0] == "--help" || wordCount > 10000 {
		help()
		return
	}

	flag.UintVar(&wordCount, "w", 100, "Enter valid numbers")
	flag.StringVar(&Prefix, "p", "", "")
	flag.UintVar(&preLen, "l", 2, "Enter valid numbers")

	flag.Parse()

	read()
	if wordCount > 0 && len(Input) == 0 && len(WordSlice) >= 2 && preLen >= 2 {
		if wordCount > 10000 {
			fmt.Println("Entered number is too much")
			return
		}
		Prefix += strings.Join(WordSlice[:2], " ")
		chain(2)
		printResult(WordSlice[:2], wordCount)

	} else if len(WordSlice) >= 2 && wordCount > 0 && preLen >= 2 {
		if wordCount > 10000 {
			fmt.Println("Entered number is too much")
			return
		} else if len(WordSlice) == 2 && wordCount == 1 {
			fmt.Println(WordSlice[0])
			return
		} else if wordCount == 1 {
			fmt.Println("Error: minimum should be 2")
			return
		} else if Prefix != "" && preLen > 0 && preLen < 6 && len(Prefix) > 2 {
			if len(strings.Fields(Prefix)) > 5 || wordCount == 0 || wordCount > 10000 || len(strings.Fields(Prefix)) != int(preLen) {
				fmt.Println("Error : prefix is above limit or invalid number")
			} else {
				chain(len(strings.Fields(Prefix)))
				printResult(strings.Fields(Prefix), wordCount)

			}

		} else if preLen < 6 && preLen >= 2 {
			chain(len(strings.Fields(Prefix)))
			printResult(strings.Fields(Prefix), wordCount)
		}

	} else if len(Prefix) == 0 {
		chain(len(strings.Fields(Prefix)))
		printResult(strings.Fields(Prefix), wordCount)

	} else if len(WordSlice) == 0 {
		fmt.Println("Provided file is empty")
	} else if len(WordSlice) == 1 {
		fmt.Println("Provided file consists only one word")
	} else {
		help()
	}
}
