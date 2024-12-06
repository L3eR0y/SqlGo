package server

import (
	"fmt"
	"internal/handlers"
	"net/http"
)

func Run(port string) {
	fmt.Printf("Server listening at %s", port)

	http.HandleFunc("/", handlers.RootHandler)
	http.ListenAndServe(port, nil)
}
