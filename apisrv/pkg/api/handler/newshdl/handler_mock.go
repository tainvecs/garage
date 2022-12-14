package newshdl

import (
	"context"
)

// MockNew func mock the creation of a news docs service handler
func MockNew() *Handler {

	mockSearchFunc := func(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
		return &SearchResponse{}, nil
	}

	return &Handler{
		SearchFunc: mockSearchFunc,
	}
}
