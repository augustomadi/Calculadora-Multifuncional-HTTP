package main 

import("net/http"
	"log"
)

func StartServer(){
	s := &http.Server{
		Addr:           "localhost:8080",
		Handler:        Handler{},
		
	}
	log.Fatal(s.ListenAndServe())
}