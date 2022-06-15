package main

import (
	"golang-gin-poc/controller"
	"golang-gin-poc/middleware"
	"golang-gin-poc/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setUpLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	//server := gin.Default()
	setUpLogOutput()
	server := gin.New()

	server.Static("/css", "./template/css/*")
	server.LoadHTMLGlob("template/*.html")

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())
	/* 	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	}) */

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message ": "Video is saved"})
			}
		})
	}

	viewRoutes := server.Group("/views")
	{
		viewRoutes.GET("/videos", func(ctx *gin.Context) {
			videoController.ShowAll(ctx)
		})
	}

	server.Run(":8080")
}
