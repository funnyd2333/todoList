package controller

import (
	"bubble/global"
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddTodo 创建记录
func AddTodo(c *gin.Context) {
	//前端页面填写待办事项 点击提交 发送数据到这里
	//1.取数据
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := global.DB.AutoMigrate(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}

func ShowTodos(c *gin.Context) {
	var todos []models.Todo
	if err := global.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

// UpdateTodo 修改
func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效id"})
		return
	}
	//先查后保存修改
	var todo models.Todo
	if err := global.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}

	if err := global.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效id"})
		return
	}
	if err := global.DB.Where("id = ?", id).First(&models.Todo{}).Delete(&models.Todo{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	}
}
