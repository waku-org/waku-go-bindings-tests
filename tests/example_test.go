package main

import (
	"testing"
	"waku-go-bindings-tests/src/libs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWakuSetup(t *testing.T) {
	
	config := libs.ConfigWakuNode(
		"0.0.0.0",
		30304,
		"11d0dcea28e86f81937a3bd1163473c7fbc0a0db54fd72914849bc47bdf78710",
		true,
		"DEBUG",
	)

	node, err := libs.CreateWakuNodeWithConfig(config)
	require.NoError(t, err, "WakuNew should not return an error")
	require.NotNil(t, node, "WakuNode should not be nil")

    node.
	err = node.WakuStart()
	assert.NoError(t, err, "WakuStart should not return an error")

}
