package elasticsearch_data_access

import (
	"errors"
)

// Term Query
type TermQuery struct {
	Term map[string]Term `json:"term"`
}

type Term struct {
	Value string `json:"value"`
	Boost string `json:"boost"`
}

func NewTermQuery(field string, term Term) (*TermQuery, error) {

	// check input
	if len(field) == 0 {
		return nil, errors.New("NewTermQuery: missing \"field\" parameter")
	}
	if len(term.Value) == 0 {
		return nil, errors.New("NewTermQuery: missing term.Value")
	}
	if len(term.Boost) == 0 {
		term.Boost = "1.0"
	}

	// result query
	termQ := TermQuery{
		map[string]Term{
			field: term,
		},
	}

	return &termQ, nil
}

// Terms Query
type TermsQuery struct {
	Terms map[string]interface{} `json:"terms"`
}

func NewTermsQuery(field string, values []string, boost string) (*TermsQuery, error) {

	// check input
	if len(field) == 0 {
		return nil, errors.New("NewTermQuery: missing \"field\" parameter")
	}
	if len(values) == 0 {
		return nil, errors.New("NewTermQuery: missing values")
	}
	if len(boost) == 0 {
		boost = "1.0"
	}

	// result query
	tremsQ := TermsQuery{
		Terms: map[string]interface{}{
			field:   values,
			"boost": boost,
		},
	}

	return &tremsQ, nil
}

// Match Query
type MatchQuery struct {
	Match map[string]Match `json:"match"`
}

type Match struct {
	Query     string `json:"query"`
	Boost     string `json:"boost"`
	Operator  string `json:"operator,omitempty"`
	Fuzziness string `json:"fuzziness,omitempty"`
}

func NewMatchQuery(field string, match Match) (*MatchQuery, error) {

	// check input
	if len(field) == 0 {
		return nil, errors.New("NewMatchQuery: missing \"field\" parameter")
	}
	if len(match.Query) == 0 {
		return nil, errors.New("NewMatchQuery: missing match.Query")
	}
	if len(match.Boost) == 0 {
		match.Boost = "1.0"
	}

	// result query
	matchQ := MatchQuery{
		map[string]Match{
			field: match,
		},
	}

	return &matchQ, nil
}

// Match Phrase Query
type MatchPhraseQuery struct {
	MatchPhrase map[string]MatchPhrase `json:"match_phrase"`
}

type MatchPhrase struct {
	Query string `json:"query"`
	Boost string `json:"boost"`
	Slop  string `json:"slop,omitempty"`
}

func NewMatchPhraseQuery(field string, matchPhrase MatchPhrase) (*MatchPhraseQuery, error) {

	// check input
	if len(field) == 0 {
		return nil, errors.New("NewMatchQuery: missing \"field\" parameter")
	}
	if len(matchPhrase.Query) == 0 {
		return nil, errors.New("NewMatchQuery: missing matchPhrase.Query")
	}
	if len(matchPhrase.Boost) == 0 {
		matchPhrase.Boost = "1.0"
	}

	// result query
	matchPhraseQ := MatchPhraseQuery{
		map[string]MatchPhrase{
			field: matchPhrase,
		},
	}

	return &matchPhraseQ, nil
}

// Multi Match Query
type MultiMatchQuery struct {
	MultiMatch MultiMatch `json:"multi_match"`
}

type MultiMatch struct {
	Query string   `json:"query"`
	Type  string   `json:"type,omitempty"`
	Field []string `json:"fields"`
}

func NewMultiMatchQuery(multiMatch MultiMatch) (*MultiMatchQuery, error) {

	// check input
	if len(multiMatch.Query) == 0 {
		return nil, errors.New("NewMultiMatchQuery: missing multiMatch.Query")
	}
	if len(multiMatch.Field) == 0 {
		return nil, errors.New("NewMultiMatchQuery: missing multiMatch.Field")
	}

	// result query
	multiMatchQ := MultiMatchQuery{
		MultiMatch: multiMatch,
	}

	return &multiMatchQ, nil
}

// Constant Score Query
type ConstantScoreQuery struct {
	ConstantScore ConstantScore `json:"constant_score"`
}

type ConstantScore struct {
	Filter interface{} `json:"filter"`
	Boost  string      `json:"boost"`
}

func NewConstantScoreQuery(constantScore ConstantScore) (*ConstantScoreQuery, error) {

	// check input
	if constantScore.Filter == nil {
		return nil, errors.New("NewConstantScoreQuery: missing constantScore.Filter")
	}
	if len(constantScore.Boost) == 0 {
		constantScore.Boost = "1.0"
	}

	// result query
	constantScoreQ := ConstantScoreQuery{
		ConstantScore: constantScore,
	}

	return &constantScoreQ, nil
}

// Simple Query String Query
type SimpleQueryStringQuery struct {
	SimpleQueryString SimpleQueryString `json:"simple_query_string"`
}

type SimpleQueryString struct {
	Query            string   `json:"query"`
	Fields           []string `json:"fields"`
	QuoteFieldSuffix string   `json:"quote_field_suffix,omitempty"`
	DefaultOperator  string   `json:"default_operator,omitempty"`
}

func NewSimpleQueryStringQuery(simpleQueryString SimpleQueryString) (*SimpleQueryStringQuery, error) {

	// check input
	if len(simpleQueryString.Query) == 0 {
		return nil, errors.New("NewSimpleQueryStringQuery: missing simpleQueryString.Query")
	}
	if len(simpleQueryString.Fields) == 0 {
		return nil, errors.New("NewSimpleQueryStringQuery: missing simpleQueryString.Fields")
	}

	// result query
	simpleQueryStringQ := SimpleQueryStringQuery{
		SimpleQueryString: simpleQueryString,
	}

	return &simpleQueryStringQ, nil
}

// NestedQuery
type NestedQuery struct {
	Nested Nested `json:"nested"`
}

// Nested
type Nested struct {
	Path      string      `json:"path"`
	Query     BoolQuery   `json:"query"`
	InnerHits interface{} `json:"inner_hits,omitempty"`
}

func NewNestedQuery(nested Nested) (*NestedQuery, error) {

	nestedQ := NestedQuery{
		Nested: nested,
	}

	return &nestedQ, nil
}

// Query body
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

// Bool Query
type BoolQuery struct {
	Bool Bool `json:"bool"`
}

type Bool struct {
	Filter             []interface{} `json:"filter,omitempty"`
	MustNot            []interface{} `json:"must_not,omitempty"`
	Must               []interface{} `json:"must,omitempty"`
	Should             []interface{} `json:"should,omitempty"`
	MinimumShouldMatch string        `json:"minimum_should_match,omitempty"`
	Boost              string        `json:"boost,omitempty"`
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
