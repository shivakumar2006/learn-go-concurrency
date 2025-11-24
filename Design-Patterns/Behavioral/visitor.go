package main

import "fmt"

type Shape interface {
	getType() string
	accept(Visitor)
}

// square
type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "square"
}

// circle
type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "circle"
}

// rectangle
type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "rectangle"
}

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// area calculator
type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("calculating area for square")
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	fmt.Println("calculating area for circle")
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("calculaitng area for rectangle")
}

// middle coordinates
type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("calculate middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("calculating middle point coordinates for circle")
}

func (a *MiddleCoordinates) visitForRectangle(r *Rectangle) {
	fmt.Println("calculating middle pount coordinates for rectangle")
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
