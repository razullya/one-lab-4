package handler

import (
	"encoding/json"
	"io/ioutil"
	"lab4/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) getUserById(c echo.Context) error {
	id := c.Param("id")
	user, err := h.Services.User.GetUserById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
func (h *Handler) updateUser(c echo.Context) error {
	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	user := models.User{}
	err = json.Unmarshal(data, &user)

	if err != nil {
		return err
	}
	if err := h.Services.User.UpdateUser(user.Name, user.Surname, user.Id); err != nil {
		return err
	}
	return c.String(http.StatusOK, "user saved")
}
func (h *Handler) createUser(c echo.Context) error {
	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	user := models.User{}
	err = json.Unmarshal(data, &user)

	if err != nil {
		return err
	}

	if err := h.Services.User.CreateUser(user.Name, user.Surname); err != nil {
		log.Fatal(err)
		return err
	}

	return c.String(http.StatusOK, "user updated")
}
func (h *Handler) deleteUser(c echo.Context) error {
	if err := h.Services.User.DeleteUser(c.Param("id")); err != nil {
		return err
	}
	return c.String(http.StatusOK, "user delete by ID")
}
