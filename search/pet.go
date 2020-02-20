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
)

type PetModel struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Pet interface {
	CheckStatus() (*esapi.Response, error)
	AddPet(context.Context, PetModel) (*esapi.Response, error)
	SearchPetByID(context.Context, string) (*esapi.Response, error)
}

// Concrete implementation
type PetClient struct {
	es *elasticsearch.Client
}

func (pc *PetClient) CheckStatus() (*esapi.Response, error) {
	res, err := pc.es.Info()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cluster returned error: %s", err))
	}

	return res, nil
}

func (pc *PetClient) AddPet(ctx context.Context, pm PetModel) (*esapi.Response, error) {
	id := uuid.New()
	bodyBytes, err := json.Marshal(pm)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not marshal struct: %s", err))
	}

	req := esapi.IndexRequest{
		Index:      "pets",
		DocumentID: id.String(),
		Body:       bytes.NewReader(bodyBytes),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, pc.es)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not index document: %s", err))
	}

	return res, nil
}

func (pc *PetClient) SearchPetByID(ctx context.Context, id string) (*esapi.Response, error) {
	res, err := pc.es.Search(
		pc.es.Search.WithIndex("pets"),
		pc.es.Search.WithQuery(fmt.Sprintf("_id")),
		pc.es.Search.WithContext(ctx),
		pc.es.Search.WithPretty(),
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not get document: %s", err))
	}

	return res, nil
}
