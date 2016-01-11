package main

import "net/http"

func runServer(port string) {
	http.HandleFunc("/checkGaleraStatus", checkGaleraStatus)
	http.ListenAndServe(":8080", nil)
}
