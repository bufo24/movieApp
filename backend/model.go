package models

type Model interface {
	GetId() string
	SetId(id string)
}

type modelImpl struct {
	id string
}

func (m *modelImpl) setId(id string) { 
	m.id = id 
}