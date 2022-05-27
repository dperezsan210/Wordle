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

	letters := input

	word := generateWord(letters)

	for win == false {

		fmt.Print("-> ")
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		valid, err := validate(input, toInt(letters))

		if valid {
			myWord := separateWordInLetter(input)

			myWord, win = checkWord(myWord, word)
			printWordWithColor(myWord)

			fmt.Println(colorReset)

			if win {
				fmt.Println("Felicidades! Has acertado")
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
