package routes

import (
	"BGP-Golang-TRPL3A/echo-rest/controllers"
	"github.com/labstack/echo"
)

// EmployeesRoute ...
func EmployeesRoute(g *echo.Group) {
	g.POST("/list", controllers.FetchAllEmployees)
	g.POST("/add", controllers.StoreEmployee)
	g.POST("/update", controllers.UpdateEmployee)
	g.POST("/delete", controllers.DeleteEmployee)
}
