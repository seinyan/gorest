package main

import (
	"fmt"
	"github.com/seinyan/gorest/internal/api"
)



// @title Swagger API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://seinyan.ru
// @contact.email narek.seinyan@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 0.0.0.0:8000
// @BasePath /
func main() {

	config, _ := api.NewConfig()
	fmt.Println(config)

	//s := api.New()
	//if err := s.Start(); err != nil {
	//	log.Fatal()
	//}
}

