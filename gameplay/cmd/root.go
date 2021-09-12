package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "marvel",
	Short: "CLI commands for Marvel LCG server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// The root command
	},
}

//Execute runs the root command for the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
