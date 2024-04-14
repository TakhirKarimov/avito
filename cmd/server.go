package cmd

import (
	_ "avito/init"
	"avito/pkg/di"
	"avito/servers"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var webCmd = &cobra.Command{
	Use: "web",
	Run: func(cmd *cobra.Command, args []string) {
		s := di.Get("servers").(*servers.ServerManager)
		if err := s.Run(); err != nil {
			log.Fatal("failed to start server")
		}
		time.Sleep(1 * time.Hour)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
