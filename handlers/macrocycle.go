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

type Macrocycle struct{}

func (Macrocycle) Index(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	macrocycles := models.Macrocycles{}
	if err := models.SelectAll(db, &macrocycles); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Macrocycles not found "), nil
		}
		return Response{}, Error(err)
	}

	resource := resources.Macrocycles{}
	resources.Resources(macrocycles, &resource, resources.Macrocycle{}.New)

	return Response{}.SuccessWithData("Macrocycles found.", resource), nil
}

func (Macrocycle) Show(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	macrocycle := models.Macrocycle{Id: id}
	if err := models.SelectOne(db, &macrocycle); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Macrocycle not found."), nil
		}
		return Response{}, Error(err)
	}
	if err := models.SelectAllWhere(db, &macrocycle.Mesocycles, "macrocycle_id", "=", macrocycle.Id); err != nil {
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Macrocycle found.", resources.MacrocycleDetails{}.New(macrocycle)), nil
}

func (Macrocycle) Store(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	createMacrocycle := requests.CreateMacrocycle{}
	requests.Read(request.Body, &createMacrocycle)
	valid, errors := createMacrocycle.Validate()
	if !valid {
		return Response{}.ValidationFailed("", errors), nil
	}

	macrocycle := models.Macrocycle{}
	macrocycle.FromCreateRequest(createMacrocycle)
	if err := models.Insert(db, &macrocycle); err != nil {
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Macrocycle created.", map[string]int{"id": macrocycle.Id}), nil
}

func (Macrocycle) Update(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	updateMacrocycle := requests.UpdateMacrocycle{Id: id}
	requests.Read(request.Body, &updateMacrocycle)
	valid, validationErrors := updateMacrocycle.Validate()
	if !valid {
		return Response{}.ValidationFailed("", validationErrors), nil
	}

	macrocycle := models.Macrocycle{}
	macrocycle.FromUpdateRequest(updateMacrocycle)
	if err := models.Update(db, macrocycle); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Macrocycle not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Macrocycle updated."), nil
}

func (Macrocycle) Destroy(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	if err := models.Delete(db, models.Macrocycle{Id: id}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Macrocycle not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Macrocycle deleted."), nil
}
