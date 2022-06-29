package main

import (
        "github.com/elastic/go-elasticsearch/v7"
        "fmt"
        "log"
        "github.com/elastic/go-elasticsearch/v7/esutil"
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

        query := map[string]interface{}{
                "query": map[string]interface{}{
                        "match": map[string]interface{}{
                                "_id": "aAwjsYEB8fcjMMM-fWrV",
                        },
                },
        }

        res, err = es.Search(
                es.Search.WithIndex("my-index-000001"),
                es.Search.WithBody(esutil.NewJSONReader(&query)),
                es.Search.WithPretty(),
        )
        if err != nil {
                log.Fatalf("Error getting response: %s", err)
        }

        log.Println(res)
}
