package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps"
)

func main() {
	r := gin.Default()

	app := apps.NewApplication(r)
	app.Create()

	r.Run()
}
