package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/nathan-osman/go-sunrise"
	"gopkg.in/resty.v1"
)

//getUserAddress asks the user what there address is from the console.
//It returns that address as a string.
func getUserAddress() (string, error) {
	fmt.Println("Find out when the Sun will rise and set at your location.")
	fmt.Println("Input your address")
	reader := bufio.NewReader(os.Stdin)
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Improper string was writen to console.")
	}
	return address, err
}

//sendGetRequestToGoogleMaps sends a REST GET request to Google map's API,
//with a given address.  A resty response object is given back.
func sendGetRequestToGoogleMaps(address string) (float64, float64, error) {
	resp, getErr := resty.R().
		SetQueryParams(map[string]string{
			"address": address,
		}).
		SetHeader("Accept", "application/json").
		Get("http://maps.google.com/maps/api/geocode/json")
	if getErr != nil {
		fmt.Printf("\nREST API Error: %v", getErr)
	}
	latitude, longitude, coordErr := getCoordinatesFromResponse(resp.String())
	var err error
	if getErr != nil || coordErr != nil {
		err = errors.New("coordinates converted improperly")
	}
	return latitude, longitude, err
}

//getCoordinatesFromResponse deals with extracting the longitude and
//latitude from the JSON String.
func getCoordinatesFromResponse(JSONString string) (float64, float64, error) {
	//I was going to see if golang had an object that could convert JSON
	//strings to JSON but after searching I decided to use regex.
	pat := regexp.MustCompile(`"location"\s*:\s*{\s*"lat"\s*:\s-?\d+.\d+,\s*"lng"\s*:\s*-?\d+.\d+\s*}`)
	s := pat.FindString(JSONString)
	//fmt.Println(s)
	var err error
	if len(s) == 0 {
		err = errors.New("JSON did not contain location")
		return -1, -1, err
	}
	pat = regexp.MustCompile(`-?\d+.\d+`)
	coordinates := pat.FindAllString(s, -1)
	latitude, latErr := strconv.ParseFloat(coordinates[0], 64)
	longitude, longErr := strconv.ParseFloat(coordinates[1], 64)
	if latErr != nil || longErr != nil {
		err = errors.New("coordinates converted improperly")
	}
	return latitude, longitude, err
}

//Calculate the sunrise Set uses the sunrise library to calculate
//the sunrise from given information.
func calculateSunRiseSet(longitude, latitude float64) (time.Time, time.Time) {
	rise, set := sunrise.SunriseSunset(longitude, latitude, time.Now().Local().Year(), time.Now().Local().Month(), time.Now().Local().Day())
	return rise, set
}
func main() {
	address, err := getUserAddress()
	//Exit project here because exiting in the function would stop a unit test
	if err != nil {
		os.Exit(-1)
	}
	latitude, longitude, coordError := sendGetRequestToGoogleMaps(address)
	if coordError != nil {
		os.Exit(-1)
	}
	sunrise, sunset := calculateSunRiseSet(longitude, latitude)

	color.Blue("For the address of :" + address)
	color.Blue("Sunrise is at " + sunrise.String())
	color.Blue("Sunset is at " + sunset.String())
}
