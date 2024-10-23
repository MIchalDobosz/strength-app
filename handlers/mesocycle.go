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

type Mesocycle struct{}

func (Mesocycle) Index(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	mesocycles := models.Mesocycles{}
	if err := models.SelectAll(db, &mesocycles); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Mesocycles not found "), nil
		}
		return Response{}, Error(err)
	}

	resource := resources.Mesocycles{}
	resources.Resources(mesocycles, &resource, resources.Mesocycle{}.New)

	return Response{}.SuccessWithData("Mesocycles found.", resource), nil
}

func (Mesocycle) Show(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	mesocycle := models.Mesocycle{Id: id}
	if err := models.SelectOne(db, &mesocycle); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Mesocycle not found."), nil
		}
		return Response{}, Error(err)
	}
	if err := models.SelectAllWhere(db, &mesocycle.Microcycles, "mesocycle_id", "=", mesocycle.Id); err != nil {
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Mesocycle found.", resources.MesocycleDetails{}.New(mesocycle)), nil
}

func (Mesocycle) Store(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	macrocycleId, err := strconv.Atoi(vars["macrocycleId"])
	if err != nil {
		return Response{}, Error(err)
	}

	createMesocycle := requests.CreateMesocycle{}
	requests.Read(request.Body, &createMesocycle)
	createMesocycle.MacrocycleId = macrocycleId
	valid, errors := createMesocycle.Validate()
	if !valid {
		return Response{}.ValidationFailed("", errors), nil
	}

	mesocycle := models.Mesocycle{}
	mesocycle.FromCreateRequest(createMesocycle)
	if err := models.Insert(db, &mesocycle); err != nil {
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Mesocycle created.", map[string]int{"id": mesocycle.Id}), nil
}

func (Mesocycle) Update(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	updateMesocycle := requests.UpdateMesocycle{Id: id}
	requests.Read(request.Body, &updateMesocycle)
	valid, validationErrors := updateMesocycle.Validate()
	if !valid {
		return Response{}.ValidationFailed("", validationErrors), nil
	}

	mesocycle := models.Mesocycle{}
	mesocycle.FromUpdateRequest(updateMesocycle)
	if err := models.Update(db, mesocycle); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Mesocycle not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Mesocycle updated."), nil
}

func (Mesocycle) Destroy(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	if err := models.Delete(db, models.Mesocycle{Id: id}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Mesocycle not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Mesocycle deleted."), nil
}
