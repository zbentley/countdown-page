package main

//go:generate go run gen.go

import (
		"log"
		"net/http"
		"fmt"
		"time"
		"flag"
)

var index int

func StatusHandler(outstream http.ResponseWriter, r *http.Request) {
	var (
		len int = len(CountdownEntries)
		current CountdownEntry = CountdownEntries[index]
		now int64 = time.Now().Unix()
		render CountdownEntry
	)

	// This solves atomicity issues with doing index++
	if now > current.Time && index < len - 1 {
		index++
		current = CountdownEntries[index]
	}

	if index >= len - 1 {
		render = CountdownEntry{
			current.Text,
			-1, // Nothing coming ever again
		}
	} else {
		render = CountdownEntry{
			CountdownEntries[index - 1].Text,
			current.Time * 1000,
		}
	}

	outstream.WriteHeader(http.StatusOK)
	PageTemplate.Execute(outstream, render)
	log.Printf("Rendered response: %s", render)
}

func main() {
	var p int
	flag.IntVar(&p, "port", 8080, "Port on which to serve")
	flag.Parse()
	port := fmt.Sprintf(":%d", p)
	fmt.Printf("Starting on %s...\n", port)
	http.HandleFunc("/", StatusHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}