/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	Version  string
	Revision string
)

var (
	// configFile 設定ファイルyamlのパス
	configFile string
	// c 設定
	config Config
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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 環境変数がすでに指定されてる場合はそちらを優先
	viper.AutomaticEnv()

	// データ構造をキャメルケースに切り替える用の設定
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("設定ファイル読み込みエラー: %s", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unmarshal error: %s", err)
		os.Exit(1)
	}

	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&configFile, "config", "c", "", "config file path")
	flags.Bool("dev", false, "development mode")
	bindPFlag(flags, "dev")

	rootCmd.MarkFlagRequired("config")
}

func bindPFlag(flags *pflag.FlagSet, key string, flag ...string) {
	if len(flag) == 0 {
		flag = []string{key}
	}
	if err := viper.BindPFlag(key, flags.Lookup(flag[0])); err != nil {
		panic(err)
	}
}

func setLogger() {
	if config.DevMode {
		logger := getCLILogger()
		slog.SetDefault(logger)
	} else {
		logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
		slog.SetDefault(logger)
	}
}

func getCLILogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func waitSIGINT() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	signal.Stop(quit)
	close(quit)
	for range quit {
		continue
	}
}
