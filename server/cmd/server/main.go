package main

import (
	"encoding/json"
	"fmt"
	"github.com/Abashinos/otus-msa-hw/server/cmd/server/views"
	"github.com/Abashinos/otus-msa-hw/server/pkg"
	"github.com/Abashinos/otus-msa-hw/server/pkg/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func hostInfo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	fmt.Print(checkHealth())
	response, _ := json.Marshal(NewHostInfo())
	c.JSON(http.StatusOK, string(response))
}

func health(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	response, _ := json.Marshal(checkHealth())
	c.JSON(http.StatusOK, string(response))
}

func debug(c *gin.Context) {
	// TODO: template
	data := struct {
		Env map[string]string `json:"env"`
	}{
		Env: util.DumpEnv(),
	}
	response, _ := json.Marshal(data)
	c.JSON(http.StatusOK, string(response))
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

	log.Printf("Starting server on %v", port)
	e.Run(portStr)
}
