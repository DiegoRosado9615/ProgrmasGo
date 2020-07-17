package main

/*
Programa prueba en el cual montare un servidor en go
*/
import (
	"log"
	"net/http"
	"github.com/diego/servidor"
)

func main() {
	s=:server.New()
	log.Fatal(http.ListenAndServe(":8000", nil))

}
