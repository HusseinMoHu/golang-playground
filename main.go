package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("-- Welcome to %v booking system --\n", color.BlueString(conferenceName))
	fmt.Printf(
		"We have total of %v tickets and %v are still available\n",
		color.BlueString(fmt.Sprint(conferenceTickets)),
		color.GreenString(fmt.Sprint(remainingTickets)),
	)

	color.New(color.Bold, color.FgGreen).Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Please enter number of tickets: ")
	fmt.Scan(&userTickets)

	if userTickets > remainingTickets {
		color.Red("Sorry, we don't have enough tickets")
		return
	}

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first element %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf(
		"Hi %v %v, you have successfully booked %v tickets. %v tickets are still available\n",
		color.BlueString(firstName),
		color.BlueString(lastName),
		color.GreenString(fmt.Sprint(userTickets)),
		color.GreenString(fmt.Sprint(remainingTickets)),
	)

	color.Green("Thank you for booking with us, You will receive an email shortly at %v\n", color.CyanString(email))

}
