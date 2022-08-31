package main

import (
	"fmt"
	"github.com/Abashinos/otus-msa-hw/app/cmd/server/views"
	"github.com/Abashinos/otus-msa-hw/app/pkg"
	"github.com/Abashinos/otus-msa-hw/app/pkg/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type App struct {
	dbConn *gorm.DB
}

func hostInfo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, NewHostInfo())
}

func (a *App) health(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, checkHealth(a.dbConn))
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

	app := App{}

	app.dbConn, err = middleware.CreateConnection()
	if err != nil {
		panic(err)
	}

	e := gin.Default()

	e.GET("/hostinfo", hostInfo)
	e.GET("/health", app.health)
	e.GET("/debug", debug)

	userService := &views.UserService{}
	userService.Register(e, app.dbConn)

	log.Infof("Starting server on %v", port)
	e.Run(fmt.Sprintf(":%v", portStr))
}
