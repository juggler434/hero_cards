package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = cobra.Command{
	Use: "marvel",
	Short: "CLI commands for Marvel LCG server",
	Long: "",
	Run: func(cmd *cobra.Command, args []string)  {
		// The root command
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
