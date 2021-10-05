package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/libs/cli"
)

var RollbackStateCmd = &cobra.Command{
	Use:   "rollback",
	Short: "rollback tendermint state by one height",
	Long: `
A state rollback is performed to recover from an incorrect application state transition,
when Tendermint has persisted an incorrect app hash and is thus unable to make
progress. Rollback overwrites a state at height n with the state at height n - 1.
The application should also roll back to height n - 1. No blocks are removed, so upon
restarting Tendermint the transactions in block n will be re-executed against the
application.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		height, hash, err := cli.RollbackState(config)
		if err != nil {
			return fmt.Errorf("failed to rollback state: %w", err)
		}

		fmt.Printf("Rolled back state to height %d and hash %v", height, hash)
		return nil
	},
}
