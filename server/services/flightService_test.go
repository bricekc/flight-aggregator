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
