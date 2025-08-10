//SOLID principle

// OCP -> Open / Closed principle

package main

import "fmt"

// Shapes interface
type shape interface {
	Area() float64
	Details() string
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Details() string {
	return fmt.Sprintf("Circle with radius %.2f", c.Radius)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Details() string {
	return fmt.Sprintf("Rectangle with Width %.2f and Height %.2f", r.Width, r.Height)
}

func main() {
	shapes := []shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
	}

	var TotalArea float64
	for _, s := range shapes {
		area := s.Area()
		fmt.Printf("%s has area %.2f\n", s.Details(), area)
		TotalArea += area
	}

	fmt.Printf("Total Area : %.2f\n", TotalArea)
}

// package main

// import "fmt"

// // Shape interface
// type shape interface {
// 	Area() float64
// }

// // Circle implements shape
// type Circle struct {
// 	Radius float64
// }

// type Rectangle struct {
// 	Width, Height float64
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// func (r Rectangle) Area() float64 {
// 	return r.Height * r.Width
// }

// func TotalArea(shapes []shape) float64 {
// 	var area float64
// 	for _, shape := range shapes {
// 		area += shape.Area()
// 	}
// 	return area
// }

// func main() {
// 	shapes := []shape{
// 		Circle{Radius: 5},
// 		Rectangle{Width: 3, Height: 4},
// 	}

// 	fmt.Println("Total Area : ", TotalArea(shapes))
// }
