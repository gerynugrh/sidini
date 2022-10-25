package message

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"net/http"
)

var usecase UseCase = NewUseCase()

func RegisterRoute(router *echo.Echo) error {
	router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/forum/get_list",
		Handler: func(c echo.Context) error {
			param := GetListParam{}
			if err := c.Bind(&param); err != nil {
				return fmt.Errorf("c.Bind: %w", err)
			}
			messages, err := usecase.GetList(c.Request().Context(), param)
			if err != nil {
				return fmt.Errorf("usecase.GetList: %w", err)
			}
			return c.JSON(http.StatusOK, messages)
		},
		Middlewares: []echo.MiddlewareFunc{
			apis.RequireUserAuth(),
		},
	})

	return nil
}
