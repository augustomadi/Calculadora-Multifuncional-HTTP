package main

import ("log"
	"net/http")

func StartServer() {
	s := &http.Server{

		Addr: "localhost:8080",
		Handler: AndreiHandler{},
		
	}
	log.Fatal(s.ListenAndServe())
}

