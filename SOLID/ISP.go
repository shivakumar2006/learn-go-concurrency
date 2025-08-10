// SOLID

// Interface Segregation Principle
// Use small, specific interfaces rather than one large interface.
// Clients should not be forced to depend on interfaces they don't use. Use many specific interfaces instead of a large general one.

package main

import (
	"fmt"
)

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type allInOnePrinter struct{}

func (a allInOnePrinter) Print() {
	fmt.Println("Printing document")
}

func (a allInOnePrinter) Scan() {
	fmt.Println("Scanning document")
}

func main() {
	p := allInOnePrinter{}
	p.Print()

	s := allInOnePrinter{}
	s.Scan()
}
