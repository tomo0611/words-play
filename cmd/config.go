package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/tomo0611/words-play/router"
	"github.com/tomo0611/words-play/router/auth"
)

// Config 設定
type Config struct {
	// DevMode 開発モードかどうか (default: false)
	DevMode bool `mapstructure:"dev" yaml:"dev"`

	// Origin オリジン (default: http://localhost)
	Origin string `mapstructure:"origin" yaml:"origin"`
	// Port サーバーポート番号 (default: 3000)
	Port int `mapstructure:"port" yaml:"port"`

	// MariaDB データベース接続設定
	MariaDB struct {
		// Host ホスト名 (default: 127.0.0.1)
		Host string `mapstructure:"host" yaml:"host"`
		// Port ポート番号 (default: 3306)
		Port int `mapstructure:"port" yaml:"port"`
		// Username ユーザー名 (default: root)
		Username string `mapstructure:"username" yaml:"username"`
		// Password パスワード (default: password)
		Password string `mapstructure:"password" yaml:"password"`
		// Database データベース名 (default: wordsplay)
		Database string `mapstructure:"database" yaml:"database"`
	} `mapstructure:"mariadb" yaml:"mariadb"`

	// ExternalAuth 外部認証設定
	ExternalAuth struct {
		Google struct {
			ClientID     string `mapstructure:"clientId" yaml:"clientId"`
			ClientSecret string `mapstructure:"clientSecret" yaml:"clientSecret"`
		} `mapstructure:"google" yaml:"google"`
	} `mapstructure:"externalAuth" yaml:"externalAuth"`
}

func init() {
	viper.SetDefault("dev", false)
	viper.SetDefault("origin", "http://localhost")
	viper.SetDefault("port", 3000)
	viper.SetDefault("mariadb.host", "127.0.0.1")
	viper.SetDefault("mariadb.port", 3306)
	viper.SetDefault("mariadb.username", "root")
	viper.SetDefault("mariadb.password", "password")
	viper.SetDefault("mariadb.database", "wordsplay")
	viper.SetDefault("externalAuth.google.clientId", "")
	viper.SetDefault("externalAuth.google.clientSecret", "")
}

func (c Config) getDatabase() (*sql.DB, error) {

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	dbConfig := mysql.Config{
		DBName:               config.MariaDB.Database,
		User:                 config.MariaDB.Username,
		Passwd:               config.MariaDB.Password,
		Addr:                 config.MariaDB.Host + ":" + fmt.Sprint(config.MariaDB.Port),
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true, // パスワード認証のために必須
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
	}

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return db, nil
}

func provideAuthGoogleProviderConfig(c *Config) auth.GoogleProviderConfig {
	return auth.GoogleProviderConfig{
		ClientID:     c.ExternalAuth.Google.ClientID,
		ClientSecret: c.ExternalAuth.Google.ClientSecret,
		CallbackURL:  c.Origin + "/api/auth/google/callback",
	}
}

func provideRouterExternalAuthConfig(c *Config) router.ExternalAuthConfig {
	return router.ExternalAuthConfig{
		Google: provideAuthGoogleProviderConfig(c),
	}
}
func provideRouterConfig(c *Config) *router.Config {
	return &router.Config{
		Development:  c.DevMode,
		Version:      Version,
		Revision:     Revision,
		ExternalAuth: provideRouterExternalAuthConfig(c),
	}
}
