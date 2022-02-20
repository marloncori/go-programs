package employee

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var database *gorm.database
var exception error
const DNS = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8md38"

type Employee struct {
	gorm.model
	FistName string `json:"first_name"`
	LastName string `json: "last_name"` 
	Email string `json:"email"`
}

func InitialMigration(){
	database, exception = gorm.Open(mysql.Open(DNS), &gorm.Config())
	if exception != nil {
		fmt.Println(exception.Error())
		panic("We cannot connect to the database. Check whether it was really created.")
	}
}

func GetAllEmployees (ctxt *fiber.Ctx) error {
	var employees []Employee
	database.Find(&users)
	return ctxt.JSON(&users)
}

func GetEmployee (ctxt *fiber.Ctx) error {
	id := ctxt.Params("id")
	var empl Employee
	database.First(&empl, id)
	return ctxt.JSON(&empl)
}


func AddEmployee (ctxt *fiber.Ctx) error {
	empl = new(Employee)
	if err := ctxt.BodyParser(empl); err != nil {
		return ctxt.Status(500).SendString(err.Error())
	}

	database.Save(&empl)
	return ctxt.JSON(&empl)
}


func UpdateEmployee (ctxt *fiber.Ctx) error {
	id := ctxt.Params("id")
	empl = new(Employee)
	database.First(&empl, id)
	if empl.Email == "" {
		return ctxt.Status(500).SendString("User has not been found in the DB.")
	}

	if err := ctxt.BodyParser; err != nil {
		return ctxt.Status(500).SendString(err.Error())
	}

	database.Save(&empl)
	return ctxt.JSON(&empl)
}


func DeleteEmployee (ctxt *fiber.Ctx) error {
	id := ctxt.Params("id")
	var empl Employee
	database.First(&empl, id)
	if empl.Email == "" {
		return ctxt.Status(500).SendString("User not availabe in the DB.")
	}

	database.Delete(&empl)
	return ctxt.SendString("User has been successfully deleted.")
}