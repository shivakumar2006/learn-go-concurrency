// // prototype pattern

package main

import "fmt"

type Inode interface {
	print(string)
	clone() Inode
}

// concrte prototype
type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
	return &File{
		name: f.name + "_clone",
	}
}

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + "  ")
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{
		name: f.name + "_clone",
	}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}

	fmt.Println("\nPrint hierarchy for folder2")
	folder2.print(" ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone folder")
	cloneFolder.print(" ")
}

// package main

// import "fmt"

// type Prototype interface {
// 	Clone() Prototype
// }

// // concrete prototype book

// type Book struct {
// 	ID       int
// 	Title    string
// 	Author   string
// 	Category string
// }

// func (b *Book) Clone() Prototype {
// 	// deep clone
// 	copy := *b
// 	return &copy
// }

// func (b *Book) Print() {
// 	fmt.Printf("Book => ID: %d, Title: %s, Author: %s, Category: %s\n", b.ID, b.Title, b.Author, b.Category)
// }

// // client
// func main() {
// 	// original object
// 	original := &Book{
// 		ID:       1,
// 		Title:    "GO programming",
// 		Author:   "shivakumar",
// 		Category: "Programming",
// 	}

// 	fmt.Println("Original : ")
// 	original.Print()

// 	// clone the object
// 	clone := original.Clone().(*Book)
// 	clone.ID = 2
// 	clone.Title = "Advanced go programming"

// 	fmt.Println("\nCloned: ")
// 	clone.Print()

// 	fmt.Println("\nOriginal after clone (unchanged)")
// 	original.Print()
// }
