package middleware

import (
	"net/http"

	"github.com/aungkoko1234/tickermaster_backend/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware () gin.HandlerFunc {
	return func(ctx *gin.Context) {
	    err := token.TokenValid(ctx)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect."})
			ctx.Abort()
		    return
		}

		ctx.Next()
		  
	}
}