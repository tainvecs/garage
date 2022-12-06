package news_doc_service

import (
	"context"
	"encoding/json"
	"log"

	es_data_access "api-server/pkg/data_access/elasticsearch_data_access"
)

type ESDAOSearchResponse struct {
	Total   int
	Results []*NewsDoc
	Scores  []float32
}

func (dao *newsDocESDAO) Search(ctx context.Context, query string) (*ESDAOSearchResponse, error) {

	// run es search
	rawResp, err := dao.Client.Search(ctx, dao.SearchIndex, query)
	if err != nil {
		return nil, err
	}

	// parsing raw ES response to DAO response
	daoResp, err := rawSearchResponseToDAOSearchResponse(rawResp)
	if err != nil {
		log.Fatalf("Error parsing the raw ES search response to the DAO response: %s", err)
		return nil, err
	}

	return daoResp, nil
}

type esRawSearchHit struct {
	es_data_access.RawSearchHit
	Source NewsDoc `json:"_source"`
}

// parse raw ES search response
func rawSearchResponseToDAOSearchResponse(
	rawResp *es_data_access.RawSearchResponse,
) (*ESDAOSearchResponse, error) {

	// parse docs
	var docSlice []*NewsDoc
	var scoreSlice []float32
	for _, rawHit := range rawResp.Hits.Hits {

		// parse individual hit
		hit := esRawSearchHit{}
		err := json.Unmarshal(rawHit, &hit)
		if err != nil {
			return nil, err
		}

		// append parsed doc
		doc := hit.Source
		doc.UUID = hit.ID
		docSlice = append(docSlice, &doc)
		scoreSlice = append(scoreSlice, hit.Score)
	}

	// response
	resp := ESDAOSearchResponse{
		Total:   rawResp.Hits.Total.Value,
		Results: docSlice,
		Scores:  scoreSlice,
	}

	return &resp, nil
}
