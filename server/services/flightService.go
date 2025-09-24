package services

import (
	"aggregator/repositories"
	"aggregator/utils"
	"fmt"
	"slices"
	"sort"
	"time"
)

type FlightService struct {
	repos []repositories.Repos
}

func NewFlightService(repos ...repositories.Repos) *FlightService {
	return &FlightService{
		repos: repos,
	}
}

func (s *FlightService) GetAllTravel(sortByParam string, orderByParam string) []repositories.Travelling {
	allTravellings := []repositories.Travelling{}
	for _, r := range s.repos {
		travellings, _ := r.GetTravel()
		allTravellings = append(allTravellings, travellings...)
	}
	sortBy(sortByParam, orderByParam, allTravellings)
	return allTravellings
}

func sortBy(sortType string, orderby string, allTravellings []repositories.Travelling) {
	things := []string{"price", "time_travel", "departure_date"}
	if !slices.Contains(things, sortType) {
		sortType = "price"
	}

	things2 := []string{"asc", "desc"}

	if !slices.Contains(things2, orderby) {
		orderby = "asc"
	}

	switch sortType {
	case "price":
		sortByPrice(allTravellings, orderby)
	case "time_travel":
		sortByTimeTravel(allTravellings, orderby)
	case "departure_date":
		sortByDepartureDate(allTravellings, orderby)
	default:
		fmt.Println("default")
	}
	fmt.Println("sort => ", sortType)
}

func sortByPrice(allTravellings []repositories.Travelling, orderBy string) {
	sort.Slice(allTravellings, func(i, j int) bool {
		return utils.Ternary(orderBy == "asc", allTravellings[i].Total.Price < allTravellings[j].Total.Price, allTravellings[i].Total.Price > allTravellings[j].Total.Price)
	})
}

func sortByDepartureDate(allTravellings []repositories.Travelling, orderBy string) {
	sort.Slice(allTravellings, func(i, j int) bool {
		timeI, _ := time.Parse(time.RFC3339, allTravellings[i].Flights[0].Depart)
		timeJ, _ := time.Parse(time.RFC3339, allTravellings[j].Flights[0].Depart)
		utils.Ternary(orderBy == "asc", timeI.Before(timeJ), timeI.After(timeJ))
		return timeI.Before(timeJ)
	})
}

func sortByTimeTravel(allTravellings []repositories.Travelling, orderBy string) {
	sort.SliceStable(allTravellings, func(i, j int) bool {
		var iTime time.Time
		for _, v := range allTravellings[i].Flights {
			arriveTime, _ := time.Parse(time.RFC3339, v.Arrive)
			departTime, _ := time.Parse(time.RFC3339, v.Depart)
			diff := arriveTime.Sub(departTime)
			iTime = iTime.Add(diff)
		}

		var jTime time.Time
		for _, v := range allTravellings[j].Flights {
			arriveTime, _ := time.Parse(time.RFC3339, v.Arrive)
			departTime, _ := time.Parse(time.RFC3339, v.Depart)
			diff := arriveTime.Sub(departTime)
			jTime = jTime.Add(diff)
		}

		orderedResult := utils.Ternary(orderBy == "desc", iTime.After(jTime), jTime.After(iTime))

		return orderedResult
		// flightsI := allTravellings[i].Flights
		// timeIDepart, _ := time.Parse(time.RFC3339, flightsI[0].Depart)
		// timeIArrive, _ := time.Parse(time.RFC3339, flightsI[len(flightsI)-1].Arrive)
		// timeIDiff := timeIArrive.Sub(timeIDepart)

		// flightsJ := allTravellings[j].Flights
		// timeJDepart, _ := time.Parse(time.RFC3339, flightsJ[0].Depart)
		// timeJArrive, _ := time.Parse(time.RFC3339, flightsJ[len(flightsJ)-1].Arrive)
		// timeJDiff := timeJArrive.Sub(timeJDepart)

		// orderedResult := utils.Ternary(orderBy == "asc", timeIDiff > timeJDiff, timeIDiff < timeJDiff)

		// return orderedResult
	})
}
