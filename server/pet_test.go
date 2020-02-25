package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tPhume/pet-search/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUp() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
}

func TestHappyPathV1(t *testing.T) {
	RegisterPetRoutes(router, searchInstance)

	// ---- Test Add Pet ----
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/pets", bytes.NewReader(jsonAdd))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected = [%v], got = [%v]", http.StatusCreated, w.Code)
	}

	temp := make(map[string]interface{})
	_ = json.Unmarshal(w.Body.Bytes(), &temp)

	if temp["id"] != sushiInstance.GetId() {
		t.Fatalf("expected = [%s], got = [%s]", sushiInstance.GetId(), temp["id"])
	}

	// ---- Test Search by id ----
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/pets/%s", sushiInstance.GetId()), nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected = [%v], got = [%v]", http.StatusOK, w.Code)
	}

	res := petResponse{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)

	if err := checkSushiRes(res); err != nil {
		t.Fatal(err)
	}

	// ---- Test Update Pet All ----
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/pets/%s", sushiInstance.GetId()), bytes.NewReader(jsonAdd))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected = [%v], got = [%v]", http.StatusOK, w.Code)
	}

	// ---- Test Delete Pet ----
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/pets/%s", sushiInstance.GetId()), nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected = [%v], got = [%v]", http.StatusOK, w.Code)
	}

	// ---- Test List Pet by name ----
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/pets/%s", sushiInstance.GetName()), nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected = [%v], got = [%v]", http.StatusOK, w.Code)
	}

	_ = json.Unmarshal(w.Body.Bytes(), &res)
	if err = checkSushiRes(petResponse.pets[0]); err != nil {
		t.Fatal(err)
	}

	// ---- Test List All Pet ----
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/v1/pets", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected = [%v], got = [%v]", http.StatusOK, w.Code)
	}

	_ = json.Unmarshal(w.Body.Bytes(), &res)
	if err = checkSushiRes(petResponse.pets[0]); err != nil {
		t.Fatal(err)
	}
}

// Variables
var (
	router *gin.Engine

	searchInstance   = petSearchHappy{}
	sushiInstance, _ = model.NewPetInstanceWithId("1", "Sushi", "Sushi is a good boy")

	notSushiId   = errors.New("incorrect id")
	notSushiName = errors.New("incorrect name")
	notSushiDesc = errors.New("incorrect description")

	bodyAdd    = map[string]string{"name": sushiInstance.GetName(), "desc": sushiInstance.GetDesc()}
	jsonAdd, _ = json.Marshal(bodyAdd)
)

// Is used for testing the happy path
type petSearchHappy struct{}

func (p petSearchHappy) CheckStatus() error {
	return nil
}

func (p petSearchHappy) AddPet(ctx context.Context, pm model.PetModel) (string, error) {
	return sushiInstance.GetId(), nil
}

func (p petSearchHappy) SearchPetByID(ctx context.Context, id string) (model.PetModel, error) {
	return sushiInstance, nil
}

func (p petSearchHappy) UpdatePetAll(ctx context.Context, pm model.PetModel) error {
	return nil
}

func (p petSearchHappy) DeletePetByID(ctx context.Context, id string) error {
	return nil
}

func (p petSearchHappy) ListPetByName(ctx context.Context, name string) ([]model.PetModel, error) {
	return []model.PetModel{sushiInstance}, nil
}

func (p petSearchHappy) ListAllPet(ctx context.Context) ([]model.PetModel, error) {
	return []model.PetModel{sushiInstance}, nil
}

// utilities function
func checkSushiRes(res petResponse) error {
	if res.Id != sushiInstance.GetId() {
		return notSushiId
	}

	if res.Name != sushiInstance.GetName() {
		return notSushiName
	}

	if res.Desc != sushiInstance.GetDesc() {
		return notSushiDesc
	}

	return nil
}
