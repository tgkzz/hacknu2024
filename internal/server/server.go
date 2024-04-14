package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func StartServer(host, port string, r *gin.Engine) error {
	srv := http.Server{
		Addr:         host + ":" + port,
		Handler:      r,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 2 * time.Minute,
	}

	log.Print(fmt.Sprintf("start server on http://%s:%s", host, port))

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
