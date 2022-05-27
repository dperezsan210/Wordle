# Wordle

Hi, this is a project based on the word game wordle: https://wordle.danielfrg.com/

This project is made 100% in Golang and works by returning a random word either from an external API or from a .txt file. In order to use a .txt file to return the random words, we will need to uncomment the lines of the "word.go" class and add the words to the "word_es" document separated by a space. We will also have to comment out the line that calls the generateWordApi() method.

This is the API that generates random words: https://rapidapi.com/AlexScigalszky/api/spanish-random-words/
