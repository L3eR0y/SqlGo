package sqlsrv

import (
	"fmt"
	"net/http"
)

func Run(port string) {
	fmt.Println("Server listening at " + port + " port...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Wolrd")
	})

	http.ListenAndServe(port, nil)
}
