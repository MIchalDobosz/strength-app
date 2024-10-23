package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strength-app/models"
	"strength-app/requests"
	"strength-app/resources"

	"github.com/jmoiron/sqlx"
)

type Exercise struct{}

func (Exercise) Index(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	exercises := models.Exercises{}
	if err := models.SelectAll(db, &exercises); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Exercises not found "), nil
		}
		return Response{}, Error(err)
	}

	resource := resources.Exercises{}
	resources.Resources(exercises, &resource, resources.Exercise{}.New)

	return Response{}.SuccessWithData("Exercises found.", resource), nil
}

func (Exercise) Show(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	exercise := models.Exercise{Id: id}
	if err := models.SelectOne(db, &exercise); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Exercise not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Exercise found.", resources.ExerciseDetails{}.New(exercise)), nil
}

func (Exercise) Store(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	createExercise := requests.CreateExercise{}
	requests.Read(request.Body, &createExercise)
	valid, errors := createExercise.Validate()
	if !valid {
		return Response{}.ValidationFailed("", errors), nil
	}

	exercise := models.Exercise{}
	exercise.FromCreateRequest(createExercise)
	if err := models.Insert(db, &exercise); err != nil {
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Exercise created.", map[string]int{"id": exercise.Id}), nil
}

func (Exercise) Update(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	updateExercise := requests.UpdateExercise{Id: id}
	requests.Read(request.Body, &updateExercise)
	valid, validationErrors := updateExercise.Validate()
	if !valid {
		return Response{}.ValidationFailed("", validationErrors), nil
	}

	exercise := models.Exercise{}
	exercise.FromUpdateRequest(updateExercise)
	if err := models.Update(db, exercise); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Exercise not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Exercise updated."), nil
}

func (Exercise) Destroy(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	if err := models.Delete(db, models.Exercise{Id: id}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Exercise not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Exercise deleted."), nil
}
