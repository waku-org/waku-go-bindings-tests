package libs

import (
	"fmt"
	"waku-go-bindings-tests/src/nwaku/examples/golang"
	"waku-go-bindings-tests/src/utilities"
)

type WakuConfigData struct {
	LocalConfigData golang.WakuConfig
}

type LocalWakuNode struct {
	Node   *golang.WakuNode
	Config WakuConfigData
}

// ConfigWakuNode generates a WakuConfigData with provided parameters.
func ConfigWakuNode(host string, port int, nodeKey string, enableRelay bool, logLevel string) WakuConfigData {
	return WakuConfigData{
		LocalConfigData: golang.WakuConfig{
			Host:        utilities.IfEmpty(host, "127.0.0.1"),
			Port:        utilities.IfZero(port, 30304),
			NodeKey:     utilities.IfEmpty(nodeKey, "default-node-key"),
			EnableRelay: enableRelay,
			LogLevel:    utilities.IfEmpty(logLevel, "INFO"),
		},
	}
}

// CreateWakuNode initializes a new Local Waku Node .
func CreateWakuNode(config WakuConfigData) (*LocalWakuNode, error) {
	node, err := golang.WakuNew(config.LocalConfigData)
	if err != nil {
		utilities.LogError("Failed to create WakuNode: " + err.Error())
		return nil, err
	}

	fmt.Println("WakuNode created successfully!")
	return &LocalWakuNode{Node: node, Config: config}, nil
}

// This function starts an existing Waku Node.
func StartWakuNode(localNode *LocalWakuNode) error {
	if err := localNode.Node.WakuStart(); err != nil {
		utilities.LogError("Failed to start WakuNode: " + err.Error())
		return err
	}

	utilities.LogDebug("WakuNode started successfully!")
	return nil
}

func StartWakuNodeWithDefaultValues(host string, port int, nodeKey string, enableRelay bool, logLevel string) (*LocalWakuNode, error) {

	utilities.LogDebug("Configure waku node with default values: ")
	config := ConfigWakuNode("", 0, "", true, "")

	utilities.LogDebug("Create waku node")
	node, err := CreateWakuNode(config)
	if err != nil {
		utilities.LogError("Failed to create Waku Node: " + err.Error())
		return nil, err
	}

	utilities.LogDebug("Start waku node ")
	err = StartWakuNode(node)
	if err != nil {
		utilities.LogError("Failed to start Waku Node: " + err.Error())
		return nil, err
	}

	utilities.LogDebug("Waku Node configured, created, and started successfully!")
	return node, err

}

// Stop stops the Waku node.
func (localNode *LocalWakuNode) Stop() error {
	if err := localNode.Node.WakuStop(); err != nil {
		utilities.LogError("Failed to stop WakuNode: " + err.Error())
		return err
	}

	utilities.LogDebug("WakuNode stoped successfully!")
	return nil
}

// Destroy destroys the Waku node.
func (localNode *LocalWakuNode) Destroy() error {
	if err := localNode.Node.WakuDestroy(); err != nil {
		utilities.LogError("Failed to Destroy WakuNode: " + err.Error())
		return err
	}

	utilities.LogDebug("WakuNode destroyed successfully!")
	return nil
}

// function returns the Waku node version.
func (localNode *LocalWakuNode) Version() (string, error) {
	version, err := localNode.Node.WakuVersion()
	if err != nil {
		utilities.LogError("Error retrieving Waku version: %v\n" + err.Error())
		return "", err
	}

	utilities.LogDebug("Waku version: %s\n" + version)
	return version, nil
}

// Function to return formatted contetn topic
func (localNode *LocalWakuNode) FormatContentTopic(appName string, appVersion int,
	contentTopicName string,
	encoding string,
) (string, error) {
	contentTopic, err := localNode.Node.FormatContentTopic(
		appName,
		appVersion,
		contentTopicName,
		encoding,
	)
	if err != nil {
		utilities.LogError("Error formatting content topic: %v\n" + err.Error())
		return "", err
	}

	utilities.LogDebug("Formatted content topic: %s\n" + contentTopic)
	return contentTopic, nil
}

// Function to format pubsubtopic
func (localNode *LocalWakuNode) FormatPubsubTopic(topicName string) (string, error) {
	pubsubTopic, err := localNode.Node.FormatPubsubTopic(topicName)
	if err != nil {
		utilities.LogError("Error formatting pubsub topic: %v\n" + err.Error())
		return "", err
	}

	utilities.LogDebug("Formatted pubsub topic: %s\n" + pubsubTopic)
	return pubsubTopic, nil
}

// Function to get default pubsubtopic

func (localNode *LocalWakuNode) WakuDefaultPubsubTopic() (string, error) {
	defaultPubsubTopic, err := localNode.Node.WakuDefaultPubsubTopic()
	if err != nil {
		utilities.LogError("Error retrieving default pubsub topic: %v\n" + err.Error())
		return "", err
	}

	utilities.LogDebug("Default pubsub topic: %s\n" + defaultPubsubTopic)
	return defaultPubsubTopic, nil
}

// Function publishes a message to a pubsub topic.

func (localNode *LocalWakuNode) WakuRelayPublish(
	pubsubTopic string,
	message string,
	timeoutMs int,
) (string, error) {
	msgHash, err := localNode.Node.WakuRelayPublish(pubsubTopic, message, timeoutMs)
	if err != nil {
		utilities.LogError("Error publishing Waku relay message: %v\n" + err.Error())
		return "", err
	}

	utilities.LogDebug("Waku relay message hash: %s\n" + msgHash)
	return msgHash, nil
}

// Function subscribes node to a relay.

func (localNode *LocalWakuNode) WakuRelaySubscribe(pubsubTopic string) error {
	err := localNode.Node.WakuRelaySubscribe(pubsubTopic)
	if err != nil {
		utilities.LogError("Error subscribing to Waku relay: " + err.Error())
		return err
	}

	utilities.LogDebug("Successfully subscribed to Waku relay topic: " + pubsubTopic)
	return nil
}

// Function unsubscribes node  from a pubsub topic.

func (localNode *LocalWakuNode) WakuRelayUnsubscribe(pubsubTopic string) error {
	err := localNode.Node.WakuRelayUnsubscribe(pubsubTopic)
	if err != nil {
		utilities.LogError("Error unsubscribing from Waku relay topic: " + err.Error())
		return err
	}

	utilities.LogDebug("Successfully unsubscribed from Waku relay topic: " + pubsubTopic)
	return nil
}

// Function connects to a peer.
func (localNode *LocalWakuNode) WakuConnect(peerMultiAddr string, timeoutMs int) error {
	err := localNode.Node.WakuConnect(peerMultiAddr, timeoutMs)
	if err != nil {
		utilities.LogError("Error connecting to Waku peer: " + err.Error())
		return err
	}

	utilities.LogDebug("Successfully connected to Waku peer: " + peerMultiAddr)
	return nil
}

// Function retrieves the node's listen addresses.
func (localNode *LocalWakuNode) WakuListenAddresses() (string, error) {
	listenAddresses, err := localNode.Node.WakuListenAddresses()
	if err != nil {
		utilities.LogError("Error retrieving Waku listen addresses: " + err.Error())
		return "", err
	}

	utilities.LogDebug("Waku listen addresses: " + listenAddresses)
	return listenAddresses, nil
}

// Function retrieves the node's ENR.
func (localNode *LocalWakuNode) WakuGetMyENR() (string, error) {
	myENR, err := localNode.Node.WakuGetMyENR()
	if err != nil {
		utilities.LogError("Error retrieving Waku ENR: " + err.Error())
		return "", err
	}

	utilities.LogDebug("Waku ENR: " + myENR)
	return myENR, nil
}
