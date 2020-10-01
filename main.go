package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	timeout, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/hello", hello(time.Duration(timeout)*time.Second))
	log.Fatal(http.ListenAndServe(":8090", nil))
}
