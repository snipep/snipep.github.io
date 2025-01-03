package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlFilePath := "script/index.html" 
		htmlContent, err := os.ReadFile(htmlFilePath)
		if err != nil {
			http.Error(w, "Error loading HTML file", http.StatusInternalServerError)
			fmt.Printf("Error reading HTML file: %s\n", err)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(htmlContent)
	})

	port := ":8080"
	fmt.Printf("Web server running on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

