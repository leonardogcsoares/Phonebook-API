package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateEntry TODO
func (r Router) CreateEntry(c *gin.Context) {

}

// GetEntry TODO
func (r Router) GetEntry(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Get:", id)
}

// UpdateEntry TODO
func (r Router) UpdateEntry(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Update:", id)
}

// DeleteEntry TODO
func (r Router) DeleteEntry(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Delete:", id)
}
