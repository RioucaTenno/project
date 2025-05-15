package middleware

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(c.Writer.Status(), ErrorResponse{
				Code:    c.Writer.Status(),
				Message: c.Errors[0].Error(),
			})
		}
	}
}
