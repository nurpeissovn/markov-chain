# Markov Chain Text Generator
This is a Markov Chain-based text generator written in Go. It reads input text (either from a file or standard input), processes it to build a word chain, and then generates new text based on the statistical properties of word sequences from the input. 
## Features
* ***Generate text*** based on a Markov Chain model.
* ***Control the number of generated words*** with the -w flag.
* ***Specify a starting prefix*** using the -p flag.
* ***Customize the prefix*** length with the -l flag.

## Command-Line Arguments (Keys)

* `-w <N>`
    
    Specify the number of words to generate. The default value is 100 words.

* `-p <S>`

    Specify the starting prefix for generating the sequence. This prefix will be used as the beginning of the generated text.

* `-l <N>`

    Specify the length of the prefix to consider when chaining words. The default is 2 and maximum goes until 5.

* `--help`

    Display the help message with all available command-line options.

## Running the Program
    You can run the program by either providing a file with text input or by piping text directly from the command line.