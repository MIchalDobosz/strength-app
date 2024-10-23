package requests

type CreateSlot struct {
	SessionId           int  `json:"session_id"`
	PlannedExerciseId   int  `json:"planned_exercise_id"`
	PerformedExerciseId int  `json:"performed_exercise_id"`
	Completed           bool `json:"completed"`
}

func (createSlot CreateSlot) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("session_id", createSlot.SessionId, &valid, errors)
	validateRequired("planned_exercise_id", createSlot.PlannedExerciseId, &valid, errors)
	return valid, errors
}

type UpdateSlot struct {
	Id int `json:"id"`
	CreateSlot
}

func (updateSlot UpdateSlot) Validate() (bool, map[string]string) {
	valid, errors := updateSlot.CreateSlot.Validate()
	validateRequired("id", updateSlot.Id, &valid, errors)
	return valid, errors
}
