package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start cyoa web application on")
	filename := flag.String("file", "gopher.json", "the json file with the CYOA story")
	flag.Parse()

	// fmt.Printf("using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v", story)

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
