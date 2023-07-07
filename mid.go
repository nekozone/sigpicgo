package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Cktoken_get(c *gin.Context) {
	qm := c.Query("token")
	if qm == config.Token {
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "token error"})
		c.Abort()
	}
}

func Numlim(rds *redis.Client) bool {
	r, err := rds.Get(ctx, "numlim").Result()
	if err == redis.Nil {
		rds.Set(ctx, "numlim", 1, time.Hour*24)
		return true
	} else if err != nil {
		panic(err)
	} else {
		rint, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}
		if rint > config.Maxnum {
			return false
		} else {
			rds.Incr(ctx, "numlim")
			return true
		}
	}
}
