package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWakuSetup(t *testing.T) {
	config := WakuConfig{
		Host:        "0.0.0.0",
		Port:        30304,
		NodeKey:     "11d0dcea28e86f81937a3bd1163473c7fbc0a0db54fd72914849bc47bdf78710",
		EnableRelay: true,
		LogLevel:    "DEBUG",
	}

	node, err := WakuNew(config)
	require.NoError(t, err, "WakuNew should not return an error")
	require.NotNil(t, node, "WakuNode should not be nil")

	err = node.WakuStart()
	assert.NoError(t, err, "WakuStart should not return an error")

}
