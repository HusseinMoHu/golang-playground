package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking system\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	// ask for username

	userName = "Tom"
	fmt.Printf(userName)
}