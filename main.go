package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	colorReset := "\033[0m"
	win := false

	printWelcome()

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	letters := toInt(input)

	word, definition := generateWord(letters)

	for win == false {

		fmt.Print("-> ")
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		valid, err := validate(input, letters)

		if valid {
			myWord := separateWordInLetter(input)

			myWord, win = checkWord(myWord, word)
			printWordWithColor(myWord)

			fmt.Println(colorReset)

			if win {
				winMessage(input, definition)
			}

		} else {
			fmt.Println(err)
		}
	}
}

func printWelcome() {
	fmt.Println("Bienvenido al Wordle en Golang")
	fmt.Println("------------------------------")
	fmt.Print("¿Cuantas letras quieres que tenga la palabra? -> ")
}

func winMessage(input string, definition string) {
	fmt.Println()
	fmt.Println("Felicidades! Has acertado.")
	fmt.Println(strings.ToUpper(input) + ": " + definition)
}
