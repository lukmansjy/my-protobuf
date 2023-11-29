package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"my-protobuf/protogen/basic"
	"os"
)

func WriteProtoToJSON(msg proto.Message, fileName string) {
	marshal, err := protojson.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	if err := os.WriteFile(fileName, marshal, 0644); err != nil {
		log.Println(err)
	}
}

func ReadProtoFromJSON(fileName string, dest proto.Message) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}

	if err := protojson.Unmarshal(file, dest); err != nil {
		log.Println(err)
	}
}

func WriteToFileJSONSample() {
	user := dummyUser()

	WriteProtoToJSON(&user, "lukman_file.json")
}

func ReadFromFileJSONSample() {
	user := basic.User{}

	ReadProtoFromJSON("lukman_file.json", &user)

	log.Println(&user)
}
