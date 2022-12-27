package server

import (
	"fmt"
	"net/http"

	"github.com/Killayt/Password-generator/config"
	"github.com/gin-gonic/gin"
)

func rend(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, message)
}

func pingHandler(c *gin.Context) {
	rend(c.Writer, "PONG")
}

func indexHandler(c *gin.Context) {
	rend(c.Writer, "index")
}

func Start() error {
	r := gin.Default()

	r.GET("/ping", pingHandler)
	r.GET("/", indexHandler)

	err := r.Run(":" + config.Init("config/.env"))

	return err
}
