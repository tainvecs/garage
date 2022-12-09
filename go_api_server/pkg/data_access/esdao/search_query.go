package esdao

import (
	"errors"
)

// TermQuery for elasticsearch
type TermQuery struct {
	Term map[string]Term `json:"term"`
}

// Term is part of the TermQuery
type Term struct {
	Value string `json:"value"`
	Boost string `json:"boost"`
}

// NewTermQuery creates a new elasticsearch TermQuery
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

// TermsQuery for elasticsearch
type TermsQuery struct {
	Terms map[string]interface{} `json:"terms"`
}

// NewTermsQuery creates a new elasticsearch TermsQuery
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

// MatchQuery for elasticsearch
type MatchQuery struct {
	Match map[string]Match `json:"match"`
}

// Match is part of the MatchQuery
type Match struct {
	Query     string `json:"query"`
	Boost     string `json:"boost"`
	Operator  string `json:"operator,omitempty"`
	Fuzziness string `json:"fuzziness,omitempty"`
}

// NewMatchQuery creates a new elasticsearch MatchQuery
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

// MatchPhraseQuery for elasticsearch
type MatchPhraseQuery struct {
	MatchPhrase map[string]MatchPhrase `json:"match_phrase"`
}

// MatchPhrase is part of the MatchPhraseQuery
type MatchPhrase struct {
	Query string `json:"query"`
	Boost string `json:"boost"`
	Slop  string `json:"slop,omitempty"`
}

// NewMatchPhraseQuery creates a new elasticsearch MatchPhraseQuery
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

// MultiMatchQuery for elasticsearch
type MultiMatchQuery struct {
	MultiMatch MultiMatch `json:"multi_match"`
}

// MultiMatch is part of the MultiMatchQuery
type MultiMatch struct {
	Query string   `json:"query"`
	Type  string   `json:"type,omitempty"`
	Field []string `json:"fields"`
}

// NewMultiMatchQuery creates a new elasticsearch MultiMatchQuery
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

// ConstantScoreQuery for elasticsearch
type ConstantScoreQuery struct {
	ConstantScore ConstantScore `json:"constant_score"`
}

// ConstantScore is part of the ConstantScoreQuery
type ConstantScore struct {
	Filter interface{} `json:"filter"`
	Boost  string      `json:"boost"`
}

// NewConstantScoreQuery creates a new elasticsearch ConstantScoreQuery
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

// SimpleQueryStringQuery for elasticsearch
type SimpleQueryStringQuery struct {
	SimpleQueryString SimpleQueryString `json:"simple_query_string"`
}

// SimpleQueryString is part of the SimpleQueryStringQuery
type SimpleQueryString struct {
	Query            string   `json:"query"`
	Fields           []string `json:"fields"`
	QuoteFieldSuffix string   `json:"quote_field_suffix,omitempty"`
	DefaultOperator  string   `json:"default_operator,omitempty"`
}

// NewSimpleQueryStringQuery creates a new elasticsearch SimpleQueryStringQuery
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

// ExistsQuery for elasticsearch
type ExistsQuery struct {
	Exists Exists `json:"exists"`
}

// Exists is part of the ExistsQuery
type Exists struct {
	Field string `json:"field"`
}

// NewExistsQuery creates a new elasticsearch ExistsQuery
func NewExistsQuery(exists Exists) (*ExistsQuery, error) {

	// check input
	if len(exists.Field) == 0 {
		return nil, errors.New("NewExistsQuery: missing exists.Field")
	}

	// result query
	existsQuery := ExistsQuery{
		Exists: exists,
	}

	return &existsQuery, nil
}

// NestedQuery for elasticsearch
type NestedQuery struct {
	Nested Nested `json:"nested"`
}

// Nested is part of the NestedQuery
type Nested struct {
	Path      string      `json:"path"`
	Query     interface{} `json:"query"`
	InnerHits interface{} `json:"inner_hits,omitempty"`
}

// NewNestedQuery creates a new elasticsearch NestedQuery
func NewNestedQuery(nested Nested) (*NestedQuery, error) {

	nestedQ := NestedQuery{
		Nested: nested,
	}

	return &nestedQ, nil
}

// BoolQuery for elasticsearch
type BoolQuery struct {
	Bool Bool `json:"bool"`
}

// Bool is part of the BoolQuery
type Bool struct {
	Filter             []interface{} `json:"filter,omitempty"`
	MustNot            []interface{} `json:"must_not,omitempty"`
	Must               []interface{} `json:"must,omitempty"`
	Should             []interface{} `json:"should,omitempty"`
	MinimumShouldMatch string        `json:"minimum_should_match,omitempty"`
	Boost              string        `json:"boost,omitempty"`
}

// QueryBody for elasticsearch
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

// Source is part of the QueryBody
type Source struct {
	Includes []string `json:"includes,omitempty"`
	Excludes []string `json:"excludes,omitempty"`
}

// TermsAggregation is one kind of Aggregation
type TermsAggregation struct {
	DocCountErrorUpperBound int                      `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int                      `json:"sum_other_doc_count"`
	Buckets                 []TermsAggregationBucket `json:"buckets"`
}

// TermsAggregationBucket is part of the TermsAggregation
type TermsAggregationBucket struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
}

// Highlight is part of the QueryBody
type Highlight struct {
	Order  string                     `json:"order,omitempty"`
	Fields map[string]*HighlightField `json:"fields,omitempty"`
}

// HighlightField is part of the Highlight
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
