package controller

import (
	"net/http"
	"tubespbbo/domains/user/dto"
	"tubespbbo/domains/user/model"
	"tubespbbo/domains/user/service"
	e "tubespbbo/err"
	"tubespbbo/mapper"
	"tubespbbo/response"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"
)

func Find(c *fiber.Ctx) error {
	users, err := service.Find()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.UserDTO
	mapper.Map(users, &DTOs)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func Create(c *fiber.Ctx) error {
	createDto := new(dto.CreateUserDTO)
	err := c.BodyParser(createDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	err = validate.Validate(&createDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var user model.User
	mapper.Map(createDto, &user)

	err = service.Create(&user)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: user,
	})
	return nil
}
