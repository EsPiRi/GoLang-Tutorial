package main

import (
	"fmt"
	"strings"
)

func bookingApplication() {
	//tempConferenceName := "Go Conference"
	//var conferenceName string = "Go Conference"

	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	//fmt.Println("Welcome to", conferenceName, "booking application")
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	//fmt.Println("Get your tickets here to attend")

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings = [50]string{} 	/////// array
	// var bookings =[]string{}		/////// slice
	// var bookings []string		/////// another slice

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets -= userTickets
			//bookings[0] = firstName + " " + lastName
			//bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			/*fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)
			fmt.Printf("Array length: %v\n", len(bookings))*/

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings { // use _ for blank identifier that we dont use in the next lines of code
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			//fmt.Printf("These are all our bookings: %v\n", bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("No tickets left for %v. Thank you for your interest. See you next year!", conferenceName)
				break
			}

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

}

func main() {
	bookingApplication()
}
