package serviceutils

import (
	"log"
	"testing"
)

func TestStatusCodeParser(t *testing.T) {
	URL := "https://localhost:8443/ping"
	if err := statusCodeParser(URL, 502); err == nil {
		log.Fatalf("%v", err)
	}
	if err := statusCodeParser(URL, 200); err != nil {
		t.Fatalf("%v", err)
	}
}
