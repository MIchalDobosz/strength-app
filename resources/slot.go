package resources

import "strength-app/models"

type Slots []Slot
type Slot struct {
	Id                    int    `json:"id"`
	SessionId             int    `json:"session_id"`
	PlannedExerciseId     int    `json:"planned_exercise_id"`
	PlannedExerciseName   string `json:"planned_exercise_name"`
	PerformedExerciseId   int    `json:"performed_exercise_id"`
	PerformedExerciseName string `json:"performed_exercise_name"`
	Completed             bool   `json:"completed"`
}

func (Slot) New(slot models.Slot) Slot {
	return Slot{
		Id:                    slot.Id,
		SessionId:             slot.SessionId,
		PlannedExerciseId:     slot.PlannedExerciseId,
		PlannedExerciseName:   slot.PlannedExercise.Name,
		PerformedExerciseId:   int(slot.PerformedExerciseId.Int64),
		PerformedExerciseName: slot.PerformedExercise.Name,
	}
}

type SlotDetails struct {
	Slot
	PlannedExercise   Exercise
	PerformedExercise Exercise
	Sets              Sets
}

func (SlotDetails) New(model models.Slot) SlotDetails {
	details := SlotDetails{
		Slot:              Slot{}.New(model),
		PlannedExercise:   Exercise{}.New(model.PlannedExercise),
		PerformedExercise: Exercise{}.New(model.PerformedExercise),
	}
	Resources(model.Sets, &details.Sets, Set{}.New)
	return details
}
