package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Quotable booting...")

	server := NewServer()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-sig:
				server.Close()
				os.Exit(0)
			}
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Launching server on port", port)
	http.ListenAndServe(":"+port, server.Router)
}
