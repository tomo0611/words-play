package router

import "github.com/tomo0611/words-play/router/auth"

type Config struct {
	Development  bool
	Version      string
	Revision     string
	ExternalAuth ExternalAuthConfig
}

// ExternalAuthConfig 外部認証設定
type ExternalAuthConfig struct {
	// ExternalAuthConfig 外部認証設定
	Google auth.GoogleProviderConfig
}

func (c ExternalAuthConfig) ValidProviders() map[string]bool {
	res := make(map[string]bool)
	if c.Google.Valid() {
		res[auth.GoogleProviderName] = true
	}
	return res
}
