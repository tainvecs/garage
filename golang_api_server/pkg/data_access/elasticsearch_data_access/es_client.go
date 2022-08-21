package elasticsearch_data_access

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ESClient struct {
	Client *elasticsearch.Client
}

func (es *ESClient) Index(
	ctx context.Context,
	index string,
	docID string,
	requestBody interface{},
) error {

	// build the request body
	docData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error marshaling document during elasticsearch Indexing: %s", err)
		return err
	}

	// set up the request object
	indexReq := esapi.IndexRequest{
		Index:      index,
		DocumentID: docID,
		Body:       bytes.NewReader(docData),
		Refresh:    "true",
	}

	// perform the request with the client
	resp, err := indexReq.Do(ctx, es.Client)
	if err != nil {
		log.Fatalf("Error getting response during elasticsearch Indexing: %s", err)
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("elasticsearch index request: status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}

func (es *ESClient) Search(
	ctx context.Context,
	index string,
	query string,
) (*RawSearchResponse, error) {

	// perform the search request
	rawResp, err := es.Client.Search(
		es.Client.Search.WithContext(ctx),
		es.Client.Search.WithIndex(index),
		es.Client.Search.WithBody(bytes.NewBufferString(query)),
		es.Client.Search.WithTrackTotalHits(true),
		// es.Client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting the ES search response: %s", err)
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

			log.Fatalf("Error parsing the raw ES search response body: %s", err)
			return nil, err

		} else {

			err = fmt.Errorf("[%s] %s: %s",
				rawResp.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			return nil, err
		}
	}

	// parse raw response from es client
	var rawResponse RawSearchResponse
	if err := json.NewDecoder(rawResp.Body).Decode(&rawResponse); err != nil {
		log.Fatalf("Error parsing the raw ES search response body: %s", err)
		return nil, err
	}

	return &rawResponse, nil
}

func (es *ESClient) Update(
	ctx context.Context,
	index string,
	docID string,
	requestBody interface{},
) error {

	// build the request body
	docData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error marshaling document during elasticsearch Updating: %s", err)
		return err
	}

	// set up the request object
	updateReq := esapi.UpdateRequest{
		Index:      index,
		DocumentID: docID,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, docData))),
		Refresh:    "true",
	}

	// perform the request with the client
	resp, err := updateReq.Do(ctx, es.Client)
	if err != nil {
		log.Fatalf("Error getting response during elasticsearch Updating: %s", err)
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("elasticsearch update request: status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}

func (es *ESClient) Delete(
	ctx context.Context,
	index string,
	docID string,
) error {

	// set up the request object
	deleteReq := esapi.DeleteRequest{
		Index:      index,
		DocumentID: docID,
		Refresh:    "true",
	}

	// perform the request with the client
	resp, err := deleteReq.Do(ctx, es.Client)
	if err != nil {
		log.Fatalf("Error getting response during elasticsearch Deleting: %s", err)
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("elasticsearch delete request: status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}
