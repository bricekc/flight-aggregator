package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Flight1 struct {
	Number string `json:"flightNumber"`
	From   string `json:"departureAirport"`
	To     string `json:"arrivalAirport"`
	Depart string `json:"departureTime"`
	Arrive string `json:"arrivalTime"`
}

type Total1 struct {
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type Travelling1 struct {
	BookingId     string `json:"bookingId"`
	Status        string `json:"status"`
	PassengerName string `json:"passengerName"`
	Flight        []Flight1
	Total         Total1
}

type Travellings1 struct {
	Travellings []Travelling1
}

func (r *Travellings1) FlightS1() ([]Travelling1, error) {
	res, err := http.Get("http://j-server1:4001/flights")
	if err != nil {
		fmt.Println("error!!!", err)
		return nil, err
	}
	fmt.Println("body -> ", res.Body)
	var data []Travelling1
	err2 := json.NewDecoder(res.Body).Decode(&data)
	if err2 != nil {
		fmt.Println("error", err2)
		return nil, err2
	}
	fmt.Println("data -> ", data)
	return data, nil
}
