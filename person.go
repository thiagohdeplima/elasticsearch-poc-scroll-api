package main

import (
	"log"
	"math/rand"
	"reflect"

	"github.com/bxcodec/faker/v3"
)

type Person struct {
	Name string `faker:"name" json:"name"`

	Location Location `json:"location"`

	DateTime string `faker:"timestamp" json:"date_time"`
	IDNumber string `faker:"len=15" json:"id_number"`
}

type Location struct {
	Latitude  float32 `faker:"customLat" json:"lat"`
	Longitude float32 `faker:"customLon" json:"lon"`
}

func GeneratePerson() (Person, error) {
	addCustomGenerators()

	person := Person{}

	if err := faker.FakeData(&person); err != nil {
		log.Printf("can't generate person `%s`", err)
		return person, err
	}

	return person, nil
}

func addCustomGenerators() {
	var minLat float32 = 21.1
	var maxLat float32 = 23.5

	var minLon float32 = 46.7
	var maxLon float32 = 49.2

	_ = faker.AddProvider("customLat", func(v reflect.Value) (interface{}, error) {
		return -(minLat + rand.Float32()*(maxLat-minLat)), nil
	})

	_ = faker.AddProvider("customLon", func(v reflect.Value) (interface{}, error) {
		return -(minLon + rand.Float32()*(maxLon-minLon)), nil
	})
}
