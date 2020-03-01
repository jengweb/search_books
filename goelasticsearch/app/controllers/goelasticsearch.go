package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticsearchContextInterface interface {
	JSON(code int, obj interface{})
	GetPostForm(key string) (string, bool)
	GetQuery(key string) (string, bool)
}

type GoElasticsearchInterface interface {
	GetGoElasticsearch(contexts ElasticsearchContextInterface, request *http.Request)
	PostGoElasticsearch(contexts ElasticsearchContextInterface, request *http.Request)
	UpdateGoElasticsearch(contexts ElasticsearchContextInterface, request *http.Request)
}

type GoElasticsearchController struct{}

func (controller *GoElasticsearchController) GetGoElasticsearch(contexts ElasticsearchContextInterface, request *http.Request) {
	log.SetFlags(0)
	querys, _ := contexts.GetQuery("query")
	var (
		results map[string]interface{}
	)
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://es01:9200",
			"http://es02:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", results["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  querys,
				"fields": []string{"title", "isbn", "author_name"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("store"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(results["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(results["took"].(float64)),
	)
	log.Println(strings.Repeat("=", 37))
	for _, hit := range results["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Println(hit.(map[string]interface{})["_source"])
	}
	contexts.JSON(200, results["hits"].(map[string]interface{})["hits"].([]interface{}))
}

func (controller *GoElasticsearchController) PostGoElasticsearch(contexts ElasticsearchContextInterface, request *http.Request) {
	dataInput, _ := contexts.GetPostForm("data")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://es01:9200",
			"http://es02:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// Build the request body.
	var b strings.Builder
	b.WriteString(dataInput)
	req := esapi.IndexRequest{
		Index: "store",
		// DocumentID: strconv.Itoa(i + 1),
		Body:    strings.NewReader(b.String()),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Printf("[%s] Error indexing", res.Status())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	contexts.JSON(200, dataInput)
}

func (controller *GoElasticsearchController) UpdateGoElasticsearch(contexts ElasticsearchContextInterface, request *http.Request) {
	dataInput, _ := contexts.GetPostForm("data")
	documentID, _ := contexts.GetPostForm("_id")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://es01:9200",
			"http://es02:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// Build the request body.
	var b strings.Builder
	b.WriteString(dataInput)
	req := esapi.IndexRequest{
		Index:      "store",
		DocumentID: documentID,
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Printf("[%s] Error indexing", res.Status())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	contexts.JSON(200, dataInput)
}
