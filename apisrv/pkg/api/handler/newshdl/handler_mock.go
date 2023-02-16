package newshdl

import (
	"context"
)

// MockNew func mock the creation of a news docs service handler
func MockNew() *Handler {

	mockGetFunc := func(ctx context.Context, request *GetRequest) (*GetResponse, error) {
		return &GetResponse{}, nil
	}

	mockSearchFunc := func(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
		return &SearchResponse{}, nil
	}

	return &Handler{
		GetFunc:    mockGetFunc,
		SearchFunc: mockSearchFunc,
	}
}
