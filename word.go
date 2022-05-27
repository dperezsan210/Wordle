package main

import (
	"fmt"
	"strconv"
	"strings"
)

type position int

type letter struct {
	value string
	pos   position
}

const (
	Wrong position = iota
	Exist
	Correct
)

type word []letter

func generateWord(letters int) word {

	/*words := getAllWordsFromFile()
	randomWord := takeRandomWord(words)*/

	randomWord := generateWordApi(letters)

	return separateWordInLetter(randomWord)
}

/*func getAllWordsFromFile() []string {
	bs, err := ioutil.ReadFile("word_es")

	if err != nil {

		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), " ")
}*/

/*func takeRandomWord(words []string) string {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	pos := r.Intn(len(words) - 1)

	return words[pos]
}*/

func separateWordInLetter(myWord string) word {

	toSeparate := strings.ToUpper(myWord)

	result := word{}

	for _, l := range toSeparate {
		result = append(result, letter{
			value: string(l),
			pos:   Wrong,
		})
	}

	return result
}

func checkWord(myWord word, answer word) (word, bool) {

	result := word{}

	for i, myLetter := range myWord {
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

	return result, checkAllLettersCorrect(result)
}

func checkAllLettersCorrect(myWord word) bool {
	for _, myLetter := range myWord {
		if myLetter.pos != Correct {
			return false
		}
	}
	return true
}

func printWordWithColor(myWord word) {
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorReset := "\033[0m"

	for _, myLetter := range myWord {
		if myLetter.pos == Correct {
			fmt.Print(colorGreen, myLetter.value, colorReset)
		} else if myLetter.pos == Exist {
			fmt.Print(colorYellow, myLetter.value, colorReset)
		} else {
			fmt.Print(myLetter.value)
		}
	}
}

func toInt(num string) int {
	intVar, _ := strconv.Atoi(num)
	return intVar
}

func toString(num int) string {
	return strconv.Itoa(num)
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
