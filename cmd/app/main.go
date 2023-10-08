package main

import (
	"rest_api/internal/app/web"
)

func main() {
	app := web.CreateApp()
	err := app.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}
