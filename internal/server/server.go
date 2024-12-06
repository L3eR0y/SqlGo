package server

import (
	"fmt"
	"internal/handlers"
	"net/http"
)

func Run(port string) {
	fmt.Println("Server listening at " + port + " port...")

	http.HandleFunc("/", handlers.RootHandler)
	http.ListenAndServe(port, nil)

}
