package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tao-yi/golang-react-template/model"
)

type User struct{}

// FindOne godoc
// @Summary Show a user
// @Description get user by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} http.HTTPError
// @Failure 500 {object} http.HTTPError
// @Failure default {object} http.DefaultError
// @Router /users/{id} [get]
func (*User) FindOne(ctx *gin.Context) {
	ctx.JSON(200, &model.User{Name: "Test User"})
}
