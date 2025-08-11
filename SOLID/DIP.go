// SOLID

// Dependency Inversion Principle

// Depend on abstractions, not on concrete implementations.
// Depend on abstractions (interfaces), not concrete types.

package main

import "fmt"

type Databse interface {
	Save(data string)
}

type MySQL struct{}

func (m MySQL) Save(data string) {
	fmt.Println("Saving to mySQL : ", data)
}

type App struct {
	DB Databse
}

func (a App) SaveData(data string) {
	a.DB.Save(data)
}

func main() {
	app := App{DB: MySQL{}}
	app.SaveData("User data")
}
