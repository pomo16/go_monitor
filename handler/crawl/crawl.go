package crawl

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AddTask(c *gin.Context) {
	logrus.Info("ping")
	c.JSON(200, gin.H{
		"message": "pong1",
	})
}

func GetTask(c *gin.Context) {
	logrus.Info("ping")
	c.JSON(200, gin.H{
		"message": "pong2",
	})
}

func UpdateTask(c *gin.Context) {
	logrus.Info("ping")
	c.JSON(200, gin.H{
		"message": "pong3",
	})
}
