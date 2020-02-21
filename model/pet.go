package model

import "errors"

var (
	EMPTY_ID = errors.New("string is empty")
)

type PetModel interface {
	GetId() (string, error)
	GetName() (string, error)
	GetDesc() (string, error)
}

type PetInstance struct {
	Id   string
	Name string
	Desc string
}

func (p *PetInstance) GetId() (string, error) {
	if p.Id == "" {
		return "", EMPTY_ID
	}

	return p.Id, nil
}
