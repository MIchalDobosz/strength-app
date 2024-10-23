package requests

type CreateMesocycle struct {
	Name         string `json:"name"`
	MacrocycleId int    `json:"macrocycle_id"`
}

func (createMesocycle CreateMesocycle) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("name", createMesocycle.Name, &valid, errors)
	validateRequired("macrocycle_id", createMesocycle.MacrocycleId, &valid, errors)
	return valid, errors
}

type UpdateMesocycle struct {
	Id           int    `json:"id"`
	MacrocycleId int    `json:"macrocycle_id"`
	Name         string `json:"name"`
}

func (updateMesocycle UpdateMesocycle) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("id", updateMesocycle.Id, &valid, errors)
	validateRequired("name", updateMesocycle.Name, &valid, errors)
	validateRequired("macrocycle_id", updateMesocycle.MacrocycleId, &valid, errors)
	return valid, errors
}
