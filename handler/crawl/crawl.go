package crawl

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UpdateTask(c *gin.Context) {
	logrus.Info("ping")
	c.JSON(200, gin.H{
		"message": "pong3",
	})
}
