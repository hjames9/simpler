package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type simpleHandler struct{}

func (handler simpleHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	pid := os.Getpid()
	hostname, _ := os.Hostname()
	currentTime := time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	var arguments []string

	if len(os.Args[1:]) > 0 {
		arguments = os.Args[1:]
	}

	data := struct {
		Path      string   `json:"path"`
		Method    string   `json:"method"`
		Pid       int      `json:"pid"`
		Hostname  string   `json:"hostname"`
		Timestamp int64    `json:"timestamp"`
		Arguments []string `json:"arguments,omitempty"`
	}{
		request.URL.Path,
		request.Method,
		pid,
		hostname,
		currentTime,
		arguments,
	}

	response.Header().Set("Content-Type", "application/json")
	str, _ := json.Marshal(data)
	fmt.Fprintf(response, "%s", str)
}

func main() {
	log.Println("Starting simpler service")

	if len(os.Args[1:]) > 0 {
		log.Println(fmt.Sprintf("CLI arguments used: %v", os.Args[1:]))
	} else {
		log.Println("No CLI arguments specified")
	}

	server := &http.Server{
		Addr:           ":8000",
		Handler:        simpleHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	//Signal handler
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	signal.Notify(signals, syscall.SIGTERM)

	sig := <-signals
	switch sig {
	case os.Interrupt:
		fallthrough
	case syscall.SIGTERM:
		log.Println("Cleanly shutting down...")
		server.Close()
		os.Exit(0)
	default:
		log.Println("Uncleanly shutting down...")
		server.Close()
		os.Exit(1)
	}
}
