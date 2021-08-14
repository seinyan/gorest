package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func HomeFunc()  {
	time.Sleep(5*time.Second)
	fmt.Println("HomeFunc")
}

func Home(ctx *gin.Context) {

	fmt.Println("start")
	go HomeFunc()

	select {
	case <-ctx.Request.Context().Done():
		fmt.Println("request cancelled")
		return
	}

	fmt.Println("end")

	ctx.JSON(200, gin.H{
		"result": "ok",
	})
}

func NewRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	router.GET("/", Home)

	return router
}
