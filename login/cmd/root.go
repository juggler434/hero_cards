package cmd

import (
	"fmt"
	"os"

	"github.com/juggler434/hero_cards/login/app"

	"github.com/spf13/cobra"
)

var port string

var rootCmd = cobra.Command{
	Use:   "login",
	Short: "CLI commands for login server",
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

func init() {
	rootCmd.AddCommand(start)
	start.Flags().StringVarP(&port, "port", "p", ":8080", "server port to run on")
}

var start = &cobra.Command{
	Use:   "start",
	Short: "starts the login server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		app.StartServer(port)
	},
}
