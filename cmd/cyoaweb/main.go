package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/b-codessoft/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Println("Starting the server on port: %d\n", *port)

}
