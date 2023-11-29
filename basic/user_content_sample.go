package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicWriteUserContentV1() {
	userContent := basic.UserContent{
		Id:   101,
		Slug: "/this-is-v1",
		//Title:       "This Is V1",
		//HtmlContent: "<p>Hello This Is V1</p>",
		//AuthorId:    1,
	}

	WriteProtoToFile(&userContent, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Read V1")
	userContent := basic.UserContent{}
	ReadProtoFromFile("user_content_v1.bin", &userContent)

	log.Println(&userContent)

	marshal, _ := protojson.Marshal(&userContent)
	log.Println(string(marshal))
}

func BasicWriteUserContentV2() {
	userContent := basic.UserContent{
		Id:   102,
		Slug: "/this-is-v2",
		//Title:       "This Is V2",
		//HtmlContent: "<p>Hello This Is V2</p>",
		//AuthorId:    1,
		//Category: "NEWS",
	}

	WriteProtoToFile(&userContent, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Read V2")
	userContent := basic.UserContent{}
	ReadProtoFromFile("user_content_v2.bin", &userContent)

	log.Println(&userContent)

	marshal, _ := protojson.Marshal(&userContent)
	log.Println(string(marshal))
}

func BasicWriteUserContentV3() {
	userContent := basic.UserContent{
		Id:   103,
		Slug: "/this-is-v3",
		//Title:       "This Is V3",
		//HtmlContent: "<p>Hello This Is V3</p>",
		//AuthorId:    1,
		//Category: "NEWS",
		//SubCategory: "TECH",
	}

	WriteProtoToFile(&userContent, "user_content_v3.bin")
}

func BasicReadUserContentV3() {
	log.Println("Read V3")
	userContent := basic.UserContent{}
	ReadProtoFromFile("user_content_v3.bin", &userContent)

	log.Println(&userContent)

	marshal, _ := protojson.Marshal(&userContent)
	log.Println(string(marshal))
}

func BasicWriteUserContentV4() {
	userContent := basic.UserContent{
		Id:   104,
		Slug: "/this-is-v4",
		//Title:       "This Is V4",
		//HtmlContent: "<p>Hello This Is V4</p>",
		//AuthorId:    1,
		//Category: "NEWS",
		//SubCategory: "TECH",
		//Rating: 5,
	}

	WriteProtoToFile(&userContent, "user_content_v4.bin")
}

func BasicReadUserContentV4() {
	log.Println("Read V4")
	userContent := basic.UserContent{}
	ReadProtoFromFile("user_content_v4.bin", &userContent)

	log.Println(&userContent)

	marshal, _ := protojson.Marshal(&userContent)
	log.Println(string(marshal))
}
