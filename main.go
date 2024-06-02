package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type UserData struct {
   userName string
   userTickets int
   userEmail string
}
var wq=sync.WaitGroup{}
func main(){
	var conferenceName ="Go Conference"
   const conferenceTickets = 50
    var remainingTickets = 50
   fmt.Println("Welcome to our",conferenceName," booking application")
   fmt.Println("We have total of ",conferenceTickets,"and ",remainingTickets,"are still available")
   fmt.Println("Get your tickets here to attend")

   var bookings=make([] UserData,0)

   var userName string
   var userEmail string
   var userTickets int
   

	fmt.Println("Enter your Name")
	fmt.Scan(&userName)

	fmt.Println("Enter your Email Address")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

		var userData = UserData{
			userName: userName,
			userEmail: userEmail,
			userTickets: userTickets,
		}
        
		isValidName := len(userName) >= 2
		isValidEmail := strings.Contains(userEmail,"@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets 
        
		if !isValidEmail {
			fmt.Println("invalid email id , enter once again")
		}
		if !isValidName {
			fmt.Println("invalid name , enter once again")
		}
		if !isValidTicketNumber {
			fmt.Println("sorry current available tickets:",remainingTickets)
		}
		if isValidEmail && isValidName && isValidTicketNumber{	
			remainingTickets=remainingTickets-userTickets

			bookings = append(bookings, userData)

			var firstNames = [] string{}
			for _, booking := range bookings{
				firstNames=append(firstNames, booking.userName)
			}
			fmt.Println(firstNames)
            
			wq.Add(1)
			go sendTicket(userName,userEmail,userTickets)
			if remainingTickets == 0 {
				fmt.Println("All conference tickets is booked out, better luck next time")
			}
	    }
		 
		wq.Wait()
}

func sendTicket(userName string ,userEmail string, userTickets int){
	time.Sleep(50*time.Second)
	var ticket= fmt.Sprintf("%v tickets for %v",userTickets,userName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket %v \n to email address %v \n",ticket,userEmail)
	fmt.Println("######################")
	wq.Done()
}