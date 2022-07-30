package user

import (
	"RestAPI/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(usersURL, h.GetList)
	router.POST(userURL, h.CreateUser)
	router.GET(userURL, h.GetUserByID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

}

func (h *handler) GetList(c *gin.Context) {
	_, err := c.Writer.Write([]byte("this is list of users"))
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) CreateUser(c *gin.Context) {
	_, err := c.Writer.Write([]byte("this is create user"))
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) GetUserByID(c *gin.Context) {
	_, err := c.Writer.Write([]byte("this is get user by id"))
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) UpdateUser(c *gin.Context) {
	_, err := c.Writer.Write([]byte("this is update user"))
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) PartiallyUpdateUser(c *gin.Context) {
	_, err := c.Writer.Write([]byte("this is partially update user"))
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) DeleteUser(c *gin.Context) {
	_, err := c.Writer.Write([]byte("this is delete user"))
	if err != nil {
		log.Println(err)
	}
}
