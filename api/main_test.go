package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainExecution(t *testing.T) {

	go main()

	resp, err := http.Get("http://localhost:9001/checkHealth")

	if err != nil {
		t.Fatalf("Cannot make get: %v\n", err)
	}

	bodySb, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatalf("Error reading body: %v\n", err)
	}

	body := string(bodySb)
	fmt.Printf("Body: %v\n", body)

	var decodedResponse interface{}
	fmt.Println(bodySb)
	err = json.Unmarshal(bodySb, &decodedResponse)

	if err != nil {
		t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
	}

	assert.Equal(t, map[string]interface{}{"status": "ok"}, decodedResponse,
		"Should return status:ok")
}
