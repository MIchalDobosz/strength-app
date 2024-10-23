package requests

type CreateMicrocycle struct {
	MesocycleId int    `json:"mesocycle_id"`
	Name        string `json:"name"`
}

func (createMicrocycle CreateMicrocycle) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("mesocycle_id", createMicrocycle.MesocycleId, &valid, errors)
	validateRequired("name", createMicrocycle.Name, &valid, errors)
	return valid, errors
}

type UpdateMicrocycle struct {
	Id          int    `json:"id"`
	MesocycleId int    `json:"mesocycle_id"`
	Name        string `json:"name"`
}

func (updateMicrocycle UpdateMicrocycle) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("id", updateMicrocycle.Id, &valid, errors)
	validateRequired("mesocycle_id", updateMicrocycle.MesocycleId, &valid, errors)
	validateRequired("name", updateMicrocycle.Name, &valid, errors)
	return valid, errors
}
