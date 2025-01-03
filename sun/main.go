package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// https://api.openweathermap.org/data/3.0/onecall?lat=$latitude&lon=$longitude&appid=$APIKey'
// http://api.weatherapi.com/v1/forecast.json?key=<key>&q=London&days=1&aqi=no&alerts=no



type Location struct {
	//Location struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
	//} `json:"location"`
}

func main() {
	#res, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=London&limit=5&appid=<id>")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Res status code", res.StatusCode)
	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(body))
	var locations []Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		panic(err)
	}
	fmt.Println(locations)
}
