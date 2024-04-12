package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 上下文中设置值，后续的处理函数能读取到该值
		c.Set("name", "小王子")
		// 调用该请求剩余的处理程序
		c.Next()

		// 计算耗时
		cost := time.Since(start)
		log.Println("============ cost:", cost)
	}
}
