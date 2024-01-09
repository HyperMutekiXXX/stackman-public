package serve

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ServerOption struct {
	Addr string `json:"addr"`
}

func Run(opt *ServerOption) {
	r := gin.Default()
	r.Use(middlewareBox...)
	loadMethod(r)

	if err := r.Run(opt.Addr); err != nil {
		log.Fatalf("start server err %v", err)
	}
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewareBox...)
	loadMethod(r)
	return r
}

func loadMethod(r *gin.Engine) {
	for _, v := range controllerBox {
		for method, handlerFunc := range v.GetRouterMap() {
			caseMethod(r, method, handlerFunc)
		}
	}
}

func caseMethod(r *gin.Engine, method *Method, handlerFunc gin.HandlerFunc) {
	switch method.Type {
	case Post:
		r.POST(method.Path, handlerFunc)
	case Get:
		r.GET(method.Path, handlerFunc)
	case Delete:
		r.DELETE(method.Path, handlerFunc)
	case Put:
		r.PUT(method.Path, handlerFunc)
	}

}
