package receiversStructWithFunctions

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func ReceiversStructWithFunctions() {
	var myVar myStruct

	myVar.FirstName = "Md Saif"

	myVar2 := myStruct{
		FirstName: "Md. Saif",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}