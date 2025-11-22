package main

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Client struct{}

func (c *Client) InsertLightningConectorIntoComputer(com Computer) {
	fmt.Println("Client inserts lightning conector into computer")
	com.InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine")
}

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}

type WindowAdapter struct {
	windowMachine *Windows
}

func (w *WindowAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter covnverts lightning signal to USB")
	w.windowMachine.insertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConectorIntoComputer(windowsMachineAdapter)
}
