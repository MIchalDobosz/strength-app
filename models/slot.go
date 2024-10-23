package models

import (
	"database/sql"
	"strength-app/requests"

	"github.com/jmoiron/sqlx"
)

type Slots []Slot
type Slot struct {
	Id                  int           `db:"id"`
	SessionId           int           `db:"session_id"`
	PlannedExerciseId   int           `db:"planned_exercise_id"`
	PerformedExerciseId sql.NullInt64 `db:"performed_exercise_id"`
	Completed           bool          `db:"completed"`
	Session             Session
	PlannedExercise     Exercise
	PerformedExercise   Exercise
	Sets                Sets
}

func (Slot) Table() string {
	return "slots"
}

func (Slot) PrimaryKeyColumn() string {
	return "id"
}

func (slot Slot) PrimaryKeyValue() int {
	return slot.Id
}

type slotsDTO []slotDTO
type slotDTO struct {
	Slot              Slot             `db:"sl"`
	PlannedExercise   Exercise         `db:"pl_ex"`
	PerformedExercise NullableExercise `db:"pe_ex"`
}

func (slot *Slot) LoadWithExercises(db *sqlx.DB) error {
	models := slotsDTO{}
	if err := models.LoadWithExercisesBySessionId(db, 0); err != nil {
		return err
	}
	slot.fromSlotDTO(models[0])
	return nil
}

func (slots *Slots) LoadWithExercises(db *sqlx.DB) error {
	return slots.LoadWithExercisesBySessionId(db, 0)
}

func (slots *Slots) LoadWithExercisesBySessionId(db *sqlx.DB, sessionId int) error {
	models := slotsDTO{}
	if err := models.LoadWithExercisesBySessionId(db, sessionId); err != nil {
		return err
	}
	for _, model := range models {
		slot := Slot{}
		slot.fromSlotDTO(model)
		*slots = append(*slots, slot)
	}
	return nil
}

func (models *slotsDTO) LoadWithExercisesBySessionId(db *sqlx.DB, sessionId int) error {
	query := `
		SELECT
			sl.id "sl.id",
			sl.session_id "sl.session_id",
			sl.planned_exercise_id "sl.planned_exercise_id",
			sl.performed_exercise_id "sl.performed_exercise_id",
			pl_ex.id "pl_ex.id",
			pl_ex.name "pl_ex.name",
			pe_ex.id "pe_ex.id",
			pe_ex.name "pe_ex.name"
		FROM ` + Slot{}.Table() + ` sl
		LEFT JOIN ` + Exercise{}.Table() + ` pl_ex ON sl.planned_exercise_id = pl_ex.id
		LEFT JOIN ` + Exercise{}.Table() + ` pe_ex ON sl.performed_exercise_id = pe_ex.id
	`
	args := []any{}
	if sessionId != 0 {
		query += "WHERE sl.session_id = ?"
		args = append(args, sessionId)
	}
	if err := db.Select(models, query, args...); err != nil {
		return err
	}
	if len(*models) == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (slot *Slot) fromSlotDTO(slotDTO slotDTO) {
	slot.Id = slotDTO.Slot.Id
	slot.PlannedExerciseId = slotDTO.Slot.PlannedExerciseId
	slot.PerformedExerciseId = slotDTO.Slot.PerformedExerciseId
	slot.Completed = slotDTO.Slot.Completed
	slot.PlannedExercise = slotDTO.PlannedExercise
	slot.PerformedExercise = Exercise{
		Id:   int(slotDTO.PerformedExercise.Id.Int64),
		Name: slotDTO.PerformedExercise.Name.String,
	}
}

func (slot *Slot) FromCreateRequest(createSlot requests.CreateSlot) {
	slot.SessionId = createSlot.SessionId
	slot.PlannedExerciseId = createSlot.PlannedExerciseId
	slot.PerformedExerciseId.Int64 = int64(createSlot.PerformedExerciseId)
	slot.Completed = createSlot.Completed
}

func (slot *Slot) FromUpdateRequest(updateSlot requests.UpdateSlot) {
	slot.FromCreateRequest(updateSlot.CreateSlot)
	slot.Id = updateSlot.Id
}
