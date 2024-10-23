package models

import (
	"strength-app/requests"
)

type Microcycles []Microcycle
type Microcycle struct {
	Id          int    `db:"id"`
	MesocycleId int    `db:"mesocycle_id"`
	Name        string `db:"name"`
	Mesocycle   Mesocycle
	Sessions    Sessions
}

func (Microcycle) Table() string {
	return "microcycles"
}

func (Microcycle) PrimaryKeyColumn() string {
	return "id"
}

func (microcycle Microcycle) PrimaryKeyValue() int {
	return microcycle.Id
}

func (microcycle *Microcycle) FromCreateRequest(createMicrocycle requests.CreateMicrocycle) {
	microcycle.MesocycleId = createMicrocycle.MesocycleId
	microcycle.Name = createMicrocycle.Name
}

func (microcycle *Microcycle) FromUpdateRequest(updateMicrocycle requests.UpdateMicrocycle) {
	microcycle.Id = updateMicrocycle.Id
	microcycle.MesocycleId = updateMicrocycle.MesocycleId
	microcycle.Name = updateMicrocycle.Name
}
