package controller

import (
	"net/http"
	_ "subscribers/docs"
	"subscribers/model"
	"subscribers/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

// CreateNewUser godoc
// @Description  Create new user
// @Tags         subscriber
// @Accept       json
// @Produce      json
// @Param user body model.User true "User data"
// @Success      200
// @Router       /api/subscriber/create [post]
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

// UpdateUser godoc
// @Description  Update user
// @Tags         subscriber
// @Accept       json
// @Produce      json
// @Param user body model.User true "User data"
// @Success      200
// @Router       /api/subscriber/update [put]
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

// GetUserById godoc
// @Summary      Show a user
// @Description  Get user by ID
// @Tags         subscriber
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200
// @Router       /api/subscriber/{id} [get]
func (s Server) GetUserById(e echo.Context) error {
	id := e.Param("id")

	user, err := repository.GetUserById(s.DB, id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if user.UserId == "" {
		return e.JSON(http.StatusNotFound, model.ErrorResponse{Status: http.StatusNotFound, Message: "User not found"})
	}

	return e.JSON(http.StatusOK, model.SuccessResponse{Status: http.StatusOK, Message: "Success", Data: user})
}

// GetUsers godoc
// @Description  Get users
// @Tags         subscriber
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/subscriber/list [get]
func (s Server) GetUsers(e echo.Context) error {
	var users []model.User

	err := repository.GetUsers(s.DB, &users)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	return e.JSON(http.StatusOK, model.SuccessResponse{Status: http.StatusOK, Message: "Success", Data: users})
}

// GetUsers godoc
// @Description  Get users
// @Tags         subscriber
// @Accept       json
// @Produce      json
// @Param        user_id path string true "User ID"
// @Param        service_name path string true "Service name"
// @Success      200
// @Router       /api/subscriber/cost [get]
func (s Server) CalculateSubsCost(e echo.Context) error {
	paramValues := e.QueryParams()

	userId := paramValues.Get("user_id")
	serviceName := paramValues.Get("service_name")

	cost, err := repository.CalculateSubsCost(s.DB, userId, serviceName)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	return e.JSON(http.StatusOK, model.SuccessResponse{Status: http.StatusOK, Message: "Success", Data: cost})
}

// DeleteUser godoc
// @Description  Delete user
// @Tags         subscriber
// @Accept       json
// @Produce      json
// @Param        user_id path string true "User ID"
// @Success      200
// @Router       /api/subscriber/delete/{id} [delete]
func (s Server) DeleteUser(e echo.Context) error {
	id := e.Param("id")

	if err := repository.DeleteUser(s.DB, id); err != nil {
		return e.JSON(http.StatusInternalServerError, model.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	return e.JSON(http.StatusOK, model.SuccessResponse{Status: http.StatusOK, Message: "Success"})
}
