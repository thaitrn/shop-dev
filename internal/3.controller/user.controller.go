package controller

import (
	service "shop-dev/internal/4.service"
	"shop-dev/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	// response.SuccessResponse(c, response.ErrorCodeSuccess, []string{"thai", "tran", "2024"})
	response.ErrorResponse(c, response.ErrorCodeEmailInvalid, nil)
}
