package serve

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	controllerBox = make([]IController, 0)
)

type HandlerMap map[*Method]gin.HandlerFunc

type IController interface {
	LoadRouter(method *Method, handlerFunc gin.HandlerFunc)
	LoadController(controller IController)
	GetRouterMap() HandlerMap
	SetRouterMap(handlerMap HandlerMap)
}

type Method struct {
	Type string
	Path string
}

type Controller struct {
	RouterMap HandlerMap
}

func (c *Controller) LoadRouter(method *Method, handlerFunc gin.HandlerFunc) {
	if len(c.GetRouterMap()) == 0 {
		c.SetRouterMap(make(HandlerMap))
	}
	c.RouterMap[method] = handlerFunc
}

func (c *Controller) LoadController(controller IController) {
	controllerBox = append(controllerBox, controller)
}

func (c *Controller) GetRouterMap() HandlerMap {
	return c.RouterMap
}

func (c *Controller) SetRouterMap(handlerMap HandlerMap) {
	c.RouterMap = handlerMap
}

func (c *Controller) LoadGetMethod(path string, handlerFunc gin.HandlerFunc) {
	c.LoadRouter(&Method{
		Type: Get,
		Path: path,
	}, handlerFunc)
}

func (c *Controller) LoadPostMethod(path string, handlerFunc gin.HandlerFunc) {
	c.LoadRouter(&Method{
		Type: Post,
		Path: path,
	}, handlerFunc)
}

func (c *Controller) LoadPutMethod(path string, handlerFunc gin.HandlerFunc) {
	c.LoadRouter(&Method{
		Type: Put,
		Path: path,
	}, handlerFunc)
}

func (c *Controller) LoadDeleteMethod(path string, handlerFunc gin.HandlerFunc) {
	c.LoadRouter(&Method{
		Type: Delete,
		Path: path,
	}, handlerFunc)
}

func (c *Controller) Success(ctx *gin.Context, msg string, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
}

func (c *Controller) Error(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusBadRequest,
		"msg":  msg,
	})
}
