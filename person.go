package main

import (
	"log"

	"github.com/bxcodec/faker/v3"
)

type Person struct {
	Name string `faker:"name" json:"name"`

	Latitude  float32 `faker:"lat" json:"lat"`
	Longitude float32 `faker:"long" json:"lng"`

	DateTime string `faker:"time" json:"date_time"`
	IDNumber string `faker:"len=15" json:"id_number"`
}

func GeneratePerson() (Person, error) {
	person := Person{}

	if err := faker.FakeData(&person); err != nil {
		log.Printf("can't generate person `%s`", err)
		return person, err
	}

	return person, nil
}
