package cmd

import (
	_ "avito/init"
	"avito/pkg/di"
	"avito/servers"
	"github.com/spf13/cobra"
	"log"
	"sync"
)

var webCmd = &cobra.Command{
	Use: "web",
	Run: func(cmd *cobra.Command, args []string) {
		wg := sync.WaitGroup{}
		wg.Add(1)
		s := di.Get("servers").(*servers.ServerManager)
		if err := s.Run(); err != nil {
			log.Fatal("failed to start server")
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
