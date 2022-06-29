package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://172.19.0.2:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		fmt.Println("Error getting response: %s", err)
	}

	defer res.Body.Close()

	ctx := context.Background()

	req := esapi.ExistsRequest{
		Index:      "my-index-000001",
		DocumentID: "dQwjsYEB8fcjMMM-5Gqe",
	}
	resp, err := req.Do(ctx, es)
	if err != nil {
		log.Fatalf("ExistsRequest ERROR: %s", err)
	}
	status := resp.StatusCode
	if status == 200 {
		fmt.Println("Exist")
		os.Exit(1)
	} else if status == 404 {
		fmt.Println("Not found")
	}
}
