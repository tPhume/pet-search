package model

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errorShortId   = errors.New("id is too short: length must be [1,128]")
	errorLongId    = errors.New("id is too long: length must be [1,128]")
	errorShortName = errors.New("name is too short: length must be [3,32]")
	errorLongName  = errors.New("name is too long: length must be between [3,32]")
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

// trims space of name and desc
func NewPetInstance(name string, desc string) (*PetInstance, error) {
	trimName := strings.TrimSpace(name)
	err := checkName(trimName)
	if err != nil {
		return nil, err
	}

	trimDesc := strings.TrimSpace(desc)

	return &PetInstance{
		name: trimName,
		desc: trimDesc,
	}, nil
}

func NewPetInstanceWithId(id string, name string, desc string) (*PetInstance, error) {
	trimId := strings.TrimSpace(id)
	err := checkId(trimId)
	if err != nil {
		return nil, err
	}

	trimName := strings.TrimSpace(name)
	err = checkName(trimName)
	if err != nil {
		return nil, err
	}

	trimDesc := strings.TrimSpace(desc)

	return &PetInstance{
		id:   id,
		name: trimName,
		desc: trimDesc,
	}, nil
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

func checkName(name string) error {
	if len(name) < 3 {
		return errorShortName
	}
	if len(name) > 32 {
		return errorLongName
	}

	return nil
}

func checkId(id string) error {
	if len(id) < 1 {
		return errorShortId
	}
	if len(id) > 128 {
		return errorLongId
	}

	return nil
}
