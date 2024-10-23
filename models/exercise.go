package models

import (
	"database/sql"
	"strength-app/requests"
)

type Exercises []Exercise
type Exercise struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type NullableExercise struct {
	Id   sql.NullInt64  `db:"id"`
	Name sql.NullString `db:"name"`
}

func (Exercise) Table() string {
	return "exercises"
}

func (Exercise) PrimaryKeyColumn() string {
	return "id"
}

func (exercise Exercise) PrimaryKeyValue() int {
	return exercise.Id
}

func (exercise *Exercise) FromCreateRequest(createExercise requests.CreateExercise) {
	exercise.Name = createExercise.Name
}

func (exercise *Exercise) FromUpdateRequest(updateExercise requests.UpdateExercise) {
	exercise.Id = updateExercise.Id
	exercise.Name = updateExercise.Name
}
