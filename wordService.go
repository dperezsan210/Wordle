package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

func generateWord(letters int) (word, string) {

	randomWord, definition := generateWordApi(letters)
	shortDefinition := strings.Split(definition, ":")

	return separateWordInLetter(randomWord), shortDefinition[0]
}

func separateWordInLetter(myWord string) word {
	result := word{}

	toSeparate := strings.ToUpper(myWord)

	for _, l := range toSeparate {
		result = append(result, letter{
			value: string(l),
			pos:   Wrong,
		})
	}

	return result
}

func checkWord(userWord string, answer word) (word, bool) {

	separatedWord := separateWordInLetter(userWord)
	return separatedWord.checkLettersPosition(answer)
}

func (separatedWord word) checkLettersPosition(answer word) (word, bool) {
	var result word

	for i, myLetter := range separatedWord {
		for j, resultLetter := range answer {

			if myLetter.value == resultLetter.value && i == j {
				myLetter.pos = Correct
			}

			if myLetter.pos == Wrong {
				if myLetter.value == resultLetter.value && i != j {
					myLetter.pos = Exist
				} else {
					myLetter.pos = Wrong
				}
			}
		}
		result = append(result, myLetter)
	}

	return result, result.checkAllLettersCorrect()
}

func (separatedWord word) checkAllLettersCorrect() bool {
	for _, myLetter := range separatedWord {
		if myLetter.pos != Correct {
			return false
		}
	}
	return true
}

func printWordWithColor(myWord word) {

	for _, myLetter := range myWord {
		if myLetter.pos == Correct {
			fmt.Print(ColorGreen, myLetter.value, ColorReset)
		} else if myLetter.pos == Exist {
			fmt.Print(ColorYellow, myLetter.value, ColorReset)
		} else {
			fmt.Print(myLetter.value)
		}
	}

	fmt.Println()
}

func validate(myWord string, letters int) (bool, string) {
	if len(myWord) > letters {
		return false, "*Tu palabra tiene demasiadas letras*"
	} else if len(myWord) < letters {
		return false, "*Tu palabra tiene muy pocas letras*"
	} else {
		return true, ""
	}
}

func toInt(num string) int {
	intVar, _ := strconv.Atoi(num)
	return intVar
}

func toString(num int) string {
	return strconv.Itoa(num)
}

func printWelcome() {
	fmt.Println("**********************************")
	fmt.Println("* Bienvenido al Wordle en Golang *")
	fmt.Println("**********************************")
}

func winMessage(input string, definition string) {
	fmt.Println()
	fmt.Println("Felicidades! Has acertado.")
	fmt.Println(strings.ToUpper(input) + ": " + definition)
}

func requestNumber() int {
	fmt.Print("¿Cuantas letras quieres que tenga la palabra? ")
	input := requestWord()
	return toInt(input)
}

func requestWord() string {
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.Replace(input, "\n", "", -1)
}
