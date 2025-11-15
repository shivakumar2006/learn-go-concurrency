//Abstract factory method

package main

import "fmt"

// food
type Food interface {
	ServeFood() string
}

type Drink interface {
	ServeDrink() string
}

type Dessert interface {
	ServeDessert() string
}

// restaurant items
// ----- italian restaurant items ------

type ItalianFood struct{}

func (ItalianFood) ServeFood() string { return "serving italian pizza" }

type ItalianDrink struct{}

func (ItalianDrink) ServeDrink() string { return "serving italian wine" }

type ItalianDessert struct{}

func (ItalianDessert) ServeDessert() string { return "serving italian Gelato" }

// ----- indian restaurant items ------

type IndianFood struct{}

func (IndianFood) ServeFood() string { return "serving indian Birayni" }

type IndianDrink struct{}

func (IndianDrink) ServeDrink() string { return "serving indian Lassi" }

type IndianDessert struct{}

func (IndianDessert) ServeDessert() string { return "serving indian Gulab jamun" }

// abstract factory method
type RestaurantFactory interface {
	CreateFood() Food
	CreateDrink() Drink
	CreateDessert() Dessert
}

// concrete factory
type ItalianRestaurant struct{}

func (ItalianRestaurant) CreateFood() Food       { return ItalianFood{} }
func (ItalianRestaurant) CreateDrink() Drink     { return ItalianDrink{} }
func (ItalianRestaurant) CreateDessert() Dessert { return ItalianDessert{} }

type IndianRestaurant struct{}

func (IndianRestaurant) CreateFood() Food       { return IndianFood{} }
func (IndianRestaurant) CreateDrink() Drink     { return IndianDrink{} }
func (IndianRestaurant) CreateDessert() Dessert { return IndianDessert{} }

// client (customer)
func ServeFullMeal(factory RestaurantFactory) {
	food := factory.CreateFood()
	drink := factory.CreateDrink()
	dessert := factory.CreateDessert()

	fmt.Println(food.ServeFood())
	fmt.Println(drink.ServeDrink())
	fmt.Println(dessert.ServeDessert())
}

func main() {
	fmt.Println("Italian restaurant meal : ")
	ServeFullMeal(ItalianRestaurant{})

	fmt.Println("Indian restaurant meal : ")
	ServeFullMeal(IndianRestaurant{})
}
