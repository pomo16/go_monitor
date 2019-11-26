package hanlder

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Ping(c *gin.Context) {
	logrus.Warn("ping")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
