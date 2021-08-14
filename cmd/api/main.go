package main

import (
	"github.com/seinyan/gorest/config"
	"github.com/seinyan/gorest/internal/api"
	"github.com/seinyan/gorest/pkg/logging"
	"os"
	"os/signal"
	"syscall"
	"time"
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


	//log := logrus.New()
	//log.SetLevel(logrus.InfoLevel)
	////log.SetReportCaller(true)
	//log.Formatter = &logrus.TextFormatter{
	//	//DisableColors: true,
	//	FullTimestamp: true,
	//}
	//
	//log.Error("Err: ", errors.New("dsada"))

	logger := logging.New()
	logger.Error("EEE")


	/*	 */

	conf, err := config.New()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info(*conf)

	router := api.NewRoutes()
	srv := api.NewServer(conf)
	go func() {
		if err := srv.Run(router); err != nil {
			logger.Error("http server run error: ", err.Error())
		}
	}()


	// Handling OS signal

	logger.Info("Listen os signal ...")

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("Shut down please wait ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Shut down Error: %s", err.Error())
	}

	logger.Info("Shut down success")



	//if err := db.Close(); err != nil {
	//	logrus.Errorf("error occured on db connection close: %s", err.Error())
	//}

	//db, err := database.NewDBConn(conf.DatabaseUrl)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//s := api.New(conf)
	//if err := s.Start(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//DBconn, _ := db.DB()
	//defer DBconn.Close()
}

