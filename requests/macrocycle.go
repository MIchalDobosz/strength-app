package requests

type CreateMacrocycle struct {
	Name string `json:"name"`
}

func (createMacrocycle CreateMacrocycle) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("name", createMacrocycle.Name, &valid, errors)
	return valid, errors
}

type UpdateMacrocycle struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (updateMacrocycle UpdateMacrocycle) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("id", updateMacrocycle.Id, &valid, errors)
	validateRequired("name", updateMacrocycle.Name, &valid, errors)
	return valid, errors
}
