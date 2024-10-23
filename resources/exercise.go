package resources

import "strength-app/models"

type Exercises []Exercise
type Exercise struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (Exercise) New(exercise models.Exercise) Exercise {
	return Exercise{
		Id:   exercise.Id,
		Name: exercise.Name,
	}
}

type ExerciseDetails struct {
	Exercise
}

func (ExerciseDetails) New(model models.Exercise) ExerciseDetails {
	return ExerciseDetails{
		Exercise: Exercise{}.New(model),
	}
}
