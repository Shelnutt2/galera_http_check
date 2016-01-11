package main

import "net/http"

func runServer(port string) {
	http.HandleFunc("/check-galera-status", checkGaleraStatus)
	http.ListenAndServe(":8080", nil)
}
