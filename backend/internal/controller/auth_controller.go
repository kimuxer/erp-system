package controller

import (
	"net/http"

	"github.com/kimuxer/erp-system/pkg/database"

	"github.com/kimuxer/erp-system/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user struct {
		Username string `json:"username" binding:"required,min=3,max=50"`
		Password string `json:"password" binding:"required,min=8"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码处理失败"})
		return
	}

	newUser := model.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Email:    user.Email,
	}

	if result := database.DB.Create(&newUser); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "用户名或邮箱已存在"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       newUser.ID,
		"username": newUser.Username,
		"email":    newUser.Email,
	})
}
