package newssvc

import (
	"context"
	"encoding/json"
	"log"

	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
)

// ESSearchResponse is the elasticsearch search response of ESNewsDoc
type ESSearchResponse struct {
	Total   int
	Results []*ESNewsDoc
	Scores  []float32
}

// Search ESNewsDoc in elasticsearch
func (dao *esDAO) Search(ctx context.Context, query string) (*ESSearchResponse, error) {

	// run es search
	rawResp, err := dao.DataAccessObject.Search(ctx, query)
	if err != nil {
		log.Printf("News docs ES search: %s", err)
		return nil, err
	}

	// parsing raw ES response to DAO response
	daoResp, err := rawSearchResponseToDAOSearchResponse(rawResp)
	if err != nil {
		log.Printf("Error parsing the raw ES search response to the DAO response: %s", err)
		return nil, err
	}

	return daoResp, nil
}

// esRawSearchHit is the raw elasticsearch search hit of ESNewsDoc
type esRawSearchHit struct {
	esdao.RawSearchHit
	Source ESNewsDoc `json:"_source"`
}

// rawSearchResponseToDAOSearchResponse parse the raw ES search response of ESNewsDoc
func rawSearchResponseToDAOSearchResponse(
	rawResp *esdao.RawSearchResponse,
) (*ESSearchResponse, error) {

	// parse docs
	var docSlice []*ESNewsDoc
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
	resp := ESSearchResponse{
		Total:   rawResp.Hits.Total.Value,
		Results: docSlice,
		Scores:  scoreSlice,
	}

	return &resp, nil
}
