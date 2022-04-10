package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sohhamm/string-svc/transports"
)

func main() {

	http.Handle("/uppercase", transports.UppercaseHandler)
	http.Handle("/count", transports.CountHandler)

	fmt.Println("server starter..... 🔥")

	log.Fatal(http.ListenAndServe(":9000", nil))

}
