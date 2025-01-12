package main

type WakuConfigData struct {
	Host        string `json:"host,omitempty"`
	Port        int    `json:"port,omitempty"`
	NodeKey     string `json:"key,omitempty"`
	EnableRelay bool   `json:"relay"`
	LogLevel    string `json:"logLevel"`
}

// Start starts the Waku node.
func StartWakuNode(host string, port int, nodeKey string, enableRelay bool, logLevel string) WakuConfigData {
	return WakuConfigData{
		Host:        host,
		Port:        port,
		NodeKey:     nodeKey,
		EnableRelay: enableRelay,
		LogLevel:    logLevel,
	}
}

// Stop stops the Waku node.
func (node *WakuNode) Stop() error {
	return node.WakuStop()
}

// Destroy destroys the Waku node.
func (node *WakuNode) Destroy() error {
	return node.WakuDestroy()
}

// Version returns the Waku node version.
func (node *WakuNode) Version() (string, error) {
	return node.WakuVersion()
}

// RelayPublish publishes a message to a pubsub topic.
func (node *WakuNode) RelayPublish(pubSubTopic, message string, timeoutMs int) error {
	_, err := node.WakuRelayPublish(pubSubTopic, message, timeoutMs)
	return err
}

// RelaySubscribe subscribes to a pubsub topic.
func (node *WakuNode) RelaySubscribe(pubSubTopic string) error {
	return node.WakuRelaySubscribe(pubSubTopic)
}

// RelayUnsubscribe unsubscribes from a pubsub topic.
func (node *WakuNode) RelayUnsubscribe(pubSubTopic string) error {
	return node.WakuRelayUnsubscribe(pubSubTopic)
}

// Connect connects to a peer.
func (node *WakuNode) Connect(peerMultiAddr string, timeoutMs int) error {
	return node.WakuConnect(peerMultiAddr, timeoutMs)
}

// ListenAddresses retrieves the node's listen addresses.
func (node *WakuNode) ListenAddresses() (string, error) {
	return node.WakuListenAddresses()
}

// GetMyENR retrieves the node's ENR.
func (node *WakuNode) GetMyENR() (string, error) {
	return node.WakuGetMyENR()
}
