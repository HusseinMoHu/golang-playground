package main

import (
	"booking-app/helper"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const conferenceName = "Go Conference"
const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings []string

func main() {
	greetingUser()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInputs(firstName, lastName, email, userTickets)

		if !isValidName {
			color.Red("Please enter a valid name")
			continue
		}

		if !isValidEmail {
			color.Red("Please enter a valid email")
			continue
		}

		if !isValidTickets {
			color.Red("Please enter a valid number of tickets")
			continue
		}

		if userTickets > remainingTickets {
			color.Red("We only have %v tickets, so you can't book %v tickets", remainingTickets, userTickets)
			continue
		}

		bookingTickets(userTickets, firstName, lastName, email)

		firstNames := getBookingsFirstNames()
		color.Blue("Here are all of our bookings %v", firstNames)
	}

	color.Green("All tickets are booked")
}

func greetingUser() {
	fmt.Printf("-- Welcome to %v booking system --\n", color.BlueString(conferenceName))
	fmt.Printf(
		"We have total of %v tickets and %v are still available\n",
		color.BlueString(fmt.Sprint(conferenceTickets)),
		color.GreenString(fmt.Sprint(remainingTickets)),
	)

	color.New(color.Bold, color.FgGreen).Println("Get your tickets here to attend")
}

func getBookingsFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookingTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf(
		"Hi %v %v, you have successfully booked %v tickets. %v tickets are still available\n",
		color.BlueString(firstName),
		color.BlueString(lastName),
		color.GreenString(fmt.Sprint(userTickets)),
		color.GreenString(fmt.Sprint(remainingTickets)),
	)

	color.Green("Thank you for booking with us, You will receive an email shortly at %v\n", color.CyanString(email))
}
