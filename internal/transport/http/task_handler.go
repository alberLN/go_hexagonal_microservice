package http

import (
	"example/microservice/internal/domain"
	"example/microservice/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterHandlers(router *gin.Engine, taskService service.TaskService) {
	router.GET("/tasks/:id", getTaskById(taskService))
	router.GET("/tasks/", getAll(taskService))
	router.POST("/tasks/", createTask(taskService))
	router.PUT("/tasks/:id", updateTask(taskService))
	router.DELETE("/tasks/:id", deleteTask(taskService))
}

func getTaskById(taskService service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		order, err := taskService.GetTaskById(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

func createTask(taskService service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var taskRQ domain.TaskRequest
		err := c.BindJSON(&taskRQ)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task, err := taskService.CreateTask(c, &taskRQ)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, task)
	}
}

func updateTask(taskService service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		taskID, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var taskRQ domain.TaskRequest
		err = c.BindJSON(&taskRQ)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = taskService.UpdateTask(c, taskID, &taskRQ)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "success"})
	}
}

func deleteTask(taskService service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		taskID, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = taskService.DeleteTask(c, taskID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "success"})
	}
}

func getAll(taskService service.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderList, err := taskService.GetAll(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orderList)
	}
}
