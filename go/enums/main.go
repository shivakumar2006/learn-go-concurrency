package main

import "fmt"

type OrderStatus string

const (
	Received  OrderStatus = "Order Received"
	Confirmed OrderStatus = "Order Confirmed"
	Shipped   OrderStatus = "Order Shipped"
	Delivered OrderStatus = "Order Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to : ", status)
}

func main() {
	changeOrderStatus(Shipped)
}
