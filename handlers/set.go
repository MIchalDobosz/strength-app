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

type Set struct{}

func (Set) Index(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	sets := models.Sets{}
	if err := models.SelectAll(db, &sets); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Sets not found "), nil
		}
		return Response{}, Error(err)
	}

	resource := resources.Sets{}
	resources.Resources(sets, &resource, resources.Set{}.New)

	return Response{}.SuccessWithData("Sets found.", resource), nil
}

func (Set) Show(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	set := models.Set{Id: id}
	if err := models.SelectOne(db, &set); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Set not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Set found.", resources.SetDetails{}.New(set)), nil
}

func (Set) Store(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	createSet := requests.CreateSet{}
	requests.Read(request.Body, &createSet)
	valid, errors := createSet.Validate()
	if !valid {
		return Response{}.ValidationFailed("", errors), nil
	}

	set := models.Set{}
	set.FromCreateRequest(createSet)
	if err := models.Insert(db, &set); err != nil {
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Set created.", map[string]int{"id": set.Id}), nil
}

func (Set) Update(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	updateSet := requests.UpdateSet{Id: id}
	requests.Read(request.Body, &updateSet)
	valid, validationErrors := updateSet.Validate()
	if !valid {
		return Response{}.ValidationFailed("", validationErrors), nil
	}

	set := models.Set{}
	set.FromUpdateRequest(updateSet)
	if err := models.Update(db, set); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Set not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Set updated."), nil
}

func (Set) Destroy(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	if err := models.Delete(db, models.Set{Id: id}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Set not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Set deleted."), nil
}
