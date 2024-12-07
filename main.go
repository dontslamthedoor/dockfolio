package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.example.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	println(res.Body)

	fmt.Printf("response body: %v", res.Body)

	http.HandleFunc("/helloKeith", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, Keith! this is net/http\n")

	})

	log.Println("starting server....")
	n, _ := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)

	}
	go func() {
		log.Fatal(http.Serve(n, nil))

	}()

}
