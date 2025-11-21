// //factory method pattern

package main

import "fmt"

// profuct interface
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// concrete products
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

// concrete product
type AK47 struct {
	Gun
}

func newAk47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// concrete product
type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// factory
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

// client code
func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun : %s", g.getName())
	fmt.Println()
	fmt.Printf("Gun : %d", g.getPower())
	fmt.Println()
}

// package main

// import "fmt"

// // product interface
// type IceCream interface {
// 	Flavour() string
// }

// // concrete icecream products
// type Choclate struct{}

// func (c Choclate) Flavour() string {
// 	return "Choclate Ice-cream"
// }

// type Vanilla struct{}

// func (v Vanilla) Flavour() string {
// 	return "Vanilla ice cream"
// }

// type Strawberry struct{}

// func (s Strawberry) Flavour() string {
// 	return "Strawberry ice cream"
// }

// // factory method
// func IceCreamFactory(choice string) IceCream {
// 	switch choice {
// 	case "choclate":
// 		return Choclate{}
// 	case "vanilla":
// 		return Vanilla{}
// 	case "strawberry":
// 		return Strawberry{}
// 	default:
// 		return nil
// 	}
// }

// // client (customer)
// func main() {
// 	ice1 := IceCreamFactory("choclate")
// 	ice2 := IceCreamFactory("strawberry")

// 	fmt.Println("Customer 1 got: ", ice1.Flavour())
// 	fmt.Println("Customer 2 got: ", ice2.Flavour())
// }
