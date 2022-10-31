package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// addr := e.Listener.Addr()
	// fmt.Println(addr.String())
	e.GET("/", home)
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUserById)
	e.GET("/users/:name", getUserByName)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello mate")
}
func getUserById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func getUserByName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}
func saveUser(c echo.Context) error {
	// psevdo code
	return c.String(http.StatusOK, "user saved")
}
func updateUser(c echo.Context) error {
	//	updating user
	return c.String(http.StatusOK, "user updated")
}
func deleteUser(c echo.Context) error {
	//	delete user by ID
	return c.String(http.StatusOK, "user delete by ID")
}
