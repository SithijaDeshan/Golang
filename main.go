package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "DK Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email : ")
		fmt.Scan(&email)

		fmt.Println("Please enter how many tickets do you want : ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isEmailValid := strings.Contains(email, "@")
		isValidTecketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isEmailValid && isValidTecketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receeive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName{
				fmt.Printf("First name or last name you entered too short.\n")
			} 
			if !isEmailValid{
				fmt.Printf("Email you entered is not correct.\n")
			} 
			if !isValidTecketNumber{
				fmt.Printf("Number of tickets you entered is invalid\n")
			} 
		}

	}

}
