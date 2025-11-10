package handlers

import (
	"GoGin/internal/model"
	"GoGin/internal/services"
	"GoGin/internal/util"

	"github.com/gin-gonic/gin"

	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	//捕获数据
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Error(c, 400, err.Error())
		return
	}

	//调用服务层
	user, err := h.userService.Register(&req)
	if err != nil {
		util.Error(c, 500, err.Error())
		return
	}

	//返回响应
	util.Success(c, gin.H{
		"username": user.Username,
		"user_id":  user.UserID,
		"email":    user.Email,
	}, "RegisterRequest registered successfully")
}

func (h *UserHandler) Login(c *gin.Context) {
	//捕获数据
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "Information is incomplete",
		})
		return
	}

	//调用服务层
	token, user, err := h.userService.Login(req.LoginKey, req.Password)
	if err != nil {
		util.Error(c, 500, err.Error())
	}

	//返回响应
	util.Success(c, gin.H{
		"username": user.Username,
		"user_id":  user.UserID,
		"email":    user.Email,
		"token":    token,
	}, "login successful")
}

func (h *UserHandler) InfoHandler(c *gin.Context) {
	//捕获数据
	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")

	//返回响应
	util.Success(c, gin.H{
		"user_id":  userID,
		"username": username,
	}, "Protected resource")
}
