package main

import (
	"cpea_monthly_usage/processing"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// server to start the app
	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port", port)

	fmt.Println("Started App.......")
	// Starting Thraed After every 10 minutes
	processing.CheckMonthlyData()

	//If Error Occurs End the processing Of Application
	go processing.CheckErrorAndEndApplication()
	http.ListenAndServe(":"+port, nil)

}

func handler(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello, World!") }
