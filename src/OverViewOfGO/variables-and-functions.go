package main

import "log"

func VariablesAndFunctions() {
	var whatToSay string
	whatToSay, _ = saySomeThing("Hello, How are you?")
	log.Println(whatToSay)

	whatEsleToSay, world := saySomeThing(("GoodBye, Amigos."))
	log.Println(whatEsleToSay," ",world)

	log.Println(saySomeThing("See you tomorrow."))
}

func saySomeThing(s string) (string, string) {
	return s, "World"
}