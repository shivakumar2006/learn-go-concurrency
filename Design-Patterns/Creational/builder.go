package main

import "fmt"

type Sandwich struct {
	Bread       string
	Sauce       string
	Fillings    []string
	Toasted     bool
	ExtraCheese bool
}

func (s Sandwich) Show() {
	fmt.Println("--------- your sandwich ---------")
	fmt.Println("Bread : ", s.Bread)
	fmt.Println("Sauce : ", s.Sauce)
	fmt.Println("Fillingd : ", s.Fillings)
	fmt.Println("Toasted : ", s.Toasted)
	fmt.Println("ExtraCheese : ", s.ExtraCheese)
	fmt.Println("---------------------------------")
}

// builder interface

type SandwichBuilder interface {
	ChooseBread()
	AddSauce()
	AddFillings()
	Toasting()
	AddExtras()
	GetSandwich() Sandwich
}

// veg sandwich
type VegSandwichBuilder struct {
	sandwich Sandwich
}

func (b *VegSandwichBuilder) ChooseBread() {
	b.sandwich.Bread = "Brown bread"
}

func (b *VegSandwichBuilder) AddSauce() {
	b.sandwich.Sauce = "Mint mayo"
}

func (b *VegSandwichBuilder) AddFillings() {
	b.sandwich.Fillings = []string{"Lettuce", "Tomato", "Cucumber", "Onion"}
}

func (b *VegSandwichBuilder) Toasting() {
	b.sandwich.Toasted = true
}

func (b *VegSandwichBuilder) AddExtras() {
	b.sandwich.ExtraCheese = false
}

func (b *VegSandwichBuilder) GetSandwich() Sandwich {
	return b.sandwich
}

// chicken sandwich
type ChickenSandwichBuilder struct {
	sandwich Sandwich
}

func (b *ChickenSandwichBuilder) ChooseBread() {
	b.sandwich.Bread = "Italian bread"
}

func (b *ChickenSandwichBuilder) AddSauce() {
	b.sandwich.Sauce = "Chipotle southwest"
}

func (b *ChickenSandwichBuilder) AddFillings() {
	b.sandwich.Fillings = []string{"Chicken Strips", "Lettuce", "Onion"}
}

func (b *ChickenSandwichBuilder) Toasting() {
	b.sandwich.Toasted = true
}

func (b *ChickenSandwichBuilder) AddExtras() {
	b.sandwich.ExtraCheese = true
}

func (b *ChickenSandwichBuilder) GetSandwich() Sandwich {
	return b.sandwich
}

// subway sandwich
type SubwayArtist struct {
	builder SandwichBuilder
}

func (a *SubwayArtist) SetBuilder(builder SandwichBuilder) {
	a.builder = builder
}

func (a *SubwayArtist) MakeSandwich() Sandwich {
	a.builder.ChooseBread()
	a.builder.AddSauce()
	a.builder.AddFillings()
	a.builder.Toasting()
	a.builder.AddExtras()
	return a.builder.GetSandwich()
}

func main() {
	artist := SubwayArtist{}

	vegBuilder := &VegSandwichBuilder{}
	artist.SetBuilder(vegBuilder)
	vegSub := artist.MakeSandwich()
	vegSub.Show()

	chickenBuilder := &ChickenSandwichBuilder{}
	artist.SetBuilder(chickenBuilder)
	chickenSub := artist.MakeSandwich()
	chickenSub.Show()
}
