package cmd


import (
	// "fmt"
	"os"
	"io"
	"log"
	"net/http"
)

// using global variable so itll feature in the second function as well
// var body []byte

// Load up the page and possibly copy the content to be saved
func Copy(link string) ([]byte, error) {
	// link := os.Args[1]
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: &s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

// Pasting content of URL into the file created 
func Paste(body []byte, filename string) string {
	err := os.WriteFile((filename+".txt"), body, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return ("File successfully downloaded")
}