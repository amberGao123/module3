package main

import (
	"context"
	"flag"
	"net/http"

	"module3/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var addr string

func main() {
	flag.StringVar(&addr, "addr", ":8080", "addr")

	r := gin.Default()

	//middleware logger
	r.Use(gin.Logger())
	//middleware recovery
	r.Use(gin.Recovery())

	//register router
	if err := handler.Register(context.Background(), r); err != nil {
		logrus.WithError(err).Errorf("handler Register")
	}

	//sever
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	logrus.Infof("start Server on %s", addr)

	//listen
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Errorf("could not listen server: %v", err)
	}

}
