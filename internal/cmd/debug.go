package cmd

import (
	"fmt"

	"github.com/dotandev/hintents/internal/rpc"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// registerDebugCommand registers the debug command with the root command.
// This function is called from RegisterCommands in root.go.
func registerDebugCommand(root *cobra.Command) {
	root.AddCommand(newDebugCommand())
}

// newDebugCommand creates and returns the debug command.
// Keeping command creation separate from registration allows for
// better testing and modularity.
func newDebugCommand() *cobra.Command {
	var (
		network string
		verbose bool
	)

	cmd := &cobra.Command{
		Use:   "debug <transaction-hash>",
		Short: "Debug a failed Stellar transaction",
		Long: `Fetch and analyze a failed Stellar smart contract transaction.

This command retrieves the transaction envelope from the Stellar network
and provides detailed information about why the transaction failed.

Example:
  erst debug abc123def456... --network testnet
  erst debug abc123def456... --network mainnet --verbose`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			return runDebug(txHash, network, verbose)
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&network, "network", "n", "testnet", "Stellar network to use (testnet, mainnet)")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	return cmd
}

// runDebug executes the debug command logic.
// Separating the logic from command creation makes it easier to test.
func runDebug(txHash, network string, verbose bool) error {
	// Validate network
	if network != "testnet" && network != "mainnet" {
		return fmt.Errorf("invalid network: %s (must be 'testnet' or 'mainnet')", network)
	}

	// Print header
	color.Cyan("🔍 Debugging transaction: %s", txHash)
	color.Cyan("📡 Network: %s", network)
	fmt.Println()

	// Initialize RPC client
	client, err := rpc.NewClient(network)
	if err != nil {
		return fmt.Errorf("failed to initialize RPC client: %w", err)
	}

	// Fetch transaction
	if verbose {
		color.Yellow("Fetching transaction envelope...")
	}

	tx, err := client.GetTransaction(txHash)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction: %w", err)
	}

	// Display transaction information
	color.Green("✓ Transaction found")
	fmt.Printf("  Envelope size: %d bytes\n", len(tx.EnvelopeXDR))
	fmt.Printf("  Result XDR size: %d bytes\n", len(tx.ResultXDR))

	if verbose {
		fmt.Printf("  Ledger: %d\n", tx.Ledger)
	}

	// TODO: Add simulation and trace decoding
	fmt.Println()
	color.Yellow("⚠️  Simulation and trace decoding coming soon...")

	return nil
}
