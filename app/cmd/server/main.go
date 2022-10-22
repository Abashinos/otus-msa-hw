package main

import (
	"fmt"
	serverMiddleware "github.com/Abashinos/otus-msa-hw/app/cmd/server/middleware"
	"github.com/Abashinos/otus-msa-hw/app/cmd/server/views"
	"github.com/Abashinos/otus-msa-hw/app/pkg"
	commonMiddleware "github.com/Abashinos/otus-msa-hw/app/pkg/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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

	app.dbConn, err = commonMiddleware.CreateDBConnection()
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	prom := serverMiddleware.NewPrometheusMiddleware("gin", nil)
	prom.Use(engine)

	// This should go after Prometheus middleware, otherwise metric measurements won't work
	simulateErrorsStr := util.GetEnv("SIMULATE_ERRORS", "false")
	simulateErrors, err := strconv.ParseBool(simulateErrorsStr)
	if err != nil {
		log.Errorf("Illegal boolean value for SIMULATE_ERRORS flag: %s", simulateErrorsStr)
		simulateErrors = false
	}
	if simulateErrors {
		log.Info("Simulating errors")
		rand.Seed(time.Now().UnixNano())
		engine.Use(serverMiddleware.NewFlakyMiddleware())
		engine.Use(serverMiddleware.NewDelayerMiddleware())
	}

	engine.GET("/hostinfo", hostInfo)
	engine.GET("/health", app.health)
	engine.GET("/debug", debug)

	userService := &views.UserService{}
	userService.Register(engine, app.dbConn)

	log.Infof("Starting server on %v", port)
	engine.Run(fmt.Sprintf(":%v", portStr))
}
