package esdao

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// DataAccessObject is the data access object for elasticsearch index and search
type DataAccessObject struct {
	Client      *elasticsearch.Client
	IndexIndex  string
	SearchIndex string
}

// New function instansiate a new DataAccessObject
func New(esURL, esIndexIndex, esSearchIndex string) (*DataAccessObject, error) {

	// new elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{esURL},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// new elasticsearch data access object
	dao := DataAccessObject{
		Client:      client,
		IndexIndex:  esIndexIndex,
		SearchIndex: esSearchIndex,
	}

	return &dao, nil
}

// Index elasticsearch docs
func (dao *DataAccessObject) Index(
	ctx context.Context,
	docID string,
	requestBody interface{},
) error {

	// build the request body
	docData, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Error marshaling document during elasticsearch Indexing: %s", err)
		return err
	}

	// set up the request object
	indexReq := esapi.IndexRequest{
		Index:      dao.IndexIndex,
		DocumentID: docID,
		Body:       bytes.NewReader(docData),
		Refresh:    "true",
	}

	// perform the request with the client
	resp, err := indexReq.Do(ctx, dao.Client)
	if err != nil {
		log.Printf("Error getting response during elasticsearch Indexing: %s", err)
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		err := fmt.Errorf("elasticsearch index request: status code %d", resp.StatusCode)
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

// Search elasticsearch docs
func (dao *DataAccessObject) Search(
	ctx context.Context,
	query string,
) (*RawSearchResponse, error) {

	// perform the search request
	rawResp, err := dao.Client.Search(
		dao.Client.Search.WithContext(ctx),
		dao.Client.Search.WithIndex(dao.SearchIndex),
		dao.Client.Search.WithBody(bytes.NewBufferString(query)),
		dao.Client.Search.WithTrackTotalHits(true),
		// es.Client.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting the ES search response: %s", err)
		return nil, err
	}
	if rawResp.StatusCode != 200 {
		return nil, fmt.Errorf("elasticsearch search request: status code %d", rawResp.StatusCode)
	}
	defer rawResp.Body.Close()

	// parse raw response err
	if rawResp.IsError() {

		var e map[string]interface{}

		if err := json.NewDecoder(rawResp.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the raw ES search response body: %s", err)
			return nil, err
		}

		err = fmt.Errorf("[%s] %s: %s",
			rawResp.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
		return nil, err
	}

	// parse raw response from es client
	var rawResponse RawSearchResponse
	if err := json.NewDecoder(rawResp.Body).Decode(&rawResponse); err != nil {
		log.Printf("Error parsing the raw ES search response body: %s", err)
		return nil, err
	}

	return &rawResponse, nil
}

// Update elasticsearch docs
func (dao *DataAccessObject) Update(
	ctx context.Context,
	docID string,
	requestBody interface{},
) error {

	// build the request body
	docData, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Error marshaling document during elasticsearch Updating: %s", err)
		return err
	}

	// set up the request object
	updateReq := esapi.UpdateRequest{
		Index:      dao.IndexIndex,
		DocumentID: docID,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, docData))),
		Refresh:    "true",
	}

	// perform the request with the client
	resp, err := updateReq.Do(ctx, dao.Client)
	if err != nil {
		log.Printf("Error getting response during elasticsearch Updating: %s", err)
		return err
	}
	if resp.StatusCode != 200 {
		err := fmt.Errorf("elasticsearch update request: status code %d", resp.StatusCode)
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

// Delete elasticsearch docs
func (dao *DataAccessObject) Delete(
	ctx context.Context,
	docID string,
) error {

	// set up the request object
	deleteReq := esapi.DeleteRequest{
		Index:      dao.IndexIndex,
		DocumentID: docID,
		Refresh:    "true",
	}

	// perform the request with the client
	resp, err := deleteReq.Do(ctx, dao.Client)
	if err != nil {
		log.Printf("Error getting response during elasticsearch Deleting: %s", err)
		return err
	}
	if resp.StatusCode != 200 {
		err := fmt.Errorf("elasticsearch delete request: status code %d", resp.StatusCode)
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
