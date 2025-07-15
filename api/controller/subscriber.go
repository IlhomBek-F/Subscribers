package controller

import (
	"net/http"
	"subscribers/model"
	"subscribers/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (s Server) CreateNewUser(e echo.Context) error {
	var user model.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, model.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if err := repository.CreateNewUser(s.DB, user); err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	return e.JSON(http.StatusCreated, model.SuccessResponse{Status: http.StatusCreated, Message: "Success"})
}

func (s Server) UpdateUser(e echo.Context) error {
	var user model.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, model.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if err := repository.UpdateUser(s.DB, user); err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	return e.JSON(http.StatusOK, model.SuccessResponse{Status: http.StatusOK, Message: "Success"})
}

func (s Server) GetUserById(e echo.Context) error {
	id := e.Param("id")

	user, err := repository.GetUserById(s.DB, id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if user.ID == "" {
		return e.JSON(http.StatusNotFound, model.ErrorResponse{Status: http.StatusNotFound, Message: "User not found"})
	}

	return e.JSON(http.StatusOK, model.SuccessResponse{Status: http.StatusOK, Message: "Success", Data: user})
}

func (s Server) GetUsers(e echo.Context) error {
	var users []model.User

	err := repository.GetUsers(s.DB, &users)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	return e.JSON(http.StatusOK, model.SuccessResponse{Status: http.StatusOK, Message: "Success", Data: users})
}

func (s Server) DeleteUser(e echo.Context) error {
	id := e.Param("id")

	if err := repository.DeleteUser(s.DB, id); err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	return e.JSON(http.StatusOK, model.SuccessResponse{Status: http.StatusOK, Message: "Success"})
}
