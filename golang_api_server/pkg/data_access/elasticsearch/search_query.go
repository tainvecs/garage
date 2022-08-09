package es_dao

type QueryBody struct {
	From      int           `json:"from,omitempty"`
	Size      int           `json:"size,omitempty"`
	MinScore  float32       `json:"min_score,omitempty"`
	Source    Source        `json:"_source,omitempty"`
	Query     interface{}   `json:"query"`
	Highlight Highlight     `json:"highlight,omitempty"`
	Aggs      interface{}   `json:"aggs,omitempty"`
	Sort      []interface{} `json:"sort,omitempty"`
}

// query
type BoolQuery struct {
	Bool struct {
		Filter             []interface{} `json:"filter,omitempty"`
		MustNot            []interface{} `json:"must_not,omitempty"`
		Must               []interface{} `json:"must,omitempty"`
		Should             []interface{} `json:"should,omitempty"`
		MinimumShouldMatch string        `json:"minimum_should_match,omitempty"`
		Boost              string        `json:"boost,omitempty"`
	} `json:"bool"`
}

type TermQuery struct {
	Term map[string]struct {
		Value string `json:"value"`
		Boost string `json:"boost"`
	} `json:"term"`
}

type TermsQuery struct {
	Terms map[string]interface{} `json:"terms"`
}

func NewTermsQuery(field string, value []string, boost string) *TermsQuery {

	if len(field) == 0 || len(value) == 0 {
		return nil
	}

	if len(boost) == 0 {
		boost = "1.0"
	}

	return &TermsQuery{
		Terms: map[string]interface{}{
			field:   value,
			"boost": boost,
		},
	}
}

type MatchQuery struct {
	Match map[string]struct {
		Query    string `json:"query"`
		Boost    string `json:"boost"`
		Operator string `json:"operator,omitempty"`
	} `json:"match"`
}

type MatchPhrase struct {
	MatchPhrase map[string]struct {
		Query string `json:"query"`
		Boost string `json:"boost"`
		Slop  string `json:"slop,omitempty"`
	} `json:"match_phrase"`
}

type MultiMatchQuery struct {
	MultiMatch struct {
		Query string   `json:"query"`
		Type  string   `json:"type"`
		Field []string `json:"fields"`
	} `json:"multi_match"`
}

type ConstantScoreQuery struct {
	ConstantScore struct {
		Filter interface{} `json:"filter"`
		Boost  string      `json:"boost"`
	} `json:"constant_score"`
}

type SimpleQueryStringESQuery struct {
	SimpleQueryString struct {
		Query            string   `json:"query"`
		Fields           []string `json:"fields"`
		QuoteFieldSuffix string   `json:"quote_field_suffix,omitempty"`
		DefaultOperator  string   `json:"default_operator,omitempty"`
	} `json:"simple_query_string"`
}

type NestedQuery struct {
	Nested struct {
		Path      string      `json:"path"`
		Query     BoolQuery   `json:"query"`
		InnerHits interface{} `json:"inner_hits,omitempty"`
	} `json:"nested"`
}

// source
type Source struct {
	Includes []string `json:"includes,omitempty"`
	Excludes []string `json:"excludes,omitempty"`
}

// aggregation
type TermsAggregation struct {
	DocCountErrorUpperBound int                      `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int                      `json:"sum_other_doc_count"`
	Buckets                 []TermsAggregationBucket `json:"buckets"`
}

type TermsAggregationBucket struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
}

// highlight
type Highlight struct {
	Order  string                     `json:"order,omitempty"`
	Fields map[string]*HighlightField `json:"fields,omitempty"`
}

type HighlightField struct {
	PreTags        []string    `json:"pre_tags,omitempty"`
	PostTags       []string    `json:"post_tags,omitempty"`
	NumOfFragments int         `json:"number_of_fragments,omitempty"`
	FragmentSize   int         `json:"fragment_size,omitempty"`
	NoMatchSize    int         `json:"no_match_size,omitempty"`
	MatchedFields  []string    `json:"matched_fields,omitempty"`
	HighlightType  string      `json:"type,omitempty"`
	HighlightQuery interface{} `json:"highlight_query,omitempty"`
}
