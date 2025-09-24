package repositories

type Repos interface {
	GetTravel() ([]Travelling, error)
}

type Flight struct {
	Number string
	From   string
	To     string
	Depart string
	Arrive string
}

type Total struct {
	Price    float64
	Currency string
}

type Travelling struct {
	BookingId     string
	Status        string
	PassengerName string
	Flights       []Flight
	Total         Total
}
