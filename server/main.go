package main

import (
	"context"

	"github.com/kitchen/redish/redish"
)

type redishServer struct {
}

func (s *redishServer) Get(ctx context.Context, key *redish.Key) (*redish.SingleValue, error) {

	return &redish.SingleValue{Value: "abc"}, nil
}
