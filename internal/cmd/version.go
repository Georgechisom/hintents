package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	// Version is the current version of erst
	Version = "0.1.0-alpha"
	// BuildDate can be set during build time
	BuildDate = "dev"
)

// registerVersionCommand registers the version command with the root command.
// This function is called from RegisterCommands in root.go.
func registerVersionCommand(root *cobra.Command) {
	root.AddCommand(newVersionCommand())
}

// newVersionCommand creates and returns the version command.
func newVersionCommand() *cobra.Command {
	var verbose bool

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of erst",
		Long:  `Display the version information for erst, including build date if available.`,
		Run: func(cmd *cobra.Command, args []string) {
			runVersion(verbose)
		},
	}

	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed version information")

	return cmd
}

// runVersion executes the version command logic.
func runVersion(verbose bool) {
	if verbose {
		fmt.Printf("erst version %s\n", Version)
		fmt.Printf("Build date: %s\n", BuildDate)
		fmt.Printf("Go version: %s\n", "go1.25.5")
	} else {
		fmt.Printf("erst version %s\n", Version)
	}
}
