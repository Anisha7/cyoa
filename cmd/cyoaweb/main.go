package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Anisha7/cyoa"
)

func main() {
	port := flag.Int("post", 3000, "the port to start the CYOA web application on")
	// option to provide filename, but uses gopher.json that we have by default
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	// open the file
	f, err := os.Open(*filename)
	if err != nil {
		panic(err) // not a good idea, print a logical error instead
	}

	// pass file to decoder
	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err) // again, print err instead
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h)) // return if error and passes it to string
}
