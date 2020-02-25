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
	SetId(string) error
	GetId() string
	SetName(string) error
	GetName() string
	SetDesc(string) error
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

func (p *PetInstance) SetId(id string) error {
	trimId := strings.TrimSpace(id)
	if err := checkId(trimId); err != nil {
		return err
	}


	p.id = trimId
	return nil
}

func (p *PetInstance) GetId() string {
	return p.id
}

func (p *PetInstance) SetName(name string) error {
	trimName := strings.TrimSpace(name)
	if err := checkName(trimName); err != nil {
		return err
	}

	p.name = trimName
	return nil
}

func (p *PetInstance) GetName() string {
	return p.name
}

func (p *PetInstance) SetDesc(desc string) error {
	trimDesc := strings.TrimSpace(desc)
	p.desc = trimDesc
	return nil
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
