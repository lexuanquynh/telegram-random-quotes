package main

import (
	"encoding/json"
	"github.com/go-co-op/gocron"
	"github.com/hashicorp/go-hclog"
	"log"
	"net/http"
	"time"
)

// main func
func main() {

	logger := NewLogger()

	s := gocron.NewScheduler(time.UTC)
	// Reset limit send mail for users.
	_, err := s.Every(1).Hour().Do(func() {
		logger.Info("Call url .")
		// get message from getQuote function
		message := getQuote()
		callURL(message)
	})

	if err != nil {
		logger.Error("Error scheduling limit data", "error", err)
		return
	}

	//s.StartAsync()
	s.StartBlocking()
}

// call url
func callURL(message string) {
	// get http with message param
	if message != "" {
		resp, err := http.Get("https://api.telegram.org/bot5:AAEO/sendMessage?chat_id=-100ABC&text=" + message + "&parse_mode=markdown")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
	}
}

// get a quotes random
func getQuote() string {
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var quotes []Quote

	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		log.Fatalln(err)
	}
	if len(quotes) > 0 {
		return quotes[0].Q + " - " + quotes[0].A
	}
	return ""
}

// Quote struct
type Quote struct {
	Q string `json:"q"`
	A string `json:"a"`
}

// NewLogger returns a new logger instance
func NewLogger() hclog.Logger {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "user-auth-service",
		Level: hclog.LevelFromString("DEBUG"),
	})

	return logger
}
