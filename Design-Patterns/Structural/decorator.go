package main

import "fmt"

type IPizza interface {
	getPrice() int
}

// concrete component
type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15
}

// toppings : concrete decorator
type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 7 // means getPrice is already 15 and now we add 7 : 15+7=22
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10 // now getprice is already 22 and we add more 10 into it : 22 + 10 = 32
}

func main() {
	pizza := &VeggieMania{}

	// add cheese toppings
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggie mania with tomato and cheese topping is %d \n", pizzaWithCheeseAndTomato.getPrice())
}
