package resources

import (
	"strength-app/models"
)

type Sets []Set
type Set struct {
	Id                int     `json:"id"`
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

func (Set) New(set models.Set) Set {
	return Set{
		Id:                set.Id,
		SlotId:            set.SlotId,
		SetNo:             set.SetNo,
		PlannedReps:       set.PlannedReps,
		PlannedRepsUnit:   set.PlannedRepsUnit,
		PerformedReps:     set.PerformedReps,
		PerformedRepsUnit: set.PerformedRepsUnit,
		PlannedPercent:    set.PlannedPercent,
		PlannedRpe:        set.PlannedRpe,
		PerformedRpe:      set.PerformedRpe,
		PlannedWeight:     set.PlannedWeight,
		PerformedWeight:   set.PerformedWeight,
		Completed:         set.Completed,
	}
}

type SetDetails struct {
	Set
}

func (SetDetails) New(model models.Set) SetDetails {
	details := SetDetails{
		Set: Set{}.New(model),
	}
	return details
}
