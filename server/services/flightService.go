package services

import "aggregator/repositories"

type FlightService struct {
    repos []repositories.Repos
}

func NewFlightService(repos ...repositories.Repos) *FlightService {
    return &FlightService{
        repos: repos,
    }
}

func (s *FlightService) GetAllTravel() []repositories.Travelling {
    allTravellings := []repositories.Travelling{}
    for _, r := range s.repos {
        travellings := r.GetTravel()
        allTravellings = append(allTravellings, travellings...)
    }
    return allTravellings
}