package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Arrivals []struct {
	UpdateTime string   `xml:"UpdateTime"`
	Flight     []Flight `xml:"Flight"`
}

type Flight struct {
	AirLineID     string `xml:"AirLineID"`
	AirLineName   string `xml:"AirLineName"`
	FlightID      string `xml:"FlightID"`
	Route1        string `xml:"Route1"`
	Route2        string `xml:"Route2"`
	Route3        string `xml:"Route3"`
	Route4        string `xml:"Route4"`
	Route5        string `xml:"Route5"`
	Route6        string `xml:"Route6"`
	Scheduled     string `xml:"Scheduled"`
	Estimated     string `xml:"Estimated"`
	Status        string `xml:"Status"`
	Gate          string `xml:"Gate"`
	Terminal      string `xml:"Terminal"`
	CheckinArea   string `xml:"CheckinArea"`
	Date          string `xml:"Date"`
	PrimaryFlight string `xml:"PrimaryFlight"`
	CarrierType   string `xml:"CarrierType"`
}

func ProcessArrivals() string {
	resp, err := http.Get("http://fis.com.mv/xml/arrive.xml")
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(jsonData)
}

func main() {
	fmt.Printf(ProcessArrivals())
}
