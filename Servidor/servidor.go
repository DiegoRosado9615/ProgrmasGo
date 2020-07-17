package main

/*
Programa prueba en el cual montare un servidor en go
*/
import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", nil))

}
