package views

import (
	"github.com/Abashinos/otus-msa-hw/app/pkg/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type UserService struct {
	repository *models.UserRepository
}

func (u *UserService) create(c *gin.Context) {
	var userPayload models.User
	err := c.ShouldBindJSON(&userPayload)
	if err != nil {
		log.Errorf("Unable to decode the request body into User model. %v", err)
		JSONResponseError(c, err)
		return
	}

	user, err := u.repository.Create(&userPayload)
	if err != nil {
		log.Errorf("Unable to create User instance. %v", err)
		JSONResponseError(c, err)
		return
	}

	JSONResponseCreated(c, gin.H{"entity": user})
}

func (u *UserService) list(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Oops")
}

func (u *UserService) get(c *gin.Context) {
	entityData := EntityURI{}
	c.BindUri(&entityData)
	user, err := u.repository.Get(entityData.ID)
	if err != nil {
		JSONResponseError(c, err)
		return
	}
	JSONResponseOK(c, gin.H{"entity": user})
}

func (u *UserService) update(c *gin.Context) {
	entityData := EntityURI{}
	c.BindUri(&entityData)

	var userPayload models.User
	err := c.ShouldBindJSON(&userPayload)
	if err != nil {
		log.Errorf("Unable to decode the request body into User model. %v", err)
		JSONResponseError(c, err)
		return
	}

	var user *models.User
	if user, err = u.repository.Update(entityData.ID, &userPayload); err != nil {
		JSONResponseError(c, err)
	}
	JSONResponseOK(c, gin.H{"entity": user})
}

func (u *UserService) delete(c *gin.Context) {
	entityData := EntityURI{}
	c.BindUri(&entityData)
	if err := u.repository.Delete(entityData.ID); err != nil {
		JSONResponseError(c, err)
	}
	JSONResponseNoContent(c)
}

func (u *UserService) Register(e *gin.Engine, db *gorm.DB) {
	e.GET("/users", u.list)
	e.POST("/users", u.create)
	e.GET("/users/:id", u.get)
	e.PUT("/users/:id", u.update)
	e.DELETE("/users/:id", u.delete)

	u.repository = models.NewUserRepository(db)
}
