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
    for i, result := range travellings {
        if result.Total.Price != expected[i].Total.Price {
            t.Errorf("Expected price %f at index %d, got %f", expected[i].Total.Price, i, result.Total.Price)
        }
    }
}

func TestSortByPriceDESC(t *testing.T) {
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
    for i, result := range travellings {
        if result.Total.Price != expected[i].Total.Price {
            t.Errorf("Expected price %f at index %d, got %f", expected[i].Total.Price, i, result.Total.Price)
        }
    }
}

func TestSortByDepartureDateASC(t *testing.T) {
    // given
    travellings := []repositories.Travelling{
        {Flights: []repositories.Flight{{Depart: "2006-01-01T15:05:05Z"},},},
        {Flights: []repositories.Flight{{Depart: "2006-01-01T07:04:05Z"},},},
        {Flights: []repositories.Flight{{Depart: "2006-01-01T12:06:05Z"},},},
    }
    expected := []repositories.Travelling{
        {Flights: []repositories.Flight{{Depart: "2006-01-01T07:04:05Z"},},},
        {Flights: []repositories.Flight{{Depart: "2006-01-01T12:06:05Z"},},},
        {Flights: []repositories.Flight{{Depart: "2006-01-01T15:05:05Z"},},},
    }

    //when
    sortByDepartureDate(travellings, "asc")

    //then
    for i, result := range travellings {
        if result.Flights[0].Depart != expected[i].Flights[0].Depart {
            t.Errorf("Expected departure date %s at index %d, got %s", expected[i].Flights[0].Depart, i, result.Flights[0].Depart)
        }
    }
}

func TestSortByDepartureDateDESC(t *testing.T) {
    // given
    travellings := []repositories.Travelling{
        {Flights: []repositories.Flight{{Depart: "2006-01-01T15:05:05Z"},},},
        {Flights: []repositories.Flight{{Depart: "2006-01-01T07:04:05Z"},},},
        {Flights: []repositories.Flight{{Depart: "2006-01-01T12:06:05Z"},},},
    }
    expected := []repositories.Travelling{
		{Flights: []repositories.Flight{{Depart: "2006-01-01T15:05:05Z"},},},
		{Flights: []repositories.Flight{{Depart: "2006-01-01T12:06:05Z"},},},
        {Flights: []repositories.Flight{{Depart: "2006-01-01T07:04:05Z"},},},
    }

    //when
    sortByDepartureDate(travellings, "desc")

    //then
    for i, result := range travellings {
        if result.Flights[0].Depart != expected[i].Flights[0].Depart {
            t.Errorf("Expected departure date %s at index %d, got %s", expected[i].Flights[0].Depart, i, result.Flights[0].Depart)
        }
    }
}