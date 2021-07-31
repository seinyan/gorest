package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/seinyan/gorest/docs"
	"github.com/seinyan/gorest/internal/database"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Start() error
}

type server struct {
	Config *Config
}

func setupLogOutput()  {
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func (s server) Start() error {
	//setupLogOutput()




	db, _ := database.NewDBConn()

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	server.GET("/users", func(context *gin.Context) {
		fmt.Println("dd")
	})


	//use ginSwagger middleware to serve the API docs
	//http://127.0.0.1:9000/swagger/index.html
	//https://github.com/swaggo/swag/tree/master/example
	server.GET("/", func(ctx *gin.Context) {
		//Redirect to docs page
		ctx.Redirect(301,  "/swagger/index.html")
	})
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run("0.0.0.0:8000")

	DBconn, _ := db.DB()
	defer DBconn.Close()

	return nil
}



func New() Server {
	return &server{}
}