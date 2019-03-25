package main

import (
	"math/rand"
	"net/http"
	"time"
)

func getMain(w http.ResponseWriter, r *http.Request) {

	var quote string
	var content Content

	quotes, err := retrieveQuotes()
	if err != nil {
		printErr(err, "Error while trying to get quotes!")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - There was an internal server error!"))
	}
	if len(quotes) > 0 {
		rand.Seed(time.Now().Unix())
		quote = quotes[rand.Intn(len(quotes))].Text
	}

	content = Content{Timer: initTime.Unix(), QuoteText: quote}

	render.HTML(w, http.StatusOK, "main", content)
}

func postQuote(w http.ResponseWriter, r *http.Request) {
	quote := r.FormValue("quote")
	if quote != "" {
		_, err := insertQuote(quote)
		if err != nil {
			printErr(err, "Error while inserting quote!")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - There was an internal server error!"))
		}
	}

	http.Redirect(w, r, "http://localhost:5000/", 301)
}
