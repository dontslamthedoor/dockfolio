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

	body, err := ioutil.ReadAll(res.Body)
    	if err != nil {
        	log.Fatalf("Failed to read response body: %v", err)
    	}
    		fmt.Printf("response body: %s\n", body)

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
