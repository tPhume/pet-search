package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tPhume/pet-search/search"
)

// to marshal for response
type petResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type petResponseList struct {
	Result []*petResponse `json:"result"`
}

func RegisterPetRoutes(router *gin.Engine, search search.Pet) {
	router.Use(setSearch(search))
}

// returns handler with search passed into value
func setSearch(search search.Pet) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("search", search)
	}
}
