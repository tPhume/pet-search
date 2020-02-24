package search

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
	"github.com/tPhume/pet-search/model"
	"net/http"
)

type Pet interface {
	CheckStatus() (*esapi.Response, error)
	AddPet(context.Context, model.PetModel) (string, error)
	SearchPetByID(context.Context, string) (model.PetModel, error)
	UpdatePetByID(context.Context, model.PetModel) (*esapi.Response, error)
	DeletePetByID(context.Context, string) error
	ListPetByName(context.Context, string) ([]model.PetModel, error)
	ListAllPet(context.Context) ([]model.PetModel, error)
}

type petRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// Concrete implementation
type PetClient struct {
	es *elasticsearch.Client
}

func NewPetClient(es *elasticsearch.Client) *PetClient {
	return &PetClient{es: es}
}

func (pc *PetClient) CheckStatus() (*esapi.Response, error) {
	res, err := pc.es.Info()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cluster returned error: %s", err))
	}

	return res, nil
}

func (pc *PetClient) AddPet(ctx context.Context, pm model.PetModel) (string, error) {
	id := uuid.New()

	bodyBytes, err := json.Marshal(petRequest{
		Name: pm.GetName(),
		Desc: pm.GetDesc(),
	})
	if err != nil {
		return "", errors.New(fmt.Sprintf("could not marshal struct: %s", err))
	}

	req := esapi.CreateRequest{
		Index:      "pets",
		DocumentID: id.String(),
		Body:       bytes.NewReader(bodyBytes),
		Refresh:    "true",
		Pretty:     true,
	}

	res, err := req.Do(ctx, pc.es)
	if err != nil {
		return "", errors.New(fmt.Sprintf("could not index document: %s", err))
	}

	if res.StatusCode != http.StatusCreated {
		return "", errors.New("could not create new pet")
	}

	indexRes, err := model.BodyToIndexResponse(res.Body)
	if err != nil {
		return "", err
	}

	return indexRes.ID, nil
}

func (pc *PetClient) SearchPetByID(ctx context.Context, id string) (model.PetModel, error) {
	res, err := pc.es.Search(
		pc.es.Search.WithIndex("pets"),
		pc.es.Search.WithQuery(fmt.Sprintf("_id:%s", id)),
		pc.es.Search.WithContext(ctx),
		pc.es.Search.WithPretty(),
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get document: %s", err))
	}

	queryRes, err := model.BodyToQueryResponse(res.Body)
	if err != nil {
		return nil, err
	}

	if len(queryRes.Hits.Hits) == 0 {
		return nil, errors.New("no id matched")
	}

	rawRes := queryRes.Hits.Hits[0]
	modelRes := model.NewPetInstanceWithId(rawRes.ID, rawRes.Source.Name, rawRes.Source.Desc)

	return modelRes, nil
}

func (pc *PetClient) UpdatePetByID(ctx context.Context, pm model.PetModel) (*esapi.Response, error) {
	bodyBytes, err := json.Marshal(petRequest{
		Name: pm.GetName(),
		Desc: pm.GetDesc(),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not index documents: %s", err))
	}

	req := esapi.UpdateRequest{
		Index:      "pets",
		DocumentID: pm.GetId(),
		Body:       bytes.NewReader(bodyBytes),
		Pretty:     true,
	}

	res, err := req.Do(ctx, pc.es)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not update document: %s", err))
	}

	return res, nil
}

func (pc *PetClient) DeletePetByID(ctx context.Context, id string) error {
	req := esapi.DeleteRequest{
		Index:      "pets",
		DocumentID: id,
		Pretty:     true,
	}

	res, err := req.Do(ctx, pc.es)
	if err != nil {
		return errors.New(fmt.Sprintf("could not delete document: %s", err))
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("an error occurred")
	}

	return nil
}

func (pc *PetClient) ListPetByName(ctx context.Context, name string) ([]model.PetModel, error) {
	res, err := pc.es.Search(
		pc.es.Search.WithIndex("pets/"),
		pc.es.Search.WithQuery(fmt.Sprintf("name:%s", name)),
		pc.es.Search.WithContext(ctx),
		pc.es.Search.WithPretty(),
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get document: %s", err))
	}

	queryRes, err := model.BodyToQueryResponse(res.Body)
	if err != nil {
		return nil, err
	}

	pmList := make([]model.PetModel, len(queryRes.Hits.Hits))
	for i, hit := range queryRes.Hits.Hits {
		pmList[i] = model.NewPetInstanceWithId(hit.ID, hit.Source.Name, hit.Source.Desc)
	}

	return pmList, nil
}

func (pc *PetClient) ListAllPet(ctx context.Context) ([]model.PetModel, error) {
	res, err := pc.es.Search(
		pc.es.Search.WithIndex("pets/"),
		pc.es.Search.WithContext(ctx),
		pc.es.Search.WithPretty(),
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get document: %s", err))
	}

	queryRes, err := model.BodyToQueryResponse(res.Body)
	if err != nil {
		return nil, err
	}

	pmList := make([]model.PetModel, len(queryRes.Hits.Hits))
	for i, hit := range queryRes.Hits.Hits {
		pmList[i] = model.NewPetInstanceWithId(hit.ID, hit.Source.Name, hit.Source.Desc)
	}

	return pmList, nil
}
