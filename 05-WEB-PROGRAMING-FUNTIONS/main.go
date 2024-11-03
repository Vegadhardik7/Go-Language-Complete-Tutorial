package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// handler function | Signature: w http.ResponseWriter, r *http.Request
// func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(w, "Hello World")
// }

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// register the helloWorldHandler function to handle requests to the root URL ("/")
	http.HandleFunc("/", getHelloWorldHandler(infoLog))

	// start the web server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func getHelloWorldHandler(infoLogger *log.Logger) func(http.ResponseWriter, *http.Request) {

	// returning the handler code without creating and calling it explicitly
	return func(w http.ResponseWriter, r *http.Request) {
		// logging message using the injection infologger
		infoLogger.Printf("Requested hello world")

		// writing "Hello World" to the response writer
		fmt.Fprintf(w, "Hello World")
	}
}
