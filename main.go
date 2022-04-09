package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const arrivalsUrl string = "http://fis.com.mv/xml/arrive.xml"
const departuresUrl string = "http://fis.com.mv/xml/depart.xml"

type Arrivals []struct {
	UpdateTime string    `json:"update_time" xml:"UpdateTime"`
	Flights    []Flights `json:"flights" xml:"Flight"`
}

type Departures []struct {
	UpdateTime string    `json:"update_time" xml:"UpdateTime"`
	Flights    []Flights `json:"flights" xml:"Flight"`
}

type Flights struct {
	AirLineID     string `json:"update_time" xml:"AirLineID"`
	AirLineName   string `json:"airline_name" xml:"AirLineName"`
	FlightID      string `json:"flight_id" xml:"FlightID"`
	Route1        string `json:"route_1" xml:"Route1"`
	Route2        string `json:"route_2" xml:"Route2"`
	Route3        string `json:"route_3" xml:"Route3"`
	Route4        string `json:"route_4" xml:"Route4"`
	Route5        string `json:"route_5" xml:"Route5"`
	Route6        string `json:"route_6" xml:"Route6"`
	Scheduled     string `json:"scheduled" xml:"Scheduled"`
	Estimated     string `json:"estimated" xml:"Estimated"`
	Status        string `json:"status" xml:"Status"`
	Gate          string `json:"gate" xml:"Gate"`
	Terminal      string `json:"terminal" xml:"Terminal"`
	CheckinArea   string `json:"checkin_area" xml:"CheckinArea"`
	Date          string `json:"date" xml:"Date"`
	PrimaryFlight string `json:"primary_flight" xml:"PrimaryFlight"`
	CarrierType   string `json:"carrier_type" xml:"CarrierType"`
}

func GetArrivals() ([]byte, error) {
	resp, err := http.Get(arrivalsUrl)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	data := &Arrivals{}
	err = xml.Unmarshal(body, data)
	if err != nil {
		panic(err)
	}

	return json.Marshal(data)
}

func GetDepartures() ([]byte, error) {
	/**
	You are going to see quite bit of redundant code below from GetArrivals func. TODO: Cleanup
	*/

	resp, err := http.Get(departuresUrl)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	data := &Departures{}
	err = xml.Unmarshal(body, data)
	if err != nil {
		panic(err)
	}

	return json.Marshal(data)
}

func main() {
	arrivals, err := GetArrivals()
	if err != nil {
		panic(err)
	}

	departures, err := GetDepartures()
	if err != nil {
		panic(err)
	}

	// Just for the sake of printing...
	fmt.Printf("ARRIVALS\n")
	fmt.Printf(string(arrivals))
	fmt.Printf("\n\n\n")
	fmt.Printf("DEPARTURES\n")
	fmt.Printf(string(departures) + "\n")
}
