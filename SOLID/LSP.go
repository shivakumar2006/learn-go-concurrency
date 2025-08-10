// SOLID

//LISKOV SUBSTITUTION PRINCIPLE

// Objects of a superclass should be replaceable with objects of a subclass without breaking correctness.

package main

import "fmt"

type Birds interface {
	Walk()
}

type FlyingBirds interface {
	Birds
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

func (s Sparrow) Walk() {
	fmt.Println("Sparrow is walking")
}

type Penguin struct{}

func (p Penguin) Walk() {
	fmt.Println("Penguin is walking")
}

func letBirdFLy(b FlyingBirds) {
	b.Fly()
}

func main() {
	s := Sparrow{}
	letBirdFLy(s)
	s.Walk()

	p := Penguin{}
	p.Walk()
}

// so, this is lsp and how we write code how to solve this
// package main

// import "fmt"

// type Birds interface {
// 	Fly()
// }

// type Sparrow struct{}

// func (s Sparrow) Fly() {
// 	fmt.Println("Sparrow is flying")
// }

// // Bad : Penguin is a bird but cannot fly
// type Penguin struct{}

// func (p Penguin) Fly() {
// 	panic("Penguin can't fly") // violates LSP
// }

// func letBirdFLy(b Birds) {
// 	b.Fly()
// }

// func main() {
// 	letBirdFLy(Sparrow{})
// 	// letBirdFLy(Penguin{})
// }
