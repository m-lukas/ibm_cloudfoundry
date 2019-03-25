package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/thedevsaddam/renderer"
)

var render *renderer.Render
var initTime time.Time

const apiPath = "/api"

func init() {
	//prepare html rendering
	options := renderer.Options{
		ParseGlobPattern: "./res/*.html",
	}
	render = renderer.New(options)
}

func main() {
	initTime = time.Now() //get initial timestamp

	r := mux.NewRouter()
	r.HandleFunc("/", getMain)
	r.HandleFunc("/health", healthCheck).Methods("GET")
	r.HandleFunc(apiPath+"/quote", postQuote).Methods("POST")

	//cors configuration
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "*"},
		Debug:            false,
	})

	handler := cors.Handler(r)

	//coonfigurate server
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler, //pass mux handler
	}
	log.Println("Configured server")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Started server")

	c := make(chan os.Signal, 1) //gracefully stop on CTRL + C
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	log.Println("Server Status: offline")
	os.Exit(0)
}

func printErr(err error, message string) {
	if err != nil {
		log.Println(message)
		log.Println(err)
	}
}
