package utility

import (
	"fmt"
	personpb "github.com/d-sense/ProtocBuf/protobuf-example-go/src/exercise"
	simplepb "github.com/d-sense/ProtocBuf/protobuf-example-go/src/simple"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"io/ioutil"
	"log"
	"github.com/gogo/protobuf/jsonpb"

)

func GetPersonExercise() *personpb.AddressBook {
	person := &personpb.Person{
		Name:  "Adeshina hammed Hassan",
		Id:    1,
		Email: "delameh@icloud.com",
		Phones: []*personpb.Person_PhoneNumber{
			&personpb.Person_PhoneNumber{
				Number: "0807321221",
				Type:   personpb.Person_HOME,
			},
			&personpb.Person_PhoneNumber{
				Number: "08083769214",
				Type:   personpb.Person_WORK,
			},
			&personpb.Person_PhoneNumber{
				Number: "080946221238",
				Type:   personpb.Person_MOBILE,
			},
		},
		LastUpdated: &timestamp.Timestamp{
			Seconds: 100,
			Nanos:   50,
		},
	}

	address := personpb.AddressBook{
		People: []*personpb.Person{
			person,
		},
	}
	return &address
}

func PersonAddressFileStorage(sm proto.Message) {
	err := writeAddressToFile("address.bin", sm)
	if err != nil {
		log.Fatal("Could not serialize file content to bytes", err)
	}

	sm2 := &simplepb.SimpleMessage{}

	err = readAddressFromFile("address.bin", sm2)
	if err != nil {
		log.Fatal("Could not un-serialize from bytes to Proto type", err)
	}

	fmt.Println("Content of stored data: ", sm2)
}

func readAddressFromFile(fileName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("can't read from file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatal("Couldn't put the bytes int the protocol buffers struct", err)
		return err2
	}

	return nil
}

func writeAddressToFile(fileName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		return err
	}
	return nil
}

func PersonAddressJSON(pb proto.Message) {
	smAsString := FromProtoToJson(pb)
	fmt.Println(smAsString)

	sm2 := &personpb.AddressBook{}
	FromJsonToProto(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func FromJsonToProto(in string, pb proto.Message){
	err := jsonpb.UnmarshalString(in, pb)
	if  err != nil {
		log.Fatal("Could not unmarshal the json", err)
	}
}

func FromProtoToJson(pb proto.Message) string {
	marshal := jsonpb.Marshaler{}
	out, err := marshal.MarshalToString(pb)
	if  err != nil {
		log.Fatal("Could not convert to JSON", err)
		return ""
	}

	return out
}