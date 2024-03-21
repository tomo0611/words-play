package v1

import (
	"context"
	"fmt"
	"net/http"
	"os"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/tomo0611/words-play/database"
)

// GetWordsRequest GET /words リクエストボディ
type GetWordsRequest struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

func (r *GetWordsRequest) Validate() error {
	if r.Limit == 0 {
		r.Limit = 50
	}
	return vd.ValidateStruct(r,
		vd.Field(&r.Limit, vd.Required, vd.Min(1), vd.Max(50)),
		vd.Field(&r.Offset, vd.Required, vd.Min(0), vd.Max(10000)),
	)
}

// GetWordsRequest GET /words
func (h *Handlers) GetWords(c echo.Context) error {
	// middlewareかなんかを使ってユーザー情報を取得?? よくわかってない
	// userID := getRequestUserID(c)

	var req GetWordsRequest
	// parse query parameters
	if err := c.Bind(&req); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	// Validate request
	err := req.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// repositoryからデータを取得
	ctx := context.Background()

	db, err := config.getDatabase()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	queries := database.New(db)

	// list words
	words, err := queries.ListWords(ctx, database.ListWordsParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Offset),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// response
	type response struct {
		Code    int             `json:"code"`
		message string          `json:"message"`
		Words   []database.Word `json:"words"`
	}

	res := response{
		Code:    200,
		message: "OK",
		Words:   words,
	}

	return c.JSON(http.StatusOK, res)
}
