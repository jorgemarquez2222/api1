package storage

import (
	"github.com/jorgemarquez2222/api1/model"
)

type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

func NewMemory() Memory {
	persons := make(map[int]model.Person)
	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNill
	}
	m.currentID++
	m.Persons[m.currentID] = *person
	return nil
}

func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNill
	}
	if _, ok := m.Persons[ID]; !ok {
		return model.ErrPersonCanNotBeNill
	}
	m.Persons[ID] = *person
	return nil
}

func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return model.ErrPersonCanNotBeNill
	}
	delete(m.Persons, ID)
	return nil
}

func (m *Memory) GetById(ID int) (*model.Person, error) {
	person := model.Person{}
	if _, ok := m.Persons[ID]; !ok {
		return &person, model.ErrPersonCanNotBeNill
	}
	person = m.Persons[ID]
	return &person, nil
}

func (m *Memory) GetById2(ID int) (*model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return &person, model.ErrPersonCanNotBeNill
	}
	return &person, nil
}

func (m *Memory) GetAll() ([]model.Person, error) {
	var persons []model.Person
	for _, v := range m.Persons {
		persons = append(persons, v)
	}
	return persons, nil
}

func (m *Memory) GetAll2() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}
	return result, nil
}
