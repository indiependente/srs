package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/context"
)

func hello2(timeout time.Duration) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(req.Context(), timeout)
		defer cancel()

		fmt.Println("server: hello handler started")
		defer fmt.Println("server: hello handler ended")

		values := req.URL.Query()
		duration := values.Get("t")
		queryDuration, err := strconv.Atoi(duration)
		if err != nil {
			log.Printf("error %q\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// worker
		go func() {
			fmt.Println("SLOWSQLQUERY...")
			time.Sleep(time.Duration(queryDuration) * time.Second)
			fmt.Println("done!")
			cancel()
		}()

		// monitor the Done channel of the context
		<-ctx.Done()
		err = ctx.Err()
		fmt.Println("server:", err)
		switch {
		case errors.Is(err, context.Canceled):
			fmt.Fprintf(w, "all good\n")
		default:
			log.Printf("error %q\n", err)
			log.Printf("forcing lock.Release()\n")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
