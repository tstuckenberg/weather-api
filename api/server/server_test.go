package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TestMain test Router configuration
func TestMain(t *testing.T) {
	// setup the test server
	// router := NewRouter()
	// testServer = httptest.NewServer(router)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	}))

	defer testServer.Close()

	res, err := http.Get(testServer.URL)

	if err != nil {
		log.Fatal(err)
	}

	hello, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", hello)

	// run tests
	// os.Exit(m.Run())
}
