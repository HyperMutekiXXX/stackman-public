package controller

import "github.com/gin-gonic/gin"

var middlewareBox = make([]gin.HandlerFunc, 0)

func LoadMiddleware(handlerFunc gin.HandlerFunc) {
	middlewareBox = append(middlewareBox, handlerFunc)
}
