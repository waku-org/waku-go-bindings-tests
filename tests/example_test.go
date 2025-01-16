package main

import (
	"testing"
	"time"
	"waku-go-bindings-tests/src/libs"
	"waku-go-bindings-tests/src/utilities"
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestWakuNodeLifecycle(t *testing.T) {

	utilities.LogDebug("Start node1 ")
	node1, err := libs.StartWakuNodeWithDefaultValues("", 0, "", true, "")

	if err != nil {
		t.Fatalf("Failed to setup and start Waku Node: %v", err)
	}

	defer func() {

		utilities.LogDebug("Stop & destroy waku node1 at the end ")
		if stopErr := node1.Stop(); stopErr != nil {
			t.Errorf("Failed to stop Waku node: %v", stopErr)
		}
		if destroyErr := node1.Destroy(); destroyErr != nil {
			t.Errorf("Failed to destroy Waku node: %v", destroyErr)
		}
	}()

	utilities.LogDebug("Sleep for 2 seconds")
	time.Sleep(2 * time.Second)

	pubsubTopic, err := node1.WakuDefaultPubsubTopic()
	if err != nil {
		t.Fatalf("Failed to get default PubSub topic: %v", err)
	}
	utilities.LogDebug("Node1 subscribed to default pubsubtopic")

	err = node1.WakuRelaySubscribe(pubsubTopic)
	if err != nil {
		t.Fatalf("Failed to subscribe to relay: %v", err)
	}

	utilities.LogDebug("Successfully subcribed node1 to topic " + pubsubTopic)

	utilities.LogDebug("Get node1 ENR")
	enr, err := node1.WakuGetMyENR()
	if err != nil {
		t.Fatalf("Failed to retrieve ENR: %v", err)
	}

	if enr == "" {
		t.Fatal("ENR is empty, expected a valid ENR")
	}

	t.Logf("Successfully retrieved ENR: %s", enr)
}
