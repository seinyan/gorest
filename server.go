package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/seinyan/gorest/config"
	_ "github.com/seinyan/gorest/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Start() error
}

type server struct {
	Config *config.Config
}

func setupLogOutput()  {
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func (s server) Start() error {
	//setupLogOutput()

	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())
	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"users": "users",
		})
	})

	//use ginSwagger middleware to serve the API docs
	//http://127.0.0.1:9000/swagger/index.html
	//https://github.com/swaggo/swag/tree/master/example
	router.GET("/", func(ctx *gin.Context) {
		//Redirect to docs page
		ctx.Redirect(301,  "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(s.Config.ServerAddr)

	return nil
}

func New(conf *config.Config) Server {
	return &server{
		Config: conf,
	}
}