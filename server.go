package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()

}

func serveFrontend(path string) {
	router.Use(static.Serve("/", static.LocalFile(path, false)))
}
