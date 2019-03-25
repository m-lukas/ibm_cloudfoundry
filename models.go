package main

//The Quote object contains an unique identifier and it's data.
type Quote struct {
	ID   int
	Text string
}

//The Content object is used to fill the HTML template.
type Content struct {
	Timer     int64
	QuoteText string
}

//The Health object is used in the CF Health Check.
type Health struct {
	Status string `json:"status"`
}
