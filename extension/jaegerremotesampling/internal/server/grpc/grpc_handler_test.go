// Copyright The OpenTelemetry Authors
// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockSamplingStore struct{}

func (mockSamplingStore) GetSamplingStrategy(_ context.Context, serviceName string) (*api_v2.SamplingStrategyResponse, error) {
	switch serviceName {
	case "error":
		return nil, errors.New("some error")
	case "nil":
		return nil, nil
	}
	return &api_v2.SamplingStrategyResponse{StrategyType: api_v2.SamplingStrategyType_PROBABILISTIC}, nil
}

func (mockSamplingStore) Close() error {
	return nil
}

func TestNewGRPCHandler(t *testing.T) {
	tests := []struct {
		req  *api_v2.SamplingStrategyParameters
		resp *api_v2.SamplingStrategyResponse
		err  string
	}{
		{req: &api_v2.SamplingStrategyParameters{ServiceName: "error"}, err: "some error"},
		{req: &api_v2.SamplingStrategyParameters{ServiceName: "nil"}, resp: nil},
		{req: &api_v2.SamplingStrategyParameters{ServiceName: "foo"}, resp: &api_v2.SamplingStrategyResponse{StrategyType: api_v2.SamplingStrategyType_PROBABILISTIC}},
	}
	h := NewGRPCHandler(mockSamplingStore{})
	for _, test := range tests {
		resp, err := h.GetSamplingStrategy(context.Background(), test.req)
		if test.err != "" {
			require.EqualError(t, err, test.err)
			require.Nil(t, resp)
		} else {
			require.NoError(t, err)
			assert.Equal(t, test.resp, resp)
		}
	}
}
