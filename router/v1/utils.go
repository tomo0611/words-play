package v1

import (
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tomo0611/words-play/model"
	"github.com/tomo0611/words-play/router/consts"
)

// getRequestUser リクエストしてきたユーザーの情報を取得
func getRequestUser(c echo.Context) model.UserInfo {
	return c.Get(consts.KeyUser).(model.UserInfo)
}

// getRequestUserID リクエストしてきたユーザーUUIDを取得
func getRequestUserID(c echo.Context) uuid.UUID {
	return getRequestUser(c).GetID()
}
