package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

// contoh serialisasi yg tidak sesuai
func BasicReadUserPayment() {
	log.Println("Read User Payment")
	userPayment := basic.UserPayment{}

	ReadProtoFromFile("user_content_v1.bin", &userPayment)
	log.Println(&userPayment)

	jsonBytes, _ := protojson.Marshal(&userPayment)
	log.Println(string(jsonBytes))
}
