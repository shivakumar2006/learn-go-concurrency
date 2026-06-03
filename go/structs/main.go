package main

import (
	"fmt"
	"time"
)

type Customer struct {
	Name  string
	Phone string
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer  Customer
}

func (o *Order) changeStatus(status string) {
	o.status = status
}

func (o *Order) getAmount() float32 {
	return o.amount
}

func NewOrder(id string, amount float32, status string) *Order {
	return &Order{
		id:     id,
		amount: amount,
		status: status,
	}
}

func main() {

	order := NewOrder("1", 50.00, "confirmed")
	order.customer = Customer{
		Name:  "shiva",
		Phone: "8888888888",
	}

	fmt.Println(order.customer)

	// order := Order{
	// 	id:        "1",
	// 	amount:    100,
	// 	status:    "received",
	// 	createdAt: time.Now(),
	// }

	// order.changeStatus("Confirmed")

	// fmt.Println(order.getAmount())
}

// func main() {
// 	order := Order{
// 		id:        "1",
// 		amount:    100,
// 		status:    "received",
// 		createdAt: time.Now(),
// 	}

// 	fmt.Println(order.status)

// 	order2 := Order{
// 		id:        "2",
// 		amount:    100,
// 		status:    "delivered",
// 		createdAt: time.Now(),
// 	}

// 	fmt.Println("Ordered struct", order2)
// 	fmt.Println("Ordered struct", order)
// }
