package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	srv := http.Server{
		Handler: router,
		Addr:    ":8888",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		hostname, _ := os.Hostname()
		str := "You've hit " + hostname + "\n"
		str2 := "I've got a secret now!\n"
		str3 := "checkout my cool secret: platform_secrets_are_awesome=" + os.Getenv("platform_secrets_are_awesome") + "\n"
		str += str2 + str3
		w.Write([]byte(str))
	})

	log.Fatal(srv.ListenAndServe())
}
