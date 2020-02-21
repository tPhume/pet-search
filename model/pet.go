package model

import "errors"

var (
	EMPTY_ID = errors.New("id is empty")
	EMPTY_NAME = errors.New("name is empty")

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

func (p *PetInstance) GetName() (string, error) {
	if p.Name == "" {
		return "", EMPTY_NAME
	}

	return p.Name, nil
}