package basic

import (
	"google.golang.org/protobuf/proto"
	"log"
	"my-protobuf/protogen/basic"
	"os"
)

func WriteProtoToFile(msg proto.Message, fileName string) {
	marshal, err := proto.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	if err := os.WriteFile(fileName, marshal, 0644); err != nil {
		log.Println(err)
	}
}

func ReadProtoFromFile(fileName string, dest proto.Message) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}

	if err := proto.Unmarshal(file, dest); err != nil {
		log.Println(err)
	}
}

func dummyUser() basic.User {
	address := basic.Address{
		Street:  "Jalan Suka Maju",
		City:    "Wonogiri",
		Country: "Indonesia",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  19.0283,
			Longitude: -7.9233,
		},
	}

	comm := randomCommunicationChannel()
	skillRating := map[string]uint32{"golang": 8, "python": 8, "nodejs": 9}

	user := basic.User{
		Id:                   1,
		Username:             "Lukman",
		IsActive:             true,
		Password:             []byte("PassSecret1234"),
		Emails:               []string{"lukam@mail.com", "kontak@mail.com"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &address,
		CommunicationChannel: &comm,
		SkillRating:          skillRating,
	}

	return user
}

func WriteToFileSample() {
	user := dummyUser()

	WriteProtoToFile(&user, "lukman_file.bin")
}

func ReadFromFileSample() {
	user := basic.User{}

	ReadProtoFromFile("lukman_file.bin", &user)

	log.Println(&user)
}
