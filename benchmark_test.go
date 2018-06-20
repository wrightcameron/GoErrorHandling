package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkCalculateSunRiseSet(b *testing.B) {
	// run the calculateSunRiseSet function b.N times
	for n := 0; n < b.N; n++ {
		calculateSunRiseSet(45, -119)
	}
}

func BenchmarkGetCoordinatesFromResponse(b *testing.B) {
	JSONString, err := ioutil.ReadFile("./jsonfiles/GoogleMapsGetEx.json")
	if err != nil {
		log.Fatal(err)
	}
	// run the calculateSunRiseSet function b.N times
	for n := 0; n < b.N; n++ {
		getCoordinatesFromResponse(string(JSONString))
	}
}
