package requests

type CreateSet struct {
	SlotId            int     `json:"slot_id"`
	SetNo             int     `json:"set_no"`
	PlannedReps       int     `json:"planned_reps"`
	PlannedRepsUnit   int     `json:"planned_reps_unit"`
	PerformedReps     int     `json:"performed_reps"`
	PerformedRepsUnit int     `json:"performed_reps_unit"`
	PlannedPercent    int     `json:"planned_percent"`
	PlannedRpe        int     `json:"planned_rpe"`
	PerformedRpe      int     `json:"performed_rpe"`
	PlannedWeight     float32 `json:"planned_weight"`
	PerformedWeight   float32 `json:"performed_weight"`
	Completed         bool    `json:"completed"`
}

func (createSet CreateSet) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("slot_id", createSet.SlotId, &valid, errors)
	validateRequired("set_no", createSet.SetNo, &valid, errors)
	validateRequired("planned_reps_unit", createSet.PlannedRepsUnit, &valid, errors)
	return valid, errors
}

type UpdateSet struct {
	Id int `json:"id"`
	CreateSet
}

func (updateSet UpdateSet) Validate() (bool, map[string]string) {
	valid, errors := updateSet.CreateSet.Validate()
	validateRequired("id", updateSet.Id, &valid, errors)
	return valid, errors
}
