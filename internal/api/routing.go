package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func HomeFunc()  {
	time.Sleep(1*time.Second)
	fmt.Println("HomeFunc")
}

func Home(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"result": "ok",
	})
}

func DummyMiddleware(c *gin.Context) {
	//fmt.Println("Im a dummy!")

	// Pass on to the next-in-chain
	c.Next()
}

func LoggerFormatter() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf( "[%s] [%s] [%s] - %s - [%d] - [%s] \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}


func NewRoutes() *gin.Engine {

	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.Use(gin.Recovery(), gin.Logger())

	router.Use(DummyMiddleware)
	router.GET("/", Home)


	return router
}
