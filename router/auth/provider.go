package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/tomo0611/words-play/utils/random"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const (
	cookieName   = "tomo_oauth_state"
	cookieMaxAge = 60 * 5 /* 5分間 */
)

type Provider interface {
	FetchUserInfo(t *oauth2.Token) (UserInfo, error)
	LoginHandler(c echo.Context) error
	CallbackHandler(c echo.Context) error
	L() *zap.Logger
}

type UserInfo interface {
	GetProviderName() string
	GetID() string
	GetRawName() string
	GetName() string
	GetDisplayName() string
	GetProfileImage() ([]byte, error)
}

func defaultLoginHandler(oac *oauth2.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.Request().Header.Get(echo.HeaderAuthorization)) > 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Authorization Header must not be set.")
		}

		state := random.SecureAlphaNumeric(32)
		c.SetCookie(&http.Cookie{
			Name:     cookieName,
			Value:    state,
			Path:     "/",
			Expires:  time.Now().Add(cookieMaxAge * time.Second),
			MaxAge:   cookieMaxAge,
			HttpOnly: true,
		})
		return c.Redirect(http.StatusFound, oac.AuthCodeURL(state))
	}
}

func defaultCallbackHandler(p Provider, oac *oauth2.Config /*repo repository.Repository,*/, sessStore sessions.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.Request().Header.Get(echo.HeaderAuthorization)) > 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Authorization Header must not be set.")
		}

		code := c.QueryParam("code")
		state := c.QueryParam("state")
		if len(code) == 0 || len(state) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "missing code or state")
		}

		cookie, err := c.Cookie(cookieName)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "missing state cookie")
		}
		if cookie.Value != state {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid state")
		}

		t, err := oac.Exchange(context.Background(), code)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to exchange code")
		}

		userinfo, err := p.FetchUserInfo(t)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to fetch user info")
		}

		fmt.Println(userinfo)

		// ユーザー追加などの処理をここでする

		return c.Redirect(http.StatusFound, "/")
	}
}
