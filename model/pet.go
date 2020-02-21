package model

import (
	"errors"
	"fmt"
)

var (
	errorEmptyId   = errors.New("id is empty")
	errorEmptyName = errors.New("name is empty")
)

// PetModel is meant to simplify the response from Elasticsearch
// Abstracts the details away from the user

type PetModel interface {
	SetId(string)
	GetId() string
	GetName() string
	GetDesc() string
}

type PetInstance struct {
	id   string
	name string
	desc string
}

func NewPetInstance(name string, desc string) *PetInstance {
	return &PetInstance{
		name: name,
		desc: desc,
	}
}

func NewPetInstanceWithId(id string, name string, desc string) *PetInstance {
	return &PetInstance{
		id:   id,
		name: name,
		desc: desc,
	}
}

func (p *PetInstance) SetId(id string) {
	p.id = id
}

func (p *PetInstance) GetId() string {
	return p.id
}

func (p *PetInstance) GetName() string {
	return p.name
}

func (p *PetInstance) GetDesc() string {
	return p.desc
}

func (p PetInstance) String() string {
	return fmt.Sprintf("ID: %s\nName: %s\nDescription: %s", p.id, p.name, p.desc)
}
