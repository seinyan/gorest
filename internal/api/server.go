package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/seinyan/gorest/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Start()
}

type server struct {}

func setupLogOutput()  {
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}


// Get godoc
// @Summary User Get
// @Description User
// @Tags Security
// @Accept json
// @Produce json
// @Param id path uint64 true "Id"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "Not fount"
// @Failure 401 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /user [get]
// @Security bearerAuth
func GetUser(ctx *gin.Context) {


}

func (s server) Start() {
	//setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	server.GET("/user", GetUser)

	//Redirect to docs page
	server.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(301,  "/swagger/index.html")
	})
	//use ginSwagger middleware to serve the API docs
	//http://127.0.0.1:9000/swagger/index.html
	//https://github.com/swaggo/swag/tree/master/example
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run("0.0.0.0:8000")
}


func NewServer() Server {
	return &server{}
}