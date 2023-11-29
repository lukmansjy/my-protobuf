package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicHello() {
	hello := basic.Hello{
		Name: "Lukman Sanjaya",
	}

	log.Println(&hello)
}
