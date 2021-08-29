package main

import (
	"context"
	"fmt"

	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/TeamChii/hello-lambda/hello"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func init() {
	runtime.GOMAXPROCS(1)

}

func main() {
	var (
		logger, _ = zap.NewProduction()
		e         = initEcho()
	)

	defer logger.Sync()

	server := Server{
		logger: logger,
	}
	server.Mount(e)

	isLambda := os.Getenv("LAMBDA")

	if isLambda == "TRUE" {
		lambdaAdapter := &LambdaAdapter{Echo: e}
		lambda.Start(lambdaAdapter.Handler)
	} else {
		logger.Info(fmt.Sprintf("Listening on port:%s", "1323"))
		go func() {
			logger.Info(fmt.Sprint(e.Start(":1323")))
		}()
	}
	gracefulShutdown(e)
}

type Server struct {
	logger *zap.Logger
}

func (s *Server) Mount(e *echo.Echo) {
	handler := hello.NewHandler(
		hello.NewService(s.logger),
	)
	e.POST("/hello", handler.HelloHandler)

}

func gracefulShutdown(e *echo.Echo) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	if err := e.Shutdown(context.Background()); err != nil {
		log.Fatal("shutdown server:", err)
	}
}

func initEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	return e
}
