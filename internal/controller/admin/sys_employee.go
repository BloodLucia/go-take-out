package adminctrl

import "github.com/gofiber/fiber/v2"

type employeeController struct {
}

type EmployeeController interface {
	AddEmployee(ctx *fiber.Ctx) error
}

// AddEmployee 添加一名员工.
func (*employeeController) AddEmployee(ctx *fiber.Ctx) error {
	return nil
}

func NewEmployeeController() EmployeeController {
	return &employeeController{}
}
