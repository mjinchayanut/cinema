package main

import (
	"fmt"

	"github.com/mjinchayanut/cinema/movie"
	"github.com/mjinchayanut/cinema/ticket"
)

func main() {
	movieName := movie.FindName("1")
	fmt.Println(movieName)

	movie.Review(movieName, 5.5)
	ticket.BuyTicket(movieName)
}
