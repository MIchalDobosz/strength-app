package resources

import (
	"strength-app/models"
)

type Mesocycles []Mesocycle
type Mesocycle struct {
	Id           int    `json:"id"`
	MacrocycleId int    `json:"macrocycle_id"`
	Name         string `json:"name"`
}

func (Mesocycle) New(model models.Mesocycle) Mesocycle {
	return Mesocycle{
		Id:           model.Id,
		MacrocycleId: model.MacrocycleId,
		Name:         model.Name,
	}
}

type MesocycleDetails struct {
	Mesocycle
	Microcycles Microcycles `json:"microcycles"`
}

func (MesocycleDetails) New(model models.Mesocycle) MesocycleDetails {
	details := MesocycleDetails{
		Mesocycle: Mesocycle{}.New(model),
	}
	Resources(model.Microcycles, &details.Microcycles, Microcycle{}.New)
	return details
}
