package main

import (
	"booking-app-golang/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0) //slice
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} //Waits for the launched go routine to finish
func main() {
	fmt.Printf("conferenceName is of type %T, and remainingTickets is of %T\n", conferenceName, remainingTickets)
	greetUsers()
	//var bookings = [50] string{} //declaring array

	for {
		//call getUserInput
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNo := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNo {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)                                               //tells the main thread to wait for 1 more threads execution to complete
			go sendTickets(userTickets, firstName, lastName, email) //go keyword for concurrency go routines
			//call function printFirstNames
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			//var noMoreTicketsLeft bool = remainingTickets == 0
			if remainingTickets == 0 {
				fmt.Println("No more tickets left. Come back next year :(")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email you entered is not valid.")
			}
			if !isValidTicketNo {
				fmt.Println("no of tickets entered is not valid.")
			}
		}

	}
	wg.Wait() // makes the main thread wait before execution of all the added threads is completed
}

func greetUsers() {
	fmt.Println("Welcome to our ", conferenceName, " booking application :)")
	fmt.Printf("We have total of %v tickets and %v remaining. Hurry!! \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend !!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { //using underscore we can ingore variables dont need to use
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your mail id")
	fmt.Scan(&email)
	fmt.Println("Enter no of tickets to be booked:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map
	// var userData = make(map[string]string) //make is a built-in method to declare an empty map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["NoOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v ticket remaining for %v.\n", remainingTickets, conferenceName)
}

// generating and sending ticket task runs in background when we place 'go' keyword in front of method call - runs in multithreaded env
func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#############################")
	fmt.Printf("Sending Ticket: \n %v \n to email %v\n", ticket, email)
	fmt.Println("#############################")
	wg.Done() //decrements the waitgroup counter by 1 and indicates that the go routine is finished
}
