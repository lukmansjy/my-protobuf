package car

import (
	"github.com/google/uuid"
	"log"
	"my-protobuf/protogen/car"
)

func ValidateCar() {
	car := car.Car{
		Id:              uuid.New().String(),
		Brand:           "Bmw",
		Model:           "202d-232d-3r92",
		ManufactureYear: 2023,
		Price:           20000000,
	}

	err := car.ValidateAll()
	if err != nil {
		log.Println(err)
	}

	log.Println(&car)

}
