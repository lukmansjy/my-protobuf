package basic

import (
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"
)

func BasicUser() {
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
		LastLogin:            timestamppb.Now(),
		BirthDate:            &date.Date{Year: 2000, Month: 5, Day: 27},
		LastLocation: &latlng.LatLng{
			Latitude:  -6.29847717,
			Longitude: 106.8290577,
		},
	}

	//log.Println(&user)
	jsonBytes, _ := protojson.Marshal(&user)
	log.Println(string(jsonBytes))
}

func ProtoUserToJson() {
	user := basic.User{
		Id:       2,
		Username: "Lukman Sanjaya",
		IsActive: true,
		Password: []byte("PassSecret789"),
		Emails:   []string{"lukam@mail.com", "kontak@mail.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	jsonBytes, _ := protojson.Marshal(&user)
	log.Println(string(jsonBytes))
}

func JsonUserToProto() {
	json := `
		 {
			"id":2,
			"username":"Lukman Sanjaya",
			"is_active":true,
			"password":"UGFzc1NlY3JldDc4OQ==",
			"emails":["lukam@mail.com","kontak@mail.com"],
			"gender":"GENDER_MALE"
		}
	`

	user := basic.User{}
	err := protojson.Unmarshal([]byte(json), &user)
	if err != nil {
		log.Println(err)
	}

	log.Println(&user)
}

func randomCommunicationChannel() anypb.Any {
	paperMail := basic.PaperMail{
		PaperMailAddress: "Some paper mail address",
	}

	socialMedia := basic.SocialMedia{
		Platform: "byteMe",
		Username: "krypton.man",
	}

	instantMessaging := basic.InstantMessaging{
		Product:  "whatsup",
		Username: "krypton.last",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		Platform: "facebook",
		Username: "lukman",
	}

	anyData := anypb.Any{}

	anypb.MarshalFrom(&anyData, &socialMedia, proto.MarshalOptions{})

	// know type (social media)

	sm := basic.SocialMedia{}
	if err := anyData.UnmarshalTo(&sm); err != nil {
		log.Println(err)
	}

	jsonBytes, err := protojson.Marshal(&sm)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonBytes))
}

// Unmarshal data yg tidak diketahui
func BasicUnmarshalAnyUnKnown() {
	random := randomCommunicationChannel()

	var unmarshal protoreflect.ProtoMessage

	unmarshal, err := random.UnmarshalNew()

	if err != nil {
		log.Println(err)
	}

	log.Println("Unmarshal as", unmarshal.ProtoReflect().Descriptor().Name())

	jsonBytes, err := protojson.Marshal(unmarshal)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonBytes))
}

// mengetahui type messege tertentu (mengubah paper mail)
func BasicUnmarshalAnyIs() {
	random := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if random.MessageIs(&pm) {
		if err := random.UnmarshalTo(&pm); err != nil {
			log.Println(err)
		}

		jsonBytes, err := protojson.Marshal(&pm)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaparMial, but :", random.TypeUrl)
	}

}

// fungsi contoh yg dapat menerima massege yg ditentukan
func BasicOneof() {
	//socoalMedia := basic.SocialMedia{
	//	Platform: "Instagram",
	//	Username: "lukman",
	//}
	//
	//ecomm := basic.User_SocialMedia{
	//	SocialMedia: &socoalMedia,
	//}

	instantMess := basic.InstantMessaging{
		Product:  "Whatsapp",
		Username: "023823293293",
	}

	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instantMess,
	}

	user := basic.User{
		Id:                   123,
		Username:             "lukman",
		ElectronicComChannel: &ecomm,
	}

	jsonBytes, err := protojson.Marshal(&user)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonBytes))
}
