package services

import (
	"aggregator/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByTimeTravel(t *testing.T) {
	t.Run("should sort by travel time desc", func(t *testing.T) {
		// given
		travellingsDesc := []repositories.Travelling{
			// milieu
			{Flights: []repositories.Flight{{Depart: "2026-01-01T01:00:00Z", Arrive: "2026-01-01T23:00:00Z"}}},
			// plus long
			{Flights: []repositories.Flight{
				{Depart: "2026-01-01T04:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
				{Depart: "2026-01-01T16:00:00Z", Arrive: "2026-01-02T10:00:00Z"},
			}},
			// plus court
			{Flights: []repositories.Flight{
				{Depart: "2026-01-01T10:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
				{Depart: "2026-01-02T02:00:00Z", Arrive: "2026-01-02T06:00:00Z"},
			}},
		}
		expectedDesc := []repositories.Travelling{
			// plus long
			{Flights: []repositories.Flight{
				{Depart: "2026-01-01T04:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
				{Depart: "2026-01-01T16:00:00Z", Arrive: "2026-01-02T10:00:00Z"},
			}},
			// milieu
			{Flights: []repositories.Flight{{Depart: "2026-01-01T01:00:00Z", Arrive: "2026-01-01T23:00:00Z"}}},
			// plus court
			{Flights: []repositories.Flight{
				{Depart: "2026-01-01T10:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
				{Depart: "2026-01-02T02:00:00Z", Arrive: "2026-01-02T06:00:00Z"},
			}},
		}

		// when
		sortByTimeTravel(travellingsDesc, "desc")

		// then
		assert.Equal(t, travellingsDesc, expectedDesc, "The two Travelling should be in the same order after desc traveil time sorting")
	})
	t.Run("should sort by travel time asc", func(t *testing.T) {
		// given
		travellingsAsc := []repositories.Travelling{
			// milieu
			{Flights: []repositories.Flight{{Depart: "2026-01-01T01:00:00Z", Arrive: "2026-01-01T23:00:00Z"}}},
			// plus long
			{Flights: []repositories.Flight{
				{Depart: "2026-01-01T04:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
				{Depart: "2026-01-01T16:00:00Z", Arrive: "2026-01-02T10:00:00Z"},
			}},
			// plus court
			{Flights: []repositories.Flight{
				{Depart: "2026-01-01T10:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
				{Depart: "2026-01-02T02:00:00Z", Arrive: "2026-01-02T06:00:00Z"},
			}},
		}
		expectedAsc := []repositories.Travelling{
			// plus court
			{Flights: []repositories.Flight{
				{Depart: "2026-01-01T10:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
				{Depart: "2026-01-02T02:00:00Z", Arrive: "2026-01-02T06:00:00Z"},
			}},
			// milieu
			{Flights: []repositories.Flight{{Depart: "2026-01-01T01:00:00Z", Arrive: "2026-01-01T23:00:00Z"}}},
			// plus long
			{Flights: []repositories.Flight{
				{Depart: "2026-01-01T04:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
				{Depart: "2026-01-01T16:00:00Z", Arrive: "2026-01-02T10:00:00Z"},
			}},
		}

		// when
		sortByTimeTravel(travellingsAsc, "asc")

		// then
		assert.Equal(t, travellingsAsc, expectedAsc, "The two Travelling should be in the same order after asc traveil time sorting")
	})
}

func TestSortByPriceASC(t *testing.T) {
	t.Run("should sort by price asc", func(t *testing.T) {
		// given
		travellings := []repositories.Travelling{
			{Total: repositories.Total{Price: 300}},
			{Total: repositories.Total{Price: 100}},
			{Total: repositories.Total{Price: 200}},
		}
		expected := []repositories.Travelling{
			{Total: repositories.Total{Price: 100}},
			{Total: repositories.Total{Price: 200}},
			{Total: repositories.Total{Price: 300}},
		}

		// when
		sortByPrice(travellings, "asc")

		// then
		assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after asc price sorting")
	})
}

func TestSortByPriceDESC(t *testing.T) {
	t.Run("should sort by price desc", func(t *testing.T) {
		// given
		travellings := []repositories.Travelling{
			{Total: repositories.Total{Price: 300}},
			{Total: repositories.Total{Price: 100}},
			{Total: repositories.Total{Price: 200}},
		}
		expected := []repositories.Travelling{
			{Total: repositories.Total{Price: 300}},
			{Total: repositories.Total{Price: 200}},
			{Total: repositories.Total{Price: 100}},
		}

		// when
		sortByPrice(travellings, "desc")

		// then
		assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after desc price sorting")
	})
}

func TestSortByDepartureDateASC(t *testing.T) {
	t.Run("should sort by departure date asc", func(t *testing.T) {
		// given
		travellings := []repositories.Travelling{
			{Flights: []repositories.Flight{{Depart: "2006-01-01T15:05:05Z"}}},
			{Flights: []repositories.Flight{{Depart: "2006-01-01T07:04:05Z"}}},
			{Flights: []repositories.Flight{{Depart: "2006-01-01T12:06:05Z"}}},
		}
		expected := []repositories.Travelling{
			{Flights: []repositories.Flight{{Depart: "2006-01-01T07:04:05Z"}}},
			{Flights: []repositories.Flight{{Depart: "2006-01-01T12:06:05Z"}}},
			{Flights: []repositories.Flight{{Depart: "2006-01-01T15:05:05Z"}}},
		}

		//when
		sortByDepartureDate(travellings, "asc")

		//then
		assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after asc departure date sorting")
	})
}

func TestSortByDepartureDateDESC(t *testing.T) {
	t.Run("should sort by departure date desc", func(t *testing.T) {
		// given
		travellings := []repositories.Travelling{
			{Flights: []repositories.Flight{{Depart: "2006-01-01T15:05:05Z"}}},
			{Flights: []repositories.Flight{{Depart: "2006-01-01T07:04:05Z"}}},
			{Flights: []repositories.Flight{{Depart: "2006-01-01T12:06:05Z"}}},
		}
		expected := []repositories.Travelling{
			{Flights: []repositories.Flight{{Depart: "2006-01-01T15:05:05Z"}}},
			{Flights: []repositories.Flight{{Depart: "2006-01-01T12:06:05Z"}}},
			{Flights: []repositories.Flight{{Depart: "2006-01-01T07:04:05Z"}}},
		}

		//when
		sortByDepartureDate(travellings, "desc")

		//then
		assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after desc departure date sorting")
	})
}

type MockTravelRepo struct {
	travels []repositories.Travelling
}

func (m *MockTravelRepo) GetTravel() ([]repositories.Travelling, error) {
	return m.travels, nil
}

func TestGetAllTravel(t *testing.T) {
	t.Run("should have travel from two differents repos", func(t *testing.T) {
		// given
		var repo1 = &MockTravelRepo{
			travels: []repositories.Travelling{
				{
					BookingId:     "B30001",
					Status:        "confirmed",
					PassengerName: "Marie Curie",
					Flights: []repositories.Flight{
						{
							Number: "AF276",
							From:   "",
							To:     "",
							Depart: "2026-01-01T10:00:00Z",
							Arrive: "2026-01-01T23:00:00Z",
						},
					},
					Total: repositories.Total{
						Price:    950,
						Currency: "EUR",
					},
				},
				{
					BookingId:     "B30004",
					Status:        "confirmed",
					PassengerName: "Marie Curie",
					Flights: []repositories.Flight{
						{
							Number: "KE902",
							From:   "",
							To:     "",
							Depart: "2026-01-01T09:30:00Z",
							Arrive: "2026-01-01T18:00:00Z",
						},
						{
							Number: "KE711",
							From:   "",
							To:     "",
							Depart: "2026-01-01T20:00:00Z",
							Arrive: "2026-01-02T00:30:00Z",
						},
					},
					Total: repositories.Total{
						Price:    880,
						Currency: "EUR",
					},
				},
			},
		}
		var repo2 = &MockTravelRepo{
			travels: []repositories.Travelling{
				{
					BookingId:     "A10002",
					Status:        "confirmed",
					PassengerName: "Albert Einstein",
					Flights: []repositories.Flight{
						{
							Number: "AF276",
							From:   "CDG",
							To:     "HND",
							Depart: "2026-01-01T10:30:00Z",
							Arrive: "2026-01-02T06:00:00Z",
						},
					},
					Total: repositories.Total{
						Price:    910,
						Currency: "EUR",
					},
				},
			},
		}
		expected := []repositories.Travelling{
			{
				BookingId:     "B30004",
				Status:        "confirmed",
				PassengerName: "Marie Curie",
				Flights: []repositories.Flight{
					{
						Number: "KE902",
						From:   "",
						To:     "",
						Depart: "2026-01-01T09:30:00Z",
						Arrive: "2026-01-01T18:00:00Z",
					},
					{
						Number: "KE711",
						From:   "",
						To:     "",
						Depart: "2026-01-01T20:00:00Z",
						Arrive: "2026-01-02T00:30:00Z",
					},
				},
				Total: repositories.Total{
					Price:    880,
					Currency: "EUR",
				},
			},
			{
				BookingId:     "A10002",
				Status:        "confirmed",
				PassengerName: "Albert Einstein",
				Flights: []repositories.Flight{
					{
						Number: "AF276",
						From:   "CDG",
						To:     "HND",
						Depart: "2026-01-01T10:30:00Z",
						Arrive: "2026-01-02T06:00:00Z",
					},
				},
				Total: repositories.Total{
					Price:    910,
					Currency: "EUR",
				},
			},
			{
				BookingId:     "B30001",
				Status:        "confirmed",
				PassengerName: "Marie Curie",
				Flights: []repositories.Flight{
					{
						Number: "AF276",
						From:   "",
						To:     "",
						Depart: "2026-01-01T10:00:00Z",
						Arrive: "2026-01-01T23:00:00Z",
					},
				},
				Total: repositories.Total{
					Price:    950,
					Currency: "EUR",
				},
			},
		}

		// when
		flightService := NewFlightService(repo1, repo2)
		allTravellings := flightService.GetAllTravel("", "")

		// then
		assert.Equal(t, expected, allTravellings, "Two repos being merge into one slice")
	})
}
