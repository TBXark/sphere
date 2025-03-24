package urlhandler

import (
	"log"
	"testing"
)

const (
	testPublicURL = "http://localhost:8080"
	testURI       = "http://localhost:8080/test?width=100"
)

func TestHandler_ExtractKeyFromURLWithMode(t *testing.T) {
	handler, _ := NewHandler(testPublicURL)
	key, err := handler.ExtractKeyFromURLWithMode(testURI, true)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(key)
}

func TestHandler_GenerateURL(t *testing.T) {
	handler, _ := NewHandler(testPublicURL)
	url := handler.GenerateURL("test")
	log.Print(url)
}
