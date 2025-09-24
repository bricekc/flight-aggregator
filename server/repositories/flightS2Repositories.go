package repositories

import (
	"encoding/json"
	"net/http"
	"os"
)

type Travelling2 struct {
	Reference string
	Status    string
	Traveler  Traveler
	Segments  []Segment
	Total     Total2
}

type Traveler struct {
	FirstName string
	LastName  string
}

type Segment struct {
	Flight Flight2
}

type Flight2 struct {
	Number string
	from   string
	to     string
	Depart string
	Arrive string
}

type Total2 struct {
	Amount   float64
	Currency string
}

func (r *Travelling2) GetTravel() ([]Travelling, error) {
	res, err := http.Get(os.Getenv("JSERVER2_URL"))
	if err != nil {
		return []Travelling{}, err
	}
	var data []Travelling2
	err2 := json.NewDecoder(res.Body).Decode(&data)
	if err2 != nil {
		return []Travelling{}, err2
	}

	return toTravellings2(data), nil
}

func toTravellings2(travellings2 []Travelling2) []Travelling {
	var travellings []Travelling
	for _, t := range travellings2 {
		var flights []Flight
		for _, s := range t.Segments {
			f := s.Flight
			flights = append(flights, Flight{
				Number: f.Number,
				From:   f.from,
				To:     f.to,
				Depart: f.Depart,
				Arrive: f.Arrive,
			})
		}

		travellings = append(travellings, Travelling{
			BookingId:     t.Reference,
			Status:        t.Status,
			PassengerName: t.Traveler.FirstName + " " + t.Traveler.LastName,
			Flights:       flights,
			Total: Total{
				Price:    t.Total.Amount,
				Currency: t.Total.Currency,
			},
		})
	}
	return travellings
}
