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
)

type Pet interface {
	CheckStatus() (*esapi.Response, error)
	AddPet(context.Context, model.PetModel) (model.PetModel, error)
	SearchPetByID(context.Context, string) (*esapi.Response, error)
	UpdatePetByID(context.Context, model.PetModel) (*esapi.Response, error)
	DeletePetByID(context.Context, string) (*esapi.Response, error)
	ListPetByName(context.Context, string) ([]model.PetModel, error)
	ListAllPet(context.Context) ([]model.PetModel, error)
}

type petRequest struct {
	name string
	desc string
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

func (pc *PetClient) AddPet(ctx context.Context, pm model.PetModel) (model.PetModel, error) {
	id := uuid.New()

	bodyBytes, err := json.Marshal(petRequest{
		name: pm.GetName(),
		desc: pm.GetDesc(),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not marshal struct: %s", err))
	}

	req := esapi.IndexRequest{
		Index:      "pets",
		DocumentID: id.String(),
		Body:       bytes.NewReader(bodyBytes),
		Refresh:    "true",
		Pretty:     true,
	}

	res, err := req.Do(ctx, pc.es)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not index document: %s", err))
	}

	indexRes, err := model.BodyToIndexResponse(res.Body)
	if err != nil {
		return nil, err
	}

	pm.SetId(indexRes.ID)
	return pm, nil
}

func (pc *PetClient) SearchPetByID(ctx context.Context, id string) (*esapi.Response, error) {
	res, err := pc.es.Search(
		pc.es.Search.WithIndex("pets"),
		pc.es.Search.WithQuery(fmt.Sprintf("_id:%s", id)),
		pc.es.Search.WithContext(ctx),
		pc.es.Search.WithPretty(),
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get document: %s", err))
	}

	return res, nil
}

func (pc *PetClient) UpdatePetByID(ctx context.Context, pm model.PetModel) (*esapi.Response, error) {
	bodyBytes, err := json.Marshal(petRequest{
		name: pm.GetName(),
		desc: pm.GetDesc(),
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

func (pc *PetClient) DeletePetByID(ctx context.Context, id string) (*esapi.Response, error) {
	req := esapi.DeleteRequest{
		Index:      "pets",
		DocumentID: id,
		Pretty:     true,
	}

	res, err := req.Do(ctx, pc.es)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not delete document: %s", err))
	}

	return res, nil
}

func (pc *PetClient) ListPetByName(ctx context.Context, name string) ([]model.PetModel, error) {
	res, err := pc.es.Search(
		pc.es.Search.WithIndex("pets"),
		pc.es.Search.WithQuery(fmt.Sprintf("name:%s", name)),
		pc.es.Search.WithContext(ctx),
		pc.es.Search.WithPretty(),
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get document: %s", err))
	}

	queryRes, err := model.BodyToQueryByNameResponse(res.Body)
	if err != nil {
		return nil, err
	}

	pmList := make([]model.PetModel, len(queryRes.Hits.Hits))
	for i, hit := range queryRes.Hits.Hits {
		pmList[i] = model.NewPetInstanceWithId(hit.Index, hit.Source.Name, hit.Source.Desc)
	}

	return pmList, nil
}

func (pc *PetClient) ListAllPet(ctx context.Context) ([]model.PetModel, error) {
	res, err := pc.es.Search(
		pc.es.Search.WithIndex("pets"),
		pc.es.Search.WithContext(ctx),
		pc.es.Search.WithPretty(),
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get document: %s", err))
	}

	queryRes, err := model.BodyToQueryByNameResponse(res.Body)
	if err != nil {
		return nil, err
	}

	pmList := make([]model.PetModel, len(queryRes.Hits.Hits))
	for i, hit := range queryRes.Hits.Hits {
		pmList[i] = model.NewPetInstanceWithId(hit.Index, hit.Source.Name, hit.Source.Desc)
	}

	return pmList, nil
}
