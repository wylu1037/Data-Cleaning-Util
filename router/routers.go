package router

import "github.com/gin-gonic/gin"

type Router func(c *gin.Engine)

var routes []Router

func Register(r ...Router) {
	routes = append(routes, r...)
}

func Init() *gin.Engine {
	r := gin.New()
	for _, route := range routes {
		route(r)
	}
	return r
}
