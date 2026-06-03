package main

import "fmt"

type Paymenter interface {
	pay(amount float32)
	refund(amount float32, reason string) error
}

type Payment struct {
	gateway Paymenter
}

type Razorpay struct {
}

func (r *Razorpay) pay(amount float32) {
	fmt.Println("Making payment using razorpay", amount)
}

// fake payment
type FakePayment struct {
}

func (f *FakePayment) pay(amount float32) {
	fmt.Println("making payment using fake payment", amount)
}

func (f *FakePayment) refund(amount float32, reason string) error {
	fmt.Println("making refund of fake payment", amount, reason)
	return nil
}

// type Stripe struct {
// }

// func (s *Stripe) Pay(amount float32) {
// 	fmt.Println("Making payment using stripe", amount)
// }

func (p *Payment) makePayment(amount float32, refund string) {
	p.gateway.pay(amount)
	p.gateway.refund(amount, refund)
}

func main() {
	// razorpayPaymentGW := Razorpay{}
	fakePaymentGW := FakePayment{}
	newpayment := Payment{
		gateway: &fakePaymentGW,
	}
	newpayment.makePayment(100, "refund")
}
