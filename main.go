package main

import (
	"github.com/d-sense/protocolBufferBasicGo/utility"
	"fmt"
	)

func main() {
	personMessage := utility.GetPersonExercise()
	utility.PersonAddressFileStorage(personMessage)

	jsonFormat := utility.FromProtoToJson(personMessage)
	fmt.Println("JSON FORMAT: ", jsonFormat)
}


