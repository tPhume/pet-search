package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tPhume/pet-search/model"
	"github.com/tPhume/pet-search/search"
	"net/http"
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

// to bind to request body
type petRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func RegisterPetRoutes(router *gin.Engine, search search.Pet) {
	router.Use(setSearch(search))

	v1 := router.Group("/api/v1/pets")
	v1.POST("", addPetHandler)
	v1.GET("/:id", searchPetByIdHandler)
}

// returns handler with search passed into value
func setSearch(search search.Pet) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("search", search)
	}
}

// handler to add pet
func addPetHandler(ctx *gin.Context) {
	temp, ok := ctx.Get("search")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	s := temp.(search.Pet)
	body := petRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	pm, err := model.NewPetInstance(body.Name, body.Desc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	id, err := s.AddPet(ctx, pm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// handler to search by id
func searchPetByIdHandler(ctx *gin.Context) {
	temp, ok := ctx.Get("search")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	s := temp.(search.Pet)
	id := ctx.Param("id")

	pm, err := s.SearchPetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusOK, petResponse{
		Id:   pm.GetId(),
		Name: pm.GetName(),
		Desc: pm.GetDesc(),
	})
}
