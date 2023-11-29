package jobsearch

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
	"my-protobuf/protogen/dummy"
	"my-protobuf/protogen/jobsearch"
)

func JobSearchSoftware() {
	jobSoftware := jobsearch.JobSoftware{
		Id: 1,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "The amazing proto",
			Platforms: []string{"Mac", "Linux", "Windows"},
		},
	}

	jsonByte, _ := protojson.Marshal(&jobSoftware)
	log.Println(string(jsonByte))
}

func JobSearchCandidate() {
	jobCandidate := jobsearch.JobCandidate{
		Id: 1,
		Application: &dummy.Application{
			Id:       1,
			FullName: "Lukman",
			Phone:    "02820482842",
			Email:    "luk@mail.com",
		},
	}

	jsonByte, _ := protojson.Marshal(&jobCandidate)
	log.Println(string(jsonByte))

}
