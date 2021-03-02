package main

import (
	"github.com/gin-gonic/gin"
	"matching-engine/config"
	"matching-engine/engine"
	"matching-engine/errcode"
	"matching-engine/handler"
	"matching-engine/log"
	"matching-engine/middleware"
	"matching-engine/process"
)

func init() {
	config.Init(config.TestPath)
	log.Init()
	errcode.Init(config.Conf.ErrCodes)
	engine.Init()
	middleware.Init()
	process.Init()
}

func main() {
	e := gin.Default()
	e.POST("/open_matching", handler.OpenMatching())
	e.POST("/close_matching", handler.CloseMatching())
	e.POST("/handle_order", handler.HandleOrder())
	if err := e.Run(":8989"); err != nil {
		panic(err)
	}
}
