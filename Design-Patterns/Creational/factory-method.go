//factory method pattern

package main

import "fmt"

// product interface
type IceCream interface {
	Flavour() string
}

// concrete icecream products
type Choclate struct{}

func (c Choclate) Flavour() string {
	return "Choclate Ice-cream"
}

type Vanilla struct{}

func (v Vanilla) Flavour() string {
	return "Vanilla ice cream"
}

type Strawberry struct{}

func (s Strawberry) Flavour() string {
	return "Strawberry ice cream"
}

// factory method
func IceCreamFactory(choice string) IceCream {
	switch choice {
	case "choclate":
		return Choclate{}
	case "vanilla":
		return Vanilla{}
	case "strawberry":
		return Strawberry{}
	default:
		return nil
	}
}

// client (customer)
func main() {
	ice1 := IceCreamFactory("choclate")
	ice2 := IceCreamFactory("strawberry")

	fmt.Println("Customer 1 got: ", ice1.Flavour())
	fmt.Println("Customer 2 got: ", ice2.Flavour())
}
