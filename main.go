package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/nathan-osman/go-sunrise"
	"gopkg.in/resty.v1"
)

func main() {
	//TODO User inputs address
	fmt.Println("Find out when the Sun will rise and set at your location.")
	fmt.Println("Input your address")
	reader := bufio.NewReader(os.Stdin)
	address, _ := reader.ReadString('\n')
	//TODO Get request to google
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"address": address,
		}).
		SetHeader("Accept", "application/json").
		Get("http://maps.google.com/maps/api/geocode/json")
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	//fmt.Printf("\nResponse Body: %v", resp)
	//fmt.Println(reflect.TypeOf(resp.String()))

	//TODO Parse JSON for lat, lng
	pat := regexp.MustCompile(`"location"\s*:\s*{\s*"lat"\s*:\s-?\d+.\d+,\s*"lng"\s*:\s*-?\d+.\d+\s*}`)
	s := pat.FindString(resp.String())
	//fmt.Println(s)

	pat = regexp.MustCompile(`-?\d+.\d+`)
	coordinates := pat.FindAllString(s, -1)
	//fmt.Println(coordinates)
	latitude, _ := strconv.ParseFloat(coordinates[0], 64)
	longitude, _ := strconv.ParseFloat(coordinates[1], 64)
	//fmt.Println(latitude)
	//fmt.Println(longitude)
	//TODO put lat, lng into sunrise sunset
	rise, set := sunrise.SunriseSunset(longitude, latitude, 2018, time.June, 19)
	//TODO display the sunrise and sunset
	color.Blue("For the address of :" + address)
	color.Blue("Sunrise is at " + rise.String())
	color.Blue("Sunset is at " + set.String())
	//getGoogleMapsCoord()
	// rise, set := sunrise.SunriseSunset(
	// 	43.607561, -116.205834, // Downtown Boise
	// 	2018, time.June, 19, // 2000-01-01
	// )
	// fmt.Println(rise)
	// fmt.Println(set)
	// calSunRiseSunSet(43.607561, -116.205834, 2018, time.June, 19)
	//data := readTxtFile("Gettysburg-Address.txt")
	//findUpperCaseWords(data)
}

func getCoordFromAddress(address string) {

}

func getGoogleMapsCoord() {
	//resp, err := resty.R().Get("http://maps.google.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA")

	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"address": "12851 West Silverbrook Ct Boise ID",
		}).
		SetHeader("Accept", "application/json").
		Get("http://maps.google.com/maps/api/geocode/json")

	// explore response object
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())

}

func calSunRiseSunSet(long float64, lat float64, year int, month time.Month, day int) (time.Time, time.Time) {
	rise, set := sunrise.SunriseSunset(
		long, lat,
		year, month, day,
	)
	return rise, set
}
