/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/
package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/tomo0611/words-play/router"
)

// サーバー起動コマンド
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve Words Play API",
	Run: func(cmd *cobra.Command, args []string) {
		// Logger
		setLogger()
		slog.Info(fmt.Sprintf("Words-Play %s (revision %s)", Version, Revision))
		// Database
		slog.Info("connecting database...")
		db, err := config.getDatabase()
		if err != nil {
			slog.Error(err.Error())
		}
		defer db.Close()

		slog.Info("serve called")

		// DBでほかのクラスに渡す
		server := router.SetupRouter(&config, db)

		go func() {
			if err := server.Start(fmt.Sprintf(":%d", config.Port)); err != nil {
				slog.Info("shutting down the server")
			}
		}()
		slog.Info("server started")

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
