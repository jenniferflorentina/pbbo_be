package controller

import (
	"net/http"
	"strconv"
	"tubespbbo/domains/user/dto"
	"tubespbbo/domains/user/model"
	"tubespbbo/domains/user/service"
	e "tubespbbo/err"
	"tubespbbo/mapper"
	"tubespbbo/response"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"
)

func FindUser(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}
	
	users, err := service.FindUser()
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

func FindOneUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	user, err := service.FindOneUser(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.UserDTO
	mapper.Map(user, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreateUser(c *fiber.Ctx) error {
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

	err = service.CreateUser(&user)
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

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	updateDto := new(dto.UpdateUserDTO)
	err = c.BodyParser(updateDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	err = validate.Validate(&updateDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	user, err := service.UpdateUser(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.UserDTO
	mapper.Map(user, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	user, err := service.DeleteUser(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.UserDTO
	mapper.Map(user, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}