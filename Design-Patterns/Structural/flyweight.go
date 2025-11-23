package main

import "fmt"

const (
	TerroristDressType       = "tDress"
	CounterTerroristDatatype = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDatatype {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

// FlyWeight interface
type Dress interface {
	getColor() string
}

// concrete flyweight object
type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{
		color: "red",
	}
}

type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{
		color: "blue",
	}
}

// player
type Player struct {
	dress      Dress
	playerType string
	lat        int
	lon        int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, lon int) {
	p.lat = lat
	p.lon = lon
}

// game
type game struct {
	terrorist        []*Player
	counterTerrorist []*Player
}

func newGame() *game {
	return &game{
		terrorist:        make([]*Player, 1),
		counterTerrorist: make([]*Player, 1),
	}
}

func (t *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	t.terrorist = append(t.terrorist, player)
	return
}

func (t *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	t.counterTerrorist = append(t.counterTerrorist, player)
	return
}

func main() {
	game := newGame()

	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	game.addCounterTerrorist(CounterTerroristDatatype)
	game.addCounterTerrorist(CounterTerroristDatatype)
	game.addCounterTerrorist(CounterTerroristDatatype)
	game.addCounterTerrorist(CounterTerroristDatatype)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
