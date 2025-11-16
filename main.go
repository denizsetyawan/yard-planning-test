package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	InitDB()

	http.HandleFunc("/suggestion", SuggestionHandler)
	http.HandleFunc("/placement", PlacementHandler)
	http.HandleFunc("/pickup", PickupHandler)

	port := "8080"
	fmt.Println("Server is running at http://localhost:" + port)
	fmt.Println("Press CTRL+C to stop the server.")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
