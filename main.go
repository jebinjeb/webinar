package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var remainingTickets int32 = 50

// Register handles webinar registration and returns the remaining tickets
func Register() int32 {
	return atomic.AddInt32(&remainingTickets, -1)
}

// Handler function to handle webinar registrations
func registrationHandler(w http.ResponseWriter, r *http.Request) {
	remaining := Register()

	message := fmt.Sprintf("Registration successful! Remaining tickets: %d", remaining)
	fmt.Fprintln(w, message)
}

func main() {
	// Set up a registration handler
	http.HandleFunc("/register", registrationHandler)

	// Start the web server on port 8080
	fmt.Println("Web server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
