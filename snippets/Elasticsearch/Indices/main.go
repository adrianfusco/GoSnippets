package main

import (
	"context"
	"fmt"

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

	resp, err := esapi.CatIndicesRequest{Format: "json"}.Do(context.Background(), es)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp.String())
}
