package resources

import "strength-app/models"

type Microcycles []Microcycle
type Microcycle struct {
	Id          int    `json:"id"`
	MesocycleId int    `json:"mesocycle_id"`
	Name        string `json:"name"`
}

func (Microcycle) New(microcycle models.Microcycle) Microcycle {
	return Microcycle{
		Id:          microcycle.Id,
		MesocycleId: microcycle.MesocycleId,
		Name:        microcycle.Name,
	}
}

type MicrocycleDetails struct {
	Microcycle
	Sessions Sessions `json:"sessions"`
}

func (MicrocycleDetails) New(model models.Microcycle) MicrocycleDetails {
	details := MicrocycleDetails{
		Microcycle: Microcycle{}.New(model),
	}
	Resources(model.Sessions, &details.Sessions, Session{}.New)
	return details
}
