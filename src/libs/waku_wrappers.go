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

	fmt.Println("WakuNode started successfully!")
	return nil
}

// Stop stops the Waku node.
func (node *LocalWakuNode) Stop() error {
	return node.WakuStop()
}

// Destroy destroys the Waku node.
func (node *LocalWakuNode) Destroy() error {
	return node.WakuDestroy()
}

// Version returns the Waku node version.
func (node *LocalWakuNode) Version() (string, error) {
	return node.WakuVersion()
}

// RelayPublish publishes a message to a pubsub topic.
func (node *LocalWakuNode) RelayPublish(pubSubTopic, message string, timeoutMs int) error {
	_, err := node.WakuRelayPublish(pubSubTopic, message, timeoutMs)
	return err
}

// RelaySubscribe subscribes to a pubsub topic.
func (node *LocalWakuNode) RelaySubscribe(pubSubTopic string) error {
	return node.WakuRelaySubscribe(pubSubTopic)
}

// RelayUnsubscribe unsubscribes from a pubsub topic.
func (node *LocalWakuNode) RelayUnsubscribe(pubSubTopic string) error {
	return node.WakuRelayUnsubscribe(pubSubTopic)
}

// Connect connects to a peer.
func (node *LocalWakuNode) Connect(peerMultiAddr string, timeoutMs int) error {
	return node.WakuConnect(peerMultiAddr, timeoutMs)
}

// ListenAddresses retrieves the node's listen addresses.
func (node *LocalWakuNode) ListenAddresses() (string, error) {
	return node.WakuListenAddresses()
}

// GetMyENR retrieves the node's ENR.
func (node *LocalWakuNode) GetMyENR() (string, error) {
	return node.WakuGetMyENR()
}
