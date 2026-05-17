package main

import "fmt"

func main() {
	conferenceName := "DK Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your ticket here to attend")

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receeive a confirmation email at %v.", firstName,lastName,userTickets,email)


	
}
