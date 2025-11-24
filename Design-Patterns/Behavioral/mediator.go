package main

import "fmt"

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}

// passenger train
type PassengerTrain struct {
	mediator Mediator
}

func (p *PassengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("Pasenger train: arrival blocked, waiting!")
		return
	}
	fmt.Println("Passenger arrived")
}

func (p *PassengerTrain) depart() {
	fmt.Println("Passenger train: Leaving")
	p.mediator.notifyAboutDeparture()
}

func (p *PassengerTrain) permitArrival() {
	fmt.Println("Passenger train: arrival permitted, arriving")
	p.arrive()
}

// freight train
type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("Freight train: arrival blocked, waiting!")
		return
	}
	fmt.Println("Freight train: arrived")
}

func (g *FreightTrain) depart() {
	fmt.Println("Freight train: leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("Freight train: arrival permitted")
	g.arrive()
}

// station manager
type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainQueue.permitArrival()
	}
}

func main() {
	stationManager := &StationManager{}

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}

	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
