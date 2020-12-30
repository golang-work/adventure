package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        requestId := ctx.Request.Header.Get("At-Request-Id")
        if requestId == "" {
            requestId = uuid.New().String()
        }

        ctx.Header("At-Request-Id", requestId)
        ctx.Next()
    }
}

func GetRequestId(c *gin.Context) string {
    return c.Writer.Header().Get("At-Request-Id")
}
