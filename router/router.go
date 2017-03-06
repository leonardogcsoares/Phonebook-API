package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardogcsoares/phonebook-api/router/repo"
	"github.com/leonardogcsoares/phonebook-api/router/validator"
	"github.com/syndtr/goleveldb/leveldb"
)

// Router is the hub for routing and handlers
type Router struct {
	port      string
	engine    *gin.Engine
	Repo      repo.Repo
	validator validator.V
}

// New TODO
func New(port string, db *leveldb.DB) Router {
	r := Router{
		port:      port,
		Repo:      repo.NewRepo(db),
		validator: validator.Impl{},
	}

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.POST("/login", r.login)

	engine.Use(r.validate())

	engine.POST("/phone", r.CreateEntry)
	engine.GET("/phone/:id", r.GetEntry)
	engine.PUT("/phone/:id", r.UpdateEntry)
	engine.DELETE("/phone/:id", r.DeleteEntry)

	r.engine = engine

	return r
}

// Start TODO
func (r Router) Start() error {
	return r.engine.Run(":" + r.port)
}

// Engine TODO
func (r *Router) Engine() *gin.Engine {
	return r.engine
}
