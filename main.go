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

	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	//fmt.Println("Welcome to", conferenceName, "booking application")
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	//fmt.Println("Get your tickets here to attend")
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

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

		fmt.Println("Enter number of tickets:")
		_, err := fmt.Scan(&userTickets)
		fmt.Println(err)
		/*for err != nil { // for learning, do not use it!!!
			var userTickets uint = 999
			fmt.Println("Enter an integer for valid ticket number!!!")
			fmt.Println("Enter number of tickets: ")
			_, err = fmt.Scan(&userTickets)
		}*/
		/*fmt.Println(userTickets)*/
		/*enter_ticket: /// only for learning, do not use it!!!
		fmt.Println("Enter number of tickets: ")
		_, err := fmt.Scan(&userTickets)
		if err != nil {
			fmt.Println("Enter an integer for valid ticket number!!!")
			goto enter_ticket /// only for learning, do not use it!!!
		}*/
		// fmt.Println(err) check if the given input is integer
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets //&& err != nil

		isValidInfo := isValidName && isValidEmail && isValidTicketNumber

		if isValidInfo {
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
			if !isValidName {
				fmt.Println("First name or last name you entered is too short!")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
				fmt.Println(userTickets)
			}

			//fmt.Println("Your input data is invalid, try again")  /// avoid generic user feedbacks!!!
			//fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	//execute code for new york tickets
	// case "Singapore":
	// 	//execute code for singapore tickets
	// case "London":
	// 	//execute code for london tickets
	// case "Ankara","Istanbul":
	//	//execute code for Ankara or Istanbul tickets
	// default:
	// 	fmt.Println("City is invalid")
	// }
}
func greetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func main() {
	bookingApplication()
}
