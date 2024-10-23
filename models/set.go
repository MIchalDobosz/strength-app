package models

import (
	"strength-app/requests"
)

type Sets []Set
type Set struct {
	Id                int     `db:"id"`
	SlotId            int     `db:"slot_id"`
	SetNo             int     `db:"set_no"`
	PlannedReps       int     `db:"planned_reps"`
	PlannedRepsUnit   int     `db:"planned_reps_unit"`
	PerformedReps     int     `db:"performed_reps"`
	PerformedRepsUnit int     `db:"performed_reps_unit"`
	PlannedPercent    int     `db:"planned_percent"`
	PlannedRpe        int     `db:"planned_rpe"`
	PerformedRpe      int     `db:"performed_rpe"`
	PlannedWeight     float32 `db:"planned_weight"`
	PerformedWeight   float32 `db:"performed_weight"`
	Completed         bool    `db:"completed"`
	Slot              Slot
}

func (Set) Table() string {
	return "sets"
}

func (Set) PrimaryKeyColumn() string {
	return "id"
}

func (set Set) PrimaryKeyValue() int {
	return set.Id
}

func (set *Set) FromCreateRequest(createSet requests.CreateSet) {
	set.SlotId = createSet.SlotId
	set.SetNo = createSet.SetNo
	set.PerformedReps = createSet.PlannedReps
	set.PlannedRepsUnit = createSet.PlannedRepsUnit
	set.PerformedReps = createSet.PerformedReps
	set.PerformedRepsUnit = createSet.PerformedRepsUnit
	set.PlannedPercent = createSet.PlannedPercent
	set.PlannedRpe = createSet.PlannedRpe
	set.PerformedRpe = createSet.PerformedRpe
	set.PlannedWeight = createSet.PlannedWeight
	set.PerformedWeight = createSet.PerformedWeight
	set.Completed = createSet.Completed
}

func (set *Set) FromUpdateRequest(updateSet requests.UpdateSet) {
	set.FromCreateRequest(updateSet.CreateSet)
	set.Id = updateSet.Id
}
