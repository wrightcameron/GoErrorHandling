package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestCalculateSunRiseSet_Happy(t *testing.T) {
	rise, set := calculateSunRiseSet(49, -119)
	fmt.Println(rise.String())
	fmt.Println(set.String())
}

func TestGetCoordinatesFromResponse_Happy(t *testing.T) {
	content, err := ioutil.ReadFile("./jsonfiles/GoogleMapsGetEx.json")
	if err != nil {
		log.Fatal(err)
	}
	latitude, longitude, err := getCoordinatesFromResponse(string(content))
	if err != nil {
		t.Errorf("Error thrown with correct input from JSON String")
	}
	fmt.Println(latitude)
	fmt.Println(longitude)
}

func TestGetCoordinatesFromResponse_CoordinatesMissing(t *testing.T) {
	content, err := ioutil.ReadFile("./jsonfiles/GoogleMapsCoordMiss.json")
	if err != nil {
		log.Fatal(err)
	}
	latitude, longitude, err := getCoordinatesFromResponse(string(content))
	if err == nil {
		t.Errorf("Error should have been thrown.")
	}
	fmt.Println(latitude)
	fmt.Println(longitude)
}
