package main

import (
	"fmt"
	"log"
	"os"
	app "prjctr/md/11_elastic"
	"text/scanner"
)

// create index of words in ES <- think about analyzers
// load words from file to index
// use bulk api to load words


func main() {
	w, err := app.NewElasticsearchWrapper()
	if err != nil {
		panic(fmt.Errorf("error creating elasticsearch wrapper: %w", err))
	}

	indexName := os.Getenv("ELASTICSEARCH_TARGET_INDEX")

	if exists, _ := w.IndexExists(indexName); !exists {
		log.Println("index does not exist, creating a new one")
		err = w.CreateAutocompleteIndex(indexName)
		if err != nil {
			panic(fmt.Errorf("error creating index: %w", err))
		}
	} else {
		log.Println("index exists, continue processing")
	}

	const fileName = "words_alpha.txt"

	var s scanner.Scanner
	f, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Errorf("error reading file: %w", err))
	}
	
	defer f.Close()

	s.Init(f)
	
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		log.Println("starting iteration of the loop")

		if s.Position.Line > 10 {
			log.Println("breaking from the loop")
			break
		}
		
		// write to elastic
		err := w.IndexWord(indexName, s.TokenText())
		if err != nil {
			panic(fmt.Errorf("error indexing word: %w", err))
		}
		log.Printf("a word %s has been written to the index %s\n", s.TokenText(), indexName)
	}
}
