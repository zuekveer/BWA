// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define a simple handler for the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "simple thoughts")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes writtes: %d", n))
	})

	// Start the server on port 8088
	port := "8088"
	log.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
