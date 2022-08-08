package es_dao

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

type Document struct {
	ID        string     `json:"id"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (dao *DAO) Index(ctx context.Context, doc *Document) error {
	return dao.Client.Index(ctx, dao.IndexIndex, doc.ID, doc)
}

type DocSearchHit struct {
	RawSearchHit
	Source Document `json:"_source"`
}

func rawSearchResponseToDAOSearchResponse(rawResp *RawSearchResponse) (*DAOSearchResponse, error) {

	// parse docs
	var docSlice []*Document
	var scoreSlice []float32
	for _, rawHit := range rawResp.Hits.Hits {

		// parse individual hit
		hit := DocSearchHit{}
		err := json.Unmarshal(rawHit, &hit)
		if err != nil {
			return nil, err
		}

		// append parsed doc
		doc := hit.Source
		doc.ID = hit.ID
		docSlice = append(docSlice, &doc)
		scoreSlice = append(scoreSlice, hit.Score)
	}

	// response
	resp := DAOSearchResponse{
		Total:   rawResp.Hits.Total.Value,
		Results: docSlice,
		Scores:  scoreSlice,
	}

	return &resp, nil
}

func (dao *DAO) Search(ctx context.Context, query string) (*DAOSearchResponse, error) {

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

func (dao *DAO) Update(ctx context.Context, doc *Document) error {
	return dao.Client.Update(ctx, dao.IndexIndex, doc.ID, doc)
}

func (dao *DAO) Delete(ctx context.Context, docID string) error {
	return dao.Client.Delete(ctx, dao.IndexIndex, docID)
}
