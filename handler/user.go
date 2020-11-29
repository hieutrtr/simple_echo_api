package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

type User struct {
	UserName string `json:"username"`
	Age int32 `json:"age"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func (h *Handler) RegisterUserHandler(g *echo.Group) {
	g.POST("", h.registerUser)
	g.PUT("", h.updateUser)
	g.GET("", h.getUser)
	g.DELETE("", h.deleteUser)
}

func (h *Handler) registerUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, "user data is invalid")
	}
	fmt.Println(user)
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "update user successfully")
}

func (h *Handler) getUser(c echo.Context) error {
	return c.String(http.StatusOK, "get user successfully")
}

func (h *Handler) deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "delete user successfully")
}
