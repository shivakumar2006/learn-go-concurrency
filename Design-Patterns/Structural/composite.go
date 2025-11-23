package main

import "fmt"

type Component interface {
	Search(string)
}

// composite
type Folder struct {
	components []Component
	name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

// lead
type File struct {
	name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

func main() {
	file1 := &File{name: "FILE1"}
	file2 := &File{name: "FILE2"}
	file3 := &File{name: "FILE3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}

	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.Search("rose")
}
