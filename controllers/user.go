package controllers

import (
	"net/http"
	"server/models"
	"server/requests"
	"server/responses"
	services "server/services/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.Service
}

func NewUserController(service services.Service) *UserController {
	return &UserController{service}
}

func convertUserResponse(o models.User) responses.UserResponse {
	return responses.UserResponse{
		ID:          o.ID,
		Name:        o.Name,
		Username:    o.Username,
		Email:       o.Email,
		Address:     o.Address,
		PhoneNumber: o.PhoneNumber,
	}
}

func convertLoginResponse(o models.User, token string) responses.UserLoginResponse {
	return responses.UserLoginResponse{
		ID:       o.ID,
		Name:     o.Name,
		Username: o.Username,
		UserRole: o.UserRole.Role,
		Token:    token,
	}
}

func (h *UserController) FindAllUsers(c *gin.Context) {
	users, err := h.service.FindAll()
	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err.Error(),
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	var userResponses []responses.UserResponse

	if len(users) == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   []responses.UserResponse{},
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	for _, user := range users {
		response := convertUserResponse(user)

		userResponses = append(userResponses, response)
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   userResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *UserController) FindUserByID(c *gin.Context) {
	idParam := c.Param("id")
	ID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	user, err := h.service.FindByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// If no user is found, return null
	if user.ID == 0 {
		webResponse := responses.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		}
		c.JSON(http.StatusOK, webResponse)
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertUserResponse(user),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *UserController) CreateUser(c *gin.Context) {
	var userForm requests.CreateUserRequest

	err := c.ShouldBindJSON(&userForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.Create(userForm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertUserResponse(user),
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *UserController) LoginUser(c *gin.Context) {
	var userForm requests.LoginUserRequest

	err := c.ShouldBindJSON(&userForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, token, err := h.service.Login(userForm)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   convertLoginResponse(user, token),
	}

	c.JSON(http.StatusOK, webResponse)
}
