package requests

type CreateExercise struct {
	Name string `json:"name"`
}

func (createExercise CreateExercise) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("name", createExercise.Name, &valid, errors)
	return valid, errors
}

type UpdateExercise struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (updateExercise UpdateExercise) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("id", updateExercise.Id, &valid, errors)
	validateRequired("name", updateExercise.Name, &valid, errors)
	return valid, errors
}
