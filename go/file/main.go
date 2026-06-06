package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(f))

	//read file
	// f, err := os.Open("file.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 12)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// f, err := os.Open("file.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name : ", fileInfo.Name())

}
