package main


// Creating a two page website application
import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// In other for a function to respond to a request from a web browser;
// It has to handle two parameters;
// A response writer called (w http.ResponseWriter, r *http.Request)
// and a request r *http.Request


// Home is the  home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	// %d is the placeholder for integer
	// _ means to ignore
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2  %d", sum))

}

// addValues adds rwo integers and returns sum
func addValues(x, y int) int {
	
	return x + y
}


// main is the main application function
func main() {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	
	
	
	
	
	
	
	
	// http.HandleFunc to tell the server which function to call
	// to handle a request to the server

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// Fprintf formats according to a format specifier and writes to w. 
		// And it returns the number of bytes written and any write error encountered
		// n int for(number of bytes written)
		// err error for(error encoutered)

		// n, err := fmt.Fprintf(w, "Hello, world!")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Sprintf() allows you to take different data types
		// and return them as a string

	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
		
	// }) 

	// http.ListenAndServe function to start the server 
	// and tell it to listen for new HTTP requests and 
	// then serve them using the handler functions you set up.

    // %s is a placeholder for string
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_= http.ListenAndServe(portNumber, nil)
}