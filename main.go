package main

import (
	"github.com/d-sense/protocolBufferBasicGo/utility"
	"fmt"
	"log"
)

func main() {
	personMessage := utility.GetPersonExercise()
	utility.PersonAddressFileStorage(personMessage)

	jsonFormat, err := utility.FromProtoToJson(personMessage)
	if  err != nil {
		log.Fatal("Could not convert to JSON: ", err)
	}

	fmt.Println("JSON FORMAT: ", jsonFormat)
}


