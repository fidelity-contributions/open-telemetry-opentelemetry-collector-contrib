// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package purefbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefbreceiver"

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefbreceiver/internal/metadata"
)

func TestStartAndShutdown(t *testing.T) {
	// prepare
	cfg, ok := createDefaultConfig().(*Config)
	require.True(t, ok)

	sink := &consumertest.MetricsSink{}
	recv := newReceiver(cfg, receivertest.NewNopSettings(metadata.Type), sink)

	// verify
	assert.NoError(t, recv.Start(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, recv.Shutdown(context.Background()))
}
