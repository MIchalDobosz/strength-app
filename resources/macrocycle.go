package resources

import "strength-app/models"

type Macrocycles []Macrocycle
type Macrocycle struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (Macrocycle) New(model models.Macrocycle) Macrocycle {
	return Macrocycle{
		Id:   model.Id,
		Name: model.Name,
	}
}

type MacrocycleDetails struct {
	Macrocycle
	Mesocycles Mesocycles `json:"mesocycles"`
}

func (MacrocycleDetails) New(model models.Macrocycle) MacrocycleDetails {
	details := MacrocycleDetails{
		Macrocycle: Macrocycle{}.New(model),
	}
	Resources(model.Mesocycles, &details.Mesocycles, Mesocycle{}.New)
	return details
}
