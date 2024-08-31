package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var reassembleCmd = &cobra.Command{
	Use:   "reassemble",
	Short: "Reassemble data from DNS lookups stored",
	Long:  `For a given exfil ID, reassemble the data from the DNS lookups stored in the SQLite database.`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func init() {
	rootCmd.AddCommand(reassembleCmd)
}

func execute() {
	fmt.Println("reassemble called")
}
