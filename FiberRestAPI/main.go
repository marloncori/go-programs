package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marloncori/go-programs/FiberRestAPI/employee"
)


func initRouters(application *fiber.App){
	application.Get("/employees", employee.GetAllEmployees)
	application.Get("/employee/:id", employee.GetEmployee)
	application.Post("/employee", employee.AddEmployee)
	application.Put("/employee/:id", employee.UpdateEmployee)
	application.Delete("/employee/:id", employee.DeleteEmployee)
}

func Hello(context *fiber.Ctx){
	context.SendString("Welcome to a Fiber-based Application written in Golang!")
}

func main(){
	app := fiber.New()
	app.Get("/", Hello)

	initRouters(app)	
	app.Listen(":3000")
}