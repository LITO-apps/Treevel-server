package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	dump, err := httputil.DumpRequest(r, true)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(dump))

	// FIXME: Unhandled error
	fmt.Fprint(w, "hello world!\n")
}

func main()  {
	var httpServer http.Server
	// Port number
	httpServer.Addr = ":8080"
	// Rooting
	http.HandleFunc("/", handler)
	// Log
	log.Println("Server Start" + httpServer.Addr)
	// Listen
	log.Println(httpServer.ListenAndServe())
}
