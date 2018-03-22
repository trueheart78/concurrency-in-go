package main

import "fmt"

func makeHotelReservation() {
	fmt.Println("Done making hotel reservation.")
}
func bookFlightTickets() {
	fmt.Println("Done booking flight tickets.")
}
func orderADress() {
	fmt.Println("Done ordering a dress.")
}
func payCreditCardBills() {
	fmt.Println("Done paying credit card bills.")
}

func writeEmail() {
	fmt.Println("Wrote a 1/3rd of the email.")
	continueWritingEmail1()
}

func continueWritingEmail1() {
	fmt.Println("Wrote a 2/3rds of the email.")
	continueWritingEmail2()
}

func continueWritingEmail2() {
	fmt.Println("Done writing the email.")
}

func listenToAudioBook() {
	fmt.Println("Listened to 1/2 of the audio book.")
	continueListeningToAudioBook()
}

func continueListeningToAudioBook() {
	fmt.Println("Done listening to audio book.")
}

var listOfTasks = []func(){
	makeHotelReservation, bookFlightTickets, orderADress,
	payCreditCardBills, writeEmail, listenToAudioBook,
}

func main() {
	for _, task := range listOfTasks {
		task()
	}
}
