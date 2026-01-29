package rpc

import (
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
)

// Client handles interactions with the Stellar Network
type Client struct {
	Horizon *horizonclient.Client
	Network string
}

// NewClient creates a new RPC client for the specified network
func NewClient(network string) (*Client, error) {
	var horizonClient *horizonclient.Client

	switch network {
	case "testnet":
		horizonClient = horizonclient.DefaultTestNetClient
	case "mainnet":
		horizonClient = horizonclient.DefaultPublicNetClient
	default:
		return nil, fmt.Errorf("unsupported network: %s", network)
	}

	return &Client{
		Horizon: horizonClient,
		Network: network,
	}, nil
}

// TransactionResponse contains the raw XDR fields needed for simulation
type TransactionResponse struct {
	EnvelopeXDR   string
	ResultXDR     string
	ResultMetaXDR string
	Ledger        int32
}

// GetTransaction fetches the transaction details and full XDR data
func (c *Client) GetTransaction(hash string) (*TransactionResponse, error) {
	tx, err := c.Horizon.TransactionDetail(hash)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction: %w", err)
	}

	return &TransactionResponse{
		EnvelopeXDR:   tx.EnvelopeXdr,
		ResultXDR:     tx.ResultXdr,
		ResultMetaXDR: tx.ResultMetaXdr,
		Ledger:        tx.Ledger,
	}, nil
}
