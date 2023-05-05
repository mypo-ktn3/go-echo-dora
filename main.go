package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echoインスタンスを作成
	e := echo.New()

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// サーバーを開始
	e.Logger.Fatal(e.Start(":8080"))
}
