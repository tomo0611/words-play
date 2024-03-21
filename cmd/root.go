/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Version  string
	Revision string
)

var (
	// configFile 設定ファイルyamlのパス
	configFile string
	// c 設定
	// c Config
)

var rootCmd = &cobra.Command{
	Use:   "words-play",
	Short: "Words PlayはCLIとHTTP上での英単語学習アプリです。",
	Long:  `Words PlayはCLIとHTTP上での英単語学習アプリです。`,
}

// Execute関数はすべての子コマンドを追加してフラグを適切にセットする。
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config.yml", "config file (default is config.yml)")

	// Cobra also supports local flags, which will only run when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
