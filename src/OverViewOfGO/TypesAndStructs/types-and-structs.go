package typesAndStructs

import (
	"log"
	"time"
)

var s = "Seven"

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func TypesAndStructs()  {
	var s2 = "Six"

	log.Println("s is ",s)
	log.Println("s2 is ",s2)
	saySomething("XXX")

	user := User {
		FirstName: "Md Saif",
		LastName: "Zaman",
		PhoneNumber: "+8801770000000",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber)
}

func saySomething(s string) (string, string) {
	log.Println(s)
	return s, "World"
}

