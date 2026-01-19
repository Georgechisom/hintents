package rpc

import (
	"context"
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
)

// Client handles interactions with the Stellar Network
type Client struct {
	Horizon *horizonclient.Client
}

// NewClient creates a new RPC client (defaults to Public Network for now)
func NewClient() *Client {
	return &Client{
		Horizon: horizonclient.DefaultPublicNetClient,
	}
}

// GetTransaction fetches the transaction details and envelope XDR
func (c *Client) GetTransaction(ctx context.Context, hash string) (string, error) {
	tx, err := c.Horizon.TransactionDetail(hash)
	if err != nil {
		return "", fmt.Errorf("failed to fetch transaction: %w", err)
	}

	return tx.EnvelopeXdr, nil
}
