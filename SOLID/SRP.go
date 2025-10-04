// // SOLID

// // S := Single Responsibility Principle

// A class (or module) should have only one reason to change; it should do one job.

package main

import (
	"fmt"
)

type userManager struct{}

func (u *userManager) SaveUser(name string) {
	fmt.Println("Save user : ", name)
}

func (u *userManager) SendEmail(email, message string) {
	fmt.Println("Send email : ", email)
}

type EmailSender struct{}

func (e *EmailSender) SendEmail(email, message string) {
	fmt.Println("Send email : ", email)
}

func main() {
	userManager := &userManager{}
	userManager.SaveUser("Shiva")

	EmailSender := &EmailSender{}
	EmailSender.SendEmail("1234@gmail.com", "Welcome!")
}

// // This is bad typo because we need to make functions like that only work for one function
// type userManager struct{}

// func (u *userManager) SaveUser(name string) {
// 	fmt.Println("Save user : ", name)
// }

// func (u *userManager) SendEmail(email, message string) {
// 	fmt.Println("Send email : ", email)
// }

// type EmailSender struct{}

// func (e *EmailSender) Send(email, message string) {
// 	fmt.Println("Send email : ", email)
// }

// func main() {
// 	userMgr := &userManager{}
// 	userMgr.SaveUser("Shiva")

// 	emailSender := &EmailSender{}
// 	emailSender.Send("Shiva@gmail.com", "Welcome!")
// }

// package main

// import (
// 	"fmt"
// )

// // Bad : user manager does two things user data and email sending
// type userManager struct{}

// func (u *userManager) SaveUser(name string) {
// 	fmt.Println("user saved : ", name)
// }

// func (u *userManager) SendEmail(email, message string) {
// 	fmt.Println("Email sent to : ", email)
// }

// // Good seperate Email sender
// type EmailSender struct{}

// func (e *EmailSender) Send(email, message string) {
// 	fmt.Println("Email sent to : ", email)
// }

// func main() {
// 	userMgr := &userManager{}
// 	userMgr.SaveUser("Alice")

// 	emailSender := &EmailSender{}
// 	emailSender.Send("alice@gmail.com", "Welcome!")
// }
