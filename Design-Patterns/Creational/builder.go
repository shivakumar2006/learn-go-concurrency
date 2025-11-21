// package main

package main

import "fmt"

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}
	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "wooden window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "wooden door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "snow window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "snow door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal house door type : %s\n", normalHouse.doorType)
	fmt.Printf("Normal house window type : %s\n", normalHouse.windowType)
	fmt.Printf("Normal house floor : %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo house door type : %s\n", iglooHouse.doorType)
	fmt.Printf("\nIgloo window door type : %s\n", iglooHouse.windowType)
	fmt.Printf("\nIgloo set floor : %d\n", iglooHouse.floor)
}

// import "fmt"

// type Sandwich struct {
// 	Bread       string
// 	Sauce       string
// 	Fillings    []string
// 	Toasted     bool
// 	ExtraCheese bool
// }

// func (s Sandwich) Show() {
// 	fmt.Println("--------- your sandwich ---------")
// 	fmt.Println("Bread : ", s.Bread)
// 	fmt.Println("Sauce : ", s.Sauce)
// 	fmt.Println("Fillingd : ", s.Fillings)
// 	fmt.Println("Toasted : ", s.Toasted)
// 	fmt.Println("ExtraCheese : ", s.ExtraCheese)
// 	fmt.Println("---------------------------------")
// }

// // builder interface

// type SandwichBuilder interface {
// 	ChooseBread()
// 	AddSauce()
// 	AddFillings()
// 	Toasting()
// 	AddExtras()
// 	GetSandwich() Sandwich
// }

// // veg sandwich
// type VegSandwichBuilder struct {
// 	sandwich Sandwich
// }

// func (b *VegSandwichBuilder) ChooseBread() {
// 	b.sandwich.Bread = "Brown bread"
// }

// func (b *VegSandwichBuilder) AddSauce() {
// 	b.sandwich.Sauce = "Mint mayo"
// }

// func (b *VegSandwichBuilder) AddFillings() {
// 	b.sandwich.Fillings = []string{"Lettuce", "Tomato", "Cucumber", "Onion"}
// }

// func (b *VegSandwichBuilder) Toasting() {
// 	b.sandwich.Toasted = true
// }

// func (b *VegSandwichBuilder) AddExtras() {
// 	b.sandwich.ExtraCheese = false
// }

// func (b *VegSandwichBuilder) GetSandwich() Sandwich {
// 	return b.sandwich
// }

// // chicken sandwich
// type ChickenSandwichBuilder struct {
// 	sandwich Sandwich
// }

// func (b *ChickenSandwichBuilder) ChooseBread() {
// 	b.sandwich.Bread = "Italian bread"
// }

// func (b *ChickenSandwichBuilder) AddSauce() {
// 	b.sandwich.Sauce = "Chipotle southwest"
// }

// func (b *ChickenSandwichBuilder) AddFillings() {
// 	b.sandwich.Fillings = []string{"Chicken Strips", "Lettuce", "Onion"}
// }

// func (b *ChickenSandwichBuilder) Toasting() {
// 	b.sandwich.Toasted = true
// }

// func (b *ChickenSandwichBuilder) AddExtras() {
// 	b.sandwich.ExtraCheese = true
// }

// func (b *ChickenSandwichBuilder) GetSandwich() Sandwich {
// 	return b.sandwich
// }

// // subway sandwich
// type SubwayArtist struct {
// 	builder SandwichBuilder
// }

// func (a *SubwayArtist) SetBuilder(builder SandwichBuilder) {
// 	a.builder = builder
// }

// func (a *SubwayArtist) MakeSandwich() Sandwich {
// 	a.builder.ChooseBread()
// 	a.builder.AddSauce()
// 	a.builder.AddFillings()
// 	a.builder.Toasting()
// 	a.builder.AddExtras()
// 	return a.builder.GetSandwich()
// }

// func main() {
// 	artist := SubwayArtist{}

// 	vegBuilder := &VegSandwichBuilder{}
// 	artist.SetBuilder(vegBuilder)
// 	vegSub := artist.MakeSandwich()
// 	vegSub.Show()

// 	chickenBuilder := &ChickenSandwichBuilder{}
// 	artist.SetBuilder(chickenBuilder)
// 	chickenSub := artist.MakeSandwich()
// 	chickenSub.Show()
// }
