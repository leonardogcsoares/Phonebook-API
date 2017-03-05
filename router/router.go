package router

import "github.com/gin-gonic/gin"

// Router is the hub for routing and handlers
type Router struct {
	port   string
	engine *gin.Engine
}

// New TODO
func New(port string) Router {
	r := Router{
		port: port,
	}

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(r.validate())

	engine.POST("/phone", r.CreateEntry)
	engine.GET("/phone/:id", r.GetEntry)
	engine.PUT("/phone/:id", r.UpdateEntry)
	engine.DELETE("/phone/:id", r.DeleteEntry)

	engine.POST("/login", r.login)

	r.engine = engine

	return r
}

// Start TODO
func (r Router) Start() error {
	return r.engine.Run(":" + r.port)
}
