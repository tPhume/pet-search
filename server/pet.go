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
	v1.PUT("", updatePetAllHandler)
	v1.DELETE("", deletePetByIdHandler)
	v1.GET("", muxGetHandler)
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
		ctx.Status(http.StatusInternalServerError)
		return
	}

	s := temp.(search.Pet)
	body := petRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	pm, err := model.NewPetInstance(body.Name, body.Desc)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, err := s.AddPet(ctx, pm)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// handler to update pet (all field)
func updatePetAllHandler(ctx *gin.Context) {
	temp, ok := ctx.Get("search")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	s := temp.(search.Pet)
	id := ctx.Query("id")

	body := petRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	pm, err := model.NewPetInstanceWithId(id, body.Name, body.Desc)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = s.UpdatePetAll(ctx, pm)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}

// handler to delete pet by id
func deletePetByIdHandler(ctx *gin.Context) {
	temp, ok := ctx.Get("search")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	s := temp.(search.Pet)
	id := ctx.Query("id")

	if err := s.DeletePetByID(ctx, id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}

// mux get handler
func muxGetHandler(ctx *gin.Context) {
	temp, ok := ctx.Get("search")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	s := temp.(search.Pet)

	id := ctx.Query("id")
	if id != "" {
		searchPetByIdHandler(ctx, s)
		return
	}

	name := ctx.Query("name")
	if name != "" {
		listPetByNameHandler(ctx, s)
		return
	}

	listAllPetHandler(ctx, s)
}

// handler to search by id
func searchPetByIdHandler(ctx *gin.Context, s search.Pet) {
	id := ctx.Param("id")

	pm, err := s.SearchPetByID(ctx, id)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, petResponse{
		Id:   pm.GetId(),
		Name: pm.GetName(),
		Desc: pm.GetDesc(),
	})
}

// handler to search by name
func listPetByNameHandler(ctx *gin.Context, s search.Pet) {
	name := ctx.Param("name")

	pmList, err := s.ListPetByName(ctx, name)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resList := petResponseList{Result: make([]*petResponse, len(pmList))}
	for i, pm := range pmList {
		resList.Result[i] = &petResponse{
			Id:   pm.GetId(),
			Name: pm.GetName(),
			Desc: pm.GetDesc(),
		}
	}

	ctx.JSON(http.StatusOK, resList)
}

func listAllPetHandler(ctx *gin.Context, s search.Pet) {
	pmList, err := s.ListAllPet(ctx)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resList := petResponseList{Result: make([]*petResponse, len(pmList))}
	for i, pm := range pmList {
		resList.Result[i] = &petResponse{
			Id:   pm.GetId(),
			Name: pm.GetName(),
			Desc: pm.GetDesc(),
		}
	}

	ctx.JSON(http.StatusOK, resList)
}
