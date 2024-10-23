package models

import (
	"strength-app/requests"
)

type Mesocycles []Mesocycle
type Mesocycle struct {
	Id           int    `db:"id"`
	MacrocycleId int    `db:"macrocycle_id"`
	Name         string `db:"name"`
	Macrocycle   Macrocycle
	Microcycles  Microcycles
}

func (Mesocycle) Table() string {
	return "mesocycles"
}

func (Mesocycle) PrimaryKeyColumn() string {
	return "id"
}

func (mesocycle Mesocycle) PrimaryKeyValue() int {
	return mesocycle.Id
}

func (mesocycle *Mesocycle) FromCreateRequest(createMesocycle requests.CreateMesocycle) {
	mesocycle.MacrocycleId = createMesocycle.MacrocycleId
	mesocycle.Name = createMesocycle.Name
}

func (mesocycle *Mesocycle) FromUpdateRequest(updateMesocycle requests.UpdateMesocycle) {
	mesocycle.Id = updateMesocycle.Id
	mesocycle.MacrocycleId = updateMesocycle.MacrocycleId
	mesocycle.Name = updateMesocycle.Name
}
