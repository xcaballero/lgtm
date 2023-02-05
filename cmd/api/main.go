package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/afiskon/promtail-client/promtail"
	"github.com/gin-gonic/gin"
	"github.com/xcaballero/lgtm/pkg/net/http/handler"
	"github.com/xcaballero/lgtm/pkg/net/http/router"
)

func main() {
	labels := "{source=\"go\",job=\"go\"}"
	conf := promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push",
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}
	loki, err := promtail.NewClientJson(conf)
	if err != nil {
		log.Printf("promtail.NewClient: %s\n", err)
		os.Exit(1)
	}

	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(router.LogMiddleware(loki))

	h := handler.New(loki)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router.New(ginEngine, h),
	}

	loki.Infof("Server starting at time=%s\n", time.Now().String())
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

// func main() {
// 	labels := "{source=\"go\",job=\"go\"}"
// 	conf := promtail.ClientConfig{
// 		PushURL:            "http://localhost:3100/api/prom/push",
// 		Labels:             labels,
// 		BatchWait:          5 * time.Second,
// 		BatchEntriesNumber: 10000,
// 		SendLevel:          promtail.INFO,
// 		PrintLevel:         promtail.ERROR,
// 	}
// 	loki, err := promtail.NewClientJson(conf)
// 	if err != nil {
// 		log.Printf("promtail.NewClient: %s\n", err)
// 		os.Exit(1)
// 	}

// 	loki.Errorf("server starting at time=%s\n", time.Now().String())

// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		loki.Infof("GET /ping at time=%s\n", time.Now().String())
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run()
// }
