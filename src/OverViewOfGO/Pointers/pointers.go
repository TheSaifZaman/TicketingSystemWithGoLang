package pointers

import "log"

func Pointers() {
	var myString string
	myString = "Hello Green"
	log.Println("My string is saying: ", myString)

	changeUsingPointer(&myString)
	log.Println("My string is saying: ", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Hello Red"
	*s = newValue
}