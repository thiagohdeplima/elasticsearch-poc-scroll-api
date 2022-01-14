package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func SaveInElastic(person Person) error {
	// convert Person{} to json
	bytes, err := json.Marshal(&person)
	if err != nil {
		return err
	}

	// connect with elastic
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("creating client: `%s`", err)
		return err
	}

	req := esapi.IndexRequest{
		Index:      "person",
		DocumentID: person.IDNumber,
		Body:       strings.NewReader(string(bytes)),
		Refresh:    "true",
	}

	log.Printf("sending person %s", string(bytes))

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	if res.IsError() {
		log.Printf("[%s|%+v] Error indexing document ID=%s", res.Status(), res, person.IDNumber)
	}

	return nil
}
