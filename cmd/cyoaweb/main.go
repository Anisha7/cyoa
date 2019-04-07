package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"story/cyoa"
)

func main() {
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
	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err) // again, print err instead
	}

	fmt.Printf("%v\n", story)
}
