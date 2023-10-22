package handler

import (
	"bwastartup/helper"
	"bwastartup/user"

	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) ResgisterUser(c *gin.Context) {
	var input user.ResgisterUserInput
	err := c.ShouldBindJSON(&input)

errors := helper.FormatValidationError(err)

	errorMessage := gin.H{"errors": errors}

	if err != nil {
		response := helper.ApiResponse("Register Account Failed", http.StatusUnprocessableEntity, "error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	
	if err != nil {
		response := helper.ApiResponse("Register Account Failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	} 

	formatter := user.UserFormat(newUser, "oprjwihrowdkofnk")

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}