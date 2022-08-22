package main

import (
	"fmt"
	"github.com/Abashinos/otus-msa-hw/app/cmd/server/views"
	"github.com/Abashinos/otus-msa-hw/app/pkg"
	"github.com/Abashinos/otus-msa-hw/app/pkg/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func hostInfo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	fmt.Print(checkHealth())
	c.JSON(http.StatusOK, NewHostInfo())
}

func health(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, checkHealth())
}

func debug(c *gin.Context) {
	// TODO: template
	c.JSON(http.StatusOK, map[string]interface{}{"env": util.DumpEnv()})
}

func main() {
	portStr := util.GetEnv("SERVER_PORT", "8000")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Illegal port value: %s", portStr))
	}

	db, err := middleware.CreateConnection()
	if err != nil {
		panic(err)
	}

	e := gin.Default()

	e.GET("/hostinfo", hostInfo)
	e.GET("/health", health)
	e.GET("/debug", debug)

	userService := &views.UserService{}
	userService.Register(e, db)

	log.Infof("Starting server on %v", port)
	e.Run(fmt.Sprintf(":%v", portStr))
}
