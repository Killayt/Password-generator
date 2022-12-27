package main

import (
	"log"

	"github.com/Killayt/Password-generator/pkg/server"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
