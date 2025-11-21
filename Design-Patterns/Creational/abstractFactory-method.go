// //Abstract factory method

package main

import "fmt"

// abstract product
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}

// concrete products
type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

type NikeShoe struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

// abstract factory interface
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

// concrete factory
type Adidas struct{}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

type Nike struct{}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 12,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 12,
		},
	}
}

// client code
func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Shoe Logo : %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Shoe Size : %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Shirt Logo : %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Shirt Size : %d", s.getSize())
	fmt.Println()
}

// package main

// import "fmt"

// // food
// type Food interface {
// 	ServeFood() string
// }

// type Drink interface {
// 	ServeDrink() string
// }

// type Dessert interface {
// 	ServeDessert() string
// }

// // restaurant items
// // ----- italian restaurant items ------

// type ItalianFood struct{}

// func (ItalianFood) ServeFood() string { return "serving italian pizza" }

// type ItalianDrink struct{}

// func (ItalianDrink) ServeDrink() string { return "serving italian wine" }

// type ItalianDessert struct{}

// func (ItalianDessert) ServeDessert() string { return "serving italian Gelato" }

// // ----- indian restaurant items ------

// type IndianFood struct{}

// func (IndianFood) ServeFood() string { return "serving indian Birayni" }

// type IndianDrink struct{}

// func (IndianDrink) ServeDrink() string { return "serving indian Lassi" }

// type IndianDessert struct{}

// func (IndianDessert) ServeDessert() string { return "serving indian Gulab jamun" }

// // abstract factory method
// type RestaurantFactory interface {
// 	CreateFood() Food
// 	CreateDrink() Drink
// 	CreateDessert() Dessert
// }

// // concrete factory
// type ItalianRestaurant struct{}

// func (ItalianRestaurant) CreateFood() Food       { return ItalianFood{} }
// func (ItalianRestaurant) CreateDrink() Drink     { return ItalianDrink{} }
// func (ItalianRestaurant) CreateDessert() Dessert { return ItalianDessert{} }

// type IndianRestaurant struct{}

// func (IndianRestaurant) CreateFood() Food       { return IndianFood{} }
// func (IndianRestaurant) CreateDrink() Drink     { return IndianDrink{} }
// func (IndianRestaurant) CreateDessert() Dessert { return IndianDessert{} }

// // client (customer)
// func ServeFullMeal(factory RestaurantFactory) {
// 	food := factory.CreateFood()
// 	drink := factory.CreateDrink()
// 	dessert := factory.CreateDessert()

// 	fmt.Println(food.ServeFood())
// 	fmt.Println(drink.ServeDrink())
// 	fmt.Println(dessert.ServeDessert())
// }

// func main() {
// 	fmt.Println("Italian restaurant meal : ")
// 	ServeFullMeal(ItalianRestaurant{})

// 	fmt.Println("Indian restaurant meal : ")
// 	ServeFullMeal(IndianRestaurant{})
// }
