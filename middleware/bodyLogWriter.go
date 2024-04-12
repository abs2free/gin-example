package middleware

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter               // 嵌入的gin框架ResponseWriter
	body               *bytes.Buffer // 用于记录response
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // 记录一份
	return w.ResponseWriter.Write(b) // 真正写入响应
}

func GinBodyMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}

	c.Writer = blw // 使用我们自定义的类型替换默认的
	c.Next()       // 执行业务逻辑

	fmt.Println("Response body:" + blw.body.String()) // 事后按照需求记录返回的响应
}
