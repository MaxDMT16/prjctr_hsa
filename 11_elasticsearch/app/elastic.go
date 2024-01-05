package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

type elasticsearchWrapper struct {
	c *elasticsearch.Client
}

func NewElasticsearchWrapper() (*elasticsearchWrapper, error) {
	c, err := newElasticsearchClient()
	if err != nil {
		return nil, fmt.Errorf("creating elasticsearch client: %w", err)
	}

	return &elasticsearchWrapper{c: c}, nil
}

func newElasticsearchClient() (*elasticsearch.Client, error) {
	rawUrl := os.Getenv("ELASTICSEARCH_URL")
	urls := strings.Split(rawUrl, ",")

	fmt.Println("urls:")
	fmt.Println(urls)

	password := os.Getenv("ELASTIC_PASSWORD")
	
	config := elasticsearch.Config{
		Addresses: urls,
		Username:  "elastic",
		Password:  password,
	}

	return elasticsearch.NewClient(config)
}

func (e *elasticsearchWrapper) CreateAutocompleteIndex(indexName string) error {
	body := `{
		"settings": {
		  "number_of_replicas": 1,
		  "number_of_shards": 3,
		  "analysis": {
			"analyzer": {
			  "autocomplete": {
				"tokenizer": "autocomplete",
				"filter": [
				  "lowercase"
				]
			  },
			  "autocomplete_search": {
				"tokenizer": "lowercase"
			  }
			},
			"tokenizer": {
			  "autocomplete": {
				"type": "edge_ngram",
				"min_gram": 2,
				"max_gram": 10,
				"token_chars": [
				  "letter"
				]
			  }
			}
		  },
		  "refresh_interval": "1s"
		},
		"mappings": {
		  "dynamic": false,
		  "properties": {
			"word": {
			  "type": "text",
			  "analyzer": "english"
			}
		  }
		}
	  }`

	res, err := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(body),
	}.Do(context.Background(), e.c)
	if err != nil {
		return fmt.Errorf("creating index: %w", err)
	}

	defer res.Body.Close()

	return nil
}

func (e *elasticsearchWrapper) IndexExists(name string) (bool, error) {
	resp, err := esapi.IndicesExistsRequest{
		Index: []string{name},
	}.Do(context.Background(), e.c)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}

func (e *elasticsearchWrapper) IndexWord(indexName string, word string) error {
	body := `{
		"word": "%s"
	}`

	body = fmt.Sprintf(body, word)

	res, err := e.c.Index(indexName, strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("indexing word: %w", err)
	}

	defer res.Body.Close()

	return nil
}
