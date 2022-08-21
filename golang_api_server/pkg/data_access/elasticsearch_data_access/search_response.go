package elasticsearch_data_access

import "encoding/json"

// response
type RawSearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	}
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float32           `json:"mas_score"`
		Hits     []json.RawMessage `json:"hits"`
	} `json:"hits"`
	Aggregations json.RawMessage `json:"aggregations,omitempty"`
	Highlight    json.RawMessage `json:"highlight,omitempty"`
}

// hit
type RawSearchHit struct {
	Index string  `json:"_index"`
	Type  string  `json:"_type"`
	ID    string  `json:"_id"`
	Score float32 `json:"_score"`
}
