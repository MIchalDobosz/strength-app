package models

import (
	"strength-app/requests"
)

type Macrocycles []Macrocycle
type Macrocycle struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Mesocycles Mesocycles
}

func (Macrocycle) Table() string {
	return "macrocycles"
}

func (Macrocycle) PrimaryKeyColumn() string {
	return "id"
}

func (macrocycle Macrocycle) PrimaryKeyValue() int {
	return macrocycle.Id
}

func (macrocycle *Macrocycle) FromCreateRequest(createMacrocycle requests.CreateMacrocycle) {
	macrocycle.Name = createMacrocycle.Name
}

func (macrocycle *Macrocycle) FromUpdateRequest(updateMacrocycle requests.UpdateMacrocycle) {
	macrocycle.Id = updateMacrocycle.Id
	macrocycle.Name = updateMacrocycle.Name
}
