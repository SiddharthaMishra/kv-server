package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	a := initApp()

	srv := &http.Server{
		Addr:    ":8000",
		Handler: a.router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		fmt.Println("Server Started")
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	os.Exit(0)

}
