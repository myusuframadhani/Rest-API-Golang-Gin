package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	alamat := "localhost:8080"
	fmt.Printf("Server berjalan pada port %s\n", alamat)
	err := http.ListenAndServe(alamat, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	pesan := "Welcome"
	w.Write([]byte(pesan))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	pesan := "Hello World!"
	w.Write([]byte(pesan))
}
