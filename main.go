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
)

// func engine(num, wordlen int) {
//  i := 0
//  for {

//   chain(wordlen)
//   i++
//   if len(Result) == num {
//    printResult()
//    break
//   }

//  }
// }

func printResult(prefix []string, wordCount int) {
	rand.Seed(time.Now().UnixNano())
	i := 0
	Sentence = append(Sentence, prefix...)
	for {
		temp := strings.Join(Sentence[i:], " ")
		i++
		if _, exists := Result[temp]; !exists || len(Sentence) == wordCount {
			break
		}

		Sentence = append(Sentence, (Result[temp][rand.Intn(len(Result[temp]))]))

	}

	for i, k := range Sentence {
		fmt.Print(k)
		if i < len(Sentence)-1 {
			fmt.Print(" ")
		}

	}
	fmt.Println()

	// fmt.Print(" ")
	//
	// }
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

// func makeSentence(given int) {
//  temp := ""

//  for i := 0; i < len(WordSlice); i++ {
//   if i+given <= len(WordSlice) {
//    temp = strings.Join(WordSlice[i:i+given], " ")

//    Sentence = append(Sentence, temp)
//    temp = ""
//   }
//  }
// }

func chain(wordlen int) {
	Result = make(map[string][]string)
	for i := range WordSlice {
		if i < len(WordSlice)-wordlen {
			tempSentence := strings.Join(WordSlice[i:i+wordlen], " ")
			// fmt.Println(tempSentence)

			Result[tempSentence] = append(Result[tempSentence], WordSlice[i+wordlen])

		}
	}
}

func help() {
	fmt.Print("Markov Chain text generator.", "\n", "\n")
	fmt.Print("Usage:", "\n", "markovchain [-w <N>] [-p <S>] [-l <N>]", "\n", "markovchain --help", "\n", "\n")
	fmt.Println("Options:", "\n", "--help   Show this screen.", "\n", "-w N   Number of maximum words", "\n", "-p S   Starting prefix. By default length is 2", "\n", "-l N   Prefix length. By default maximum is 5")
}

func main() {
	input := os.Args[1:]

	wordCount := flag.Int("w", 100, "")
	prefix := flag.String("p", "", "")

	flag.Parse()

	if len(input) != 0 && input[0] == "--help" {
		help()
		return
	}
	read()

	if len(WordSlice) <= 2 {

		fmt.Println("Error: invalid input or file should contain more than 2 words")
		return
	}

	if *wordCount > 0 && *prefix != "" {
		chain(2)
		printResult(strings.Fields(*prefix), *wordCount)

	}
	// if *wordCount == 10 {
	// 	chain(2)
	// 	printResult(WordSlice[:2], *wordCount)

	// }
	// else if len(input) == 0 {
	//  Result = append(Result, WordSlice[:2]...)
	//  engine(100, 2)

	// } else if *wordCount > 0 {
	//  Result = append(Result, WordSlice[:2]...)

	//  if num, err := strconv.Atoi(input[1]); err == nil && num < 10001 && num >= 0 {
	//   engine(num, 2)
	//  }
	// else if err != nil  num > 10000  num < 0 {
	//   fmt.Println("Error: provide valid number.")
	//  }
	// } else if len(input) == 4 && input[0] == "-w" && input[2] == "-p" {
	//  if len(strings.Fields(input[3])) <= 1 {
	//   fmt.Println("Length of prefix not valid")
	//   return
	//  }
	//  if num, err := strconv.Atoi(input[1]); err == nil && num < 10001 && num >= 0 {
	//   Result = append(Result, strings.Fields(input[3])...)
	//   engine(num, len(strings.Fields(input[3])))

	//  } else if err != nil  num > 10000  num < 0 {
	//   fmt.Println("Error: provide valid number.")
	//  }
	// } else if len(input) == 6 && input[0] == "-w" && input[2] == "-p" && input[4] == "-l" {

	//  num, err := strconv.Atoi(input[1])
	//  san, errors := strconv.Atoi(input[5])
	//  if len(strings.Fields(input[3])) <= 1 {
	//   fmt.Println("Length of prefix not valid")
	//   return
	//  }
	//  if errors != nil  err != nil  san > 5  num > 10000  num <= 0  san < 2  len(strings.Fields(input[3])) < san {
	//   fmt.Println("Error: provide valid numbers")
	//  } else if len(strings.Fields(input[3])) >= san {
	//   slice := strings.Fields(input[3])
	//   Result = append(Result, slice[:san]...)
	//   engine(num, san)
	//  }

	// } else {
	//  help()
	// }
}
