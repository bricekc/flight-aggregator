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

func TestSortByPrice(t *testing.T) {
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

func TestSortByDepartureDate(t *testing.T) {
	t.Run("should sort by departure date asc", func(t *testing.T) {
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
	assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after asc departure date sorting")
	})
	t.Run("should sort by departure date desc", func(t *testing.T) {
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
	assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after desc departure date sorting")
	})
}

func TestSortBy(t *testing.T) {
	t.Run("should sort by default (price asc)", func(t *testing.T) {
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
	sortBy("", "", travellings)

	// then
	assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after default price sorting")
	})

	t.Run("should sort by price asc when price param is given", func(t *testing.T) {
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
	sortBy("price", "asc", travellings)

	// then
	assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after price asc sorting")
	})

	t.Run("should sort by price desc when price param is given", func(t *testing.T) {
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
	sortBy("price", "desc", travellings)

	// then
	assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after price desc sorting")
	})

	t.Run("should sort by departure date asc when departure_date param is given", func(t *testing.T) {
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
	sortBy("departure_date", "asc", travellings)

	//then
	assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after asc departure date sorting")
	})

	t.Run("should sort by departure date desc when departure_date param is given", func(t *testing.T) {
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
	sortBy("departure_date", "desc", travellings)

	//then
	assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after desc departure date sorting")
	})

	t.Run("should sort by time travel asc when time_travel param is given", func(t *testing.T) {
	// given
	travellings := []repositories.Travelling{
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
	expected := []repositories.Travelling{
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
	sortBy("time_travel", "asc", travellings)
	
	// then
	assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after asc travel time sorting")
	})

	t.Run("should sort by time travel desc when time_travel param is given", func(t *testing.T) {
	// given
	travellings := []repositories.Travelling{
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
	expected := []repositories.Travelling{
		// plus long
		{Flights: []repositories.Flight{
			{Depart: "2026-01-01T04:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
			{Depart: "2026-01-01T16:00:00Z", Arrive: "2026-01-02T10:00:00Z"},
		}},
		// milieu
		{Flights: []repositories.Flight{{
			Depart: "2026-01-01T01:00:00Z", Arrive: "2026-01-01T23:00:00Z"}}},
		// plus court
		{Flights: []repositories.Flight{
			{Depart: "2026-01-01T10:00:00Z", Arrive: "2026-01-01T12:00:00Z"},
			{Depart: "2026-01-02T02:00:00Z", Arrive: "2026-01-02T06:00:00Z"},
		}},
	}

	// when
	sortBy("time_travel", "desc", travellings)
	
	// then
	assert.Equal(t, travellings, expected, "The two Travelling should be in the same order after desc travel time sorting")
	})
}