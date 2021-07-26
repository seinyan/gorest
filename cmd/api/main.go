package main

import "github.com/seinyan/gorest/internal/api"

// @title Swagger Example API TEST
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host petstore.swagger.io:8080
// @BasePath /v2
func main() {
	s := api.NewServer()
	s.Start()
}

