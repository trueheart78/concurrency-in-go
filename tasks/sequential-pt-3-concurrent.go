package main

import (
	"fmt"
	"sync"
)

func makeHotelReservation(wg *sync.WaitGroup) {
	fmt.Println("Done making hotel reservation.")
	wg.Done()
}
func bookFlightTickets(wg *sync.WaitGroup) {
	fmt.Println("Done booking flight tickets.")
	wg.Done()
}
func orderADress(wg *sync.WaitGroup) {
	fmt.Println("Done ordering a dress.")
	wg.Done()
}
func payCreditCardBills(wg *sync.WaitGroup) {
	fmt.Println("Done paying credit card bills.")
	wg.Done()
}

func writeEmail(wg *sync.WaitGroup) {
	fmt.Println("Wrote a 1/3rd of the email.")
	go continueWritingEmail1(wg)
}

func continueWritingEmail1(wg *sync.WaitGroup) {
	fmt.Println("Wrote a 2/3rds of the email.")
	go continueWritingEmail2(wg)
}

func continueWritingEmail2(wg *sync.WaitGroup) {
	fmt.Println("Done writing the email.")
	wg.Done()
}

func listenToAudioBook(wg *sync.WaitGroup) {
	fmt.Println("Listened to 1/2 of the audio book.")
	go continueListeningToAudioBook(wg)
}

func continueListeningToAudioBook(wg *sync.WaitGroup) {
	fmt.Println("Done listening to audio book.")
	wg.Done()
}

var listOfTasks = []func(*sync.WaitGroup){
	makeHotelReservation, bookFlightTickets, orderADress,
	payCreditCardBills, writeEmail, listenToAudioBook,
}

func main() {
	var waitGroup sync.WaitGroup    // new WaitGroup
	waitGroup.Add(len(listOfTasks)) // add number of tasks to wait on

	for _, task := range listOfTasks {
		go task(&waitGroup) // pass the WaitGroup via reference, and call via goroutine
	}

	waitGroup.Wait() // wait for completion
}
