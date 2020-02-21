package model

import "errors"

var (
	errorEmptyId   = errors.New("id is empty")
	errorEmptyName = errors.New("name is empty")
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
		return "", errorEmptyId
	}

	return p.Id, nil
}

func (p *PetInstance) GetName() (string, error) {
	if p.Name == "" {
		return "", errorEmptyName
	}

	return p.Name, nil
}

func (p *PetInstance) GetDesc() (string, error) {
	return p.Desc, nil
}
