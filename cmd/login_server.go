package cmd

import (
	"github.com/juggler434/marvelServer/internal/login"
	"github.com/spf13/cobra"
)

var port string

func init() {
	rootCmd.AddCommand(&startLoginServer)
	startLoginServer.Flags().StringVarP(&port, "port", "", ":8080", "port to start server")
}

var startLoginServer = cobra.Command{
	Use:   "login",
	Short: "starts login server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		login.Start(port)
	},
}
