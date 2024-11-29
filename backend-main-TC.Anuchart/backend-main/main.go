package main

import (
	"backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Employee Routes
	router.GET("/employees", employee.GetEmployee)         // Get all employees
	router.GET("/employees/:id", employee.GetEmployeeByID) // Get employee by ID
	router.POST("/employees", employee.PostEmployee)       // Add new employee
	router.PUT("/employees", employee.PutEmployee)         // Update employee by ID
	router.DELETE("/employees", employee.DeleteEmployee)   // Delete employee by ID

	port := ":8082" //
	router.Run(port)
}
