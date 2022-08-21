package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Service interface {
	Register(e *gin.Engine, db *gorm.DB)
}

type EntityURI struct {
	ID uint `uri:"id" binding:"required"`
}

func JSONResponse(c *gin.Context, body interface{}, status int) {
	c.JSON(status, body)
}

func JSONResponseOK(c *gin.Context, body interface{}) {
	JSONResponse(c, body, http.StatusOK)
}

func JSONResponseCreated(c *gin.Context, body interface{}) {
	JSONResponse(c, body, http.StatusCreated)
}

func JSONResponseNoContent(c *gin.Context) {
	JSONResponse(c, nil, http.StatusNoContent)
}

func JSONResponseError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
