package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	num1 := r.URL.Query().Get("num1")
	num2 := r.URL.Query().Get("num2")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "status code 200:ok")
	fmt.Fprintln(w, "num1:", num1)
	fmt.Fprintln(w, "num2:", num2)
	fmt.Fprintln(w, num1, "+", num2, "=", (int)(num1)+(int)(num2))
}

func main() {
	http.HandleFunc("/api/vi/health", healthHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
