package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "erst",
		Short: "Stellar smart contract debugging tool",
		Long: `Erst is a specialized developer tool for the Stellar network,
designed to solve the "black box" debugging experience on Soroban.

It helps clarify why a Stellar smart contract transaction failed by:
  - Fetching failed transaction envelopes and ledger state
  - Re-executing transactions locally for detailed analysis
  - Mapping execution failures back to readable source code`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Register all subcommands using the modular registry pattern
	RegisterCommands(rootCmd)
}

// RegisterCommands registers all subcommands to the root command.
// This function is called during initialization and provides a central
// place to manage command registration, keeping root.go clean and focused.
func RegisterCommands(root *cobra.Command) {
	// Commands are registered in alphabetical order to maintain
	// consistent help output ordering

	// Register the debug command
	registerDebugCommand(root)

	// Register the version command
	registerVersionCommand(root)

	// Future commands can be registered here:
	// registerAnalyzeCommand(root)
	// registerReplayCommand(root)
	// registerTraceCommand(root)
}

// getErrWriter returns the error writer for commands
func getErrWriter() *os.File {
	return os.Stderr
}

// getOutWriter returns the standard output writer for commands
func getOutWriter() *os.File {
	return os.Stdout
}

// printError prints an error message to stderr
func printError(err error) {
	fmt.Fprintf(getErrWriter(), "Error: %v\n", err)
}
