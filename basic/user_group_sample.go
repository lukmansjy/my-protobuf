package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUserGroup() {
	batman := basic.User{
		Id:       3,
		Username: "Batman",
		IsActive: true,
		Password: []byte("passwd"),
		Emails:   []string{"bat@man.com", "batman@web.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	nightwing := basic.User{
		Id:       4,
		Username: "Nightwing",
		IsActive: true,
		Password: []byte("passwdnightwing"),
		Emails:   []string{"nightwing@mail.com", "nightwing@web.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	robin := basic.User{
		Id:       5,
		Username: "Robin",
		IsActive: true,
		Password: []byte("passwdrobin"),
		Emails:   []string{"robin@mail.com", "robin@web.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	batFamily := basic.UserGroup{
		GroupId:     1,
		GroupName:   "Bat Family",
		Role:        []string{"admin", "user"},
		User:        []*basic.User{&batman, &nightwing, &robin},
		Description: "The classic bat family",
	}

	jsonBytes, err := protojson.Marshal(&batFamily)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonBytes))
}
