package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/service/database"
	"testing"
)

func TestDB(t *testing.T) {
	database.InitDB()
	res, err := database.ListTask(&gin.Context{})
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
