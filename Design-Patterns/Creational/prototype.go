// prototype pattern

package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

// concrete prototype book

type Book struct {
	ID       int
	Title    string
	Author   string
	Category string
}

func (b *Book) Clone() Prototype {
	// deep clone
	copy := *b
	return &copy
}

func (b *Book) Print() {
	fmt.Printf("Book => ID: %d, Title: %s, Author: %s, Category: %s\n", b.ID, b.Title, b.Author, b.Category)
}

// client
func main() {
	// original object
	original := &Book{
		ID:       1,
		Title:    "GO programming",
		Author:   "shivakumar",
		Category: "Programming",
	}

	fmt.Println("Original : ")
	original.Print()

	// clone the object
	clone := original.Clone().(*Book)
	clone.ID = 2
	clone.Title = "Advanced go programming"

	fmt.Println("\nCloned: ")
	clone.Print()

	fmt.Println("\nOriginal after clone (unchanged)")
	original.Print()
}
