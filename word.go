package main

type word []letter

type letter struct {
	value string
	pos   position
}

type position int

const (
	Wrong position = iota
	Exist
	Correct
)
