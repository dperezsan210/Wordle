package main

import (
	"fmt"
)

func main() {

	var myWord word
	win := false

	//Print the welcome message.
	printWelcome()

	//Ask the number of letters you want to play with.
	numberOfLetters := requestNumber()

	//Generate the word with its definition.
	word, definition := generateWord(numberOfLetters)

	for win == false {

		//Ask the user for a word.
		userWord := requestWord()

		//Validates that the entered word contains the same characters as the word to guess.
		valid, err := validate(userWord, numberOfLetters)

		if valid {

			//Check if it is the correct word from the position of its letters.
			myWord, win = checkWord(userWord, word)

			//Prints the existing letters in yellow and the letters that are in the correct place in green.
			printWordWithColor(myWord)

			if win {
				//Print wins message
				winMessage(userWord, definition)
			}

		} else {
			fmt.Println(err)
		}
	}
}
