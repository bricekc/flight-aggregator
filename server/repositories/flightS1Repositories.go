package repositories

import (
	"encoding/json"
	"net/http"
	"os"
)

type Travelling1 struct {
	BookingId        string
	Status           string
	PassengerName    string
	FlightNumber     string
	DepartureAirport string
	ArrivalAirport   string
	DepartureTime    string
	ArrivalTime      string
	Price            float64
	Currency         string
}

func (r *Travelling1) GetTravel() ([]Travelling, error) {
	res, err := http.Get(os.Getenv("JSERVER1_URL"))
	if err != nil {
		return []Travelling{}, err
	}
	var data []Travelling1
	err2 := json.NewDecoder(res.Body).Decode(&data)
	if err2 != nil {
		return []Travelling{}, err2
	}

	return toTravellings1(data), nil
}

func toTravellings1(travellings1 []Travelling1) []Travelling {
	var travellings []Travelling
	for _, t := range travellings1 {
		travellings = append(travellings, Travelling{
			BookingId:     t.BookingId,
			Status:        t.Status,
			PassengerName: t.PassengerName,
			Flights: []Flight{
				{
					Number: t.FlightNumber,
					From:   t.DepartureAirport,
					To:     t.ArrivalAirport,
					Depart: t.DepartureTime,
					Arrive: t.ArrivalTime,
				},
			},
			Total: Total{
				Price:    t.Price,
				Currency: t.Currency,
			},
		})
	}
	return travellings
}
