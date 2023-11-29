package main

import (
	"fmt"
	car "my-protobuf/Car"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Println(time.Now().Format("15:04:05") + " " + string(bytes))
}
func main() {
	//log.SetFlags(0)
	//log.SetOutput(new(logWriter))

	//basic.BasicHello()
	//basic.BasicUser()
	//basic.ProtoUserToJson
	//basic.JsonUserToProto()
	//basic.BasicUserGroup()

	//jobsearch.JobSearchCandidate()
	//jobsearch.JobSearchSoftware()

	//basic.BasicUnmarshalAnyKnown()
	//basic.BasicUnmarshalAnyUnKnown()
	//basic.BasicUnmarshalAnyIs()
	//basic.BasicOneof()

	//basic.WriteToFileSample()
	//basic.ReadFromFileSample()
	//basic.WriteToFileJSONSample()
	//basic.ReadFromFileJSONSample()
	//basic.BasicWriteUserContentV1()
	//basic.BasicReadUserContentV1()
	//basic.BasicWriteUserContentV2()
	//basic.BasicReadUserContentV2()
	//basic.BasicWriteUserContentV3()
	//basic.BasicReadUserContentV3()
	//basic.BasicWriteUserContentV4()
	//basic.BasicReadUserContentV4()
	//basic.BasicReadUserPayment()
	car.ValidateCar()
}
