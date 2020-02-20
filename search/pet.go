package search

import (
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
)

type Pet interface {
	CheckStatus() error // Must return nil if there is no error
}

// Concrete implementation
type PetClient struct {
	es *elasticsearch.Client
}

func (pc *PetClient) CheckStatus() error {
	_, err := pc.es.Info()
	if err != nil {
		return errors.New(fmt.Sprintf("cluster returned error: %s", err))
	}

	return nil
}
