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
	WordCount uint
	PreLen    uint
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

func isValid() {
	if WordCount < 2 || WordCount > 10000 {
		fmt.Fprintln(os.Stderr, "Words should be more than 2 or less than 10000")
		os.Exit(1)
	}
	if PreLen < 2 || PreLen > 5 {
		fmt.Fprintln(os.Stderr, "Words should be more than 2 or less than 5")
		os.Exit(1)

	}
	if Prefix == "" || len(strings.Fields(Prefix)) > 5 {
		fmt.Fprintln(os.Stderr, "Words should be more than 2 or less than 5")
		os.Exit(1)
	}
	if len(strings.Fields(Prefix)) != int(PreLen) {
		fmt.Fprintln(os.Stderr, "Len of Prefix and length must be equal")
		os.Exit(1)
	}
}

func help() {
	fmt.Print("Markov Chain text generator.", "\n", "\n")
	fmt.Print("Usage:", "\n", "markovchain [-w <N>] [-p <S>] [-l <N>]", "\n", "markovchain --help", "\n", "\n")
	fmt.Println("Options:", "\n", "--help   Show this screen.", "\n", "-w N   Number of maximum words", "\n", "-p S   Starting prefix", "\n", "-l N   Prefix length")
}

func main() {
	Input = os.Args[1:]
	if len(Input) != 0 && Input[0] == "--help"  {
		help()
		return
	} 
	read()
	 if len(WordSlice) <= 2 {
		fmt.Println("Should be more than 2")
		os.Exit(1)
	}

	flag.UintVar(&WordCount, "w", 100, "Enter valid numbers")
	flag.StringVar(&Prefix, "p", strings.Join(WordSlice[:2], " "), "")
	flag.UintVar(&PreLen, "l", 2, "Enter valid numbers")

	flag.Parse()
	isValid()

	if len(Input) == 0 {

		chain(2)
		printResult(WordSlice[:2], WordCount)

	} else {
		chain(len(strings.Fields(Prefix)))
		printResult(strings.Fields(Prefix), WordCount)
	}
}
