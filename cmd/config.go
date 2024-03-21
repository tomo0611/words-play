package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type User struct {
	Name string `yaml:"name"`
}

// マッピング用の構造体
type Config struct {
	User User `yaml:user`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/")

	// 環境変数がすでに指定されてる場合はそちらを優先
	viper.AutomaticEnv()

	// データ構造をキャメルケースに切り替える用の設定
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー: %s", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s", err)
	}

	return &cfg, nil
}
