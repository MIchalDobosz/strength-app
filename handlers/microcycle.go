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

type Microcycle struct{}

func (Microcycle) Index(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	microcycles := models.Microcycles{}
	if err := models.SelectAll(db, &microcycles); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Microcycles not found "), nil
		}
		return Response{}, Error(err)
	}

	resource := resources.Microcycles{}
	resources.Resources(microcycles, &resource, resources.Microcycle{}.New)

	return Response{}.SuccessWithData("Microcycles found.", resource), nil
}

func (Microcycle) Show(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	microcycle := models.Microcycle{Id: id}
	if err := models.SelectOne(db, &microcycle); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Microcycle not found."), nil
		}
		return Response{}, Error(err)
	}
	if err := models.SelectAllWhere(db, &microcycle.Sessions, "microcycle_id", "=", microcycle.Id); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return Response{}, Error(err)
		}
	}

	return Response{}.SuccessWithData("Microcycle found.", resources.MicrocycleDetails{}.New(microcycle)), nil
}

func (Microcycle) Store(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	createMicrocycle := requests.CreateMicrocycle{}
	requests.Read(request.Body, &createMicrocycle)
	valid, errors := createMicrocycle.Validate()
	if !valid {
		return Response{}.ValidationFailed("", errors), nil
	}

	microcycle := models.Microcycle{}
	microcycle.FromCreateRequest(createMicrocycle)
	if err := models.Insert(db, &microcycle); err != nil {
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Microcycle created.", map[string]int{"id": microcycle.Id}), nil
}

func (Microcycle) Update(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	updateMicrocycle := requests.UpdateMicrocycle{Id: id}
	requests.Read(request.Body, &updateMicrocycle)
	valid, validationErrors := updateMicrocycle.Validate()
	if !valid {
		return Response{}.ValidationFailed("", validationErrors), nil
	}

	microcycle := models.Microcycle{}
	microcycle.FromUpdateRequest(updateMicrocycle)
	if err := models.Update(db, microcycle); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Microcycle not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Microcycle updated."), nil
}

func (Microcycle) Destroy(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	if err := models.Delete(db, models.Microcycle{Id: id}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Microcycle not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Microcycle deleted."), nil
}
