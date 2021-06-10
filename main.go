package main

import (
	"fmt"
	"net/http"

	"github.com/gustavocarmoh/routes"
)

func main() {
	routes.Routes()

	fmt.Println("Server is running in port 8000.")
	http.ListenAndServe(":8000", nil)
}

