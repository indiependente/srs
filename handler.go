package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

func hello(timeout time.Duration) func(w http.ResponseWriter, req *http.Request) {
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
		var eg errgroup.Group

		// monitor the Done channel of the context
		eg.Go(func() error {
			<-ctx.Done()
			err := ctx.Err()
			fmt.Println("server:", err)
			return err

		})

		// worker
		eg.Go(func() error {
			fmt.Println("SLOWSQLQUERY...")
			time.Sleep(time.Duration(queryDuration) * time.Second)
			fmt.Println("done!")
			cancel()
			return nil
		})

		// wait for goroutines
		if err := eg.Wait(); err != nil {
			// the deadline was exceeded => release the lock
			if errors.Is(err, context.DeadlineExceeded) {
				log.Printf("error %q\n", err)
				log.Printf("forcing lock.Release()\n")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		fmt.Fprintf(w, "all good\n")
	}
}
