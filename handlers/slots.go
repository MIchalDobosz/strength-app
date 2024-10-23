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

type Slot struct{}

func (Slot) Index(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	slots := models.Slots{}
	if err := slots.LoadWithExercises(db); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Slots not found "), nil
		}
		return Response{}, Error(err)
	}

	resource := resources.Slots{}
	resources.Resources(slots, &resource, resources.Slot{}.New)

	return Response{}.SuccessWithData("Slots found.", resource), nil
}

func (Slot) Show(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	slot := models.Slot{Id: id}
	if err := slot.LoadWithExercises(db); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Slot not found."), nil
		}
		return Response{}, Error(err)
	}
	if err := models.SelectAllWhere(db, &slot.Sets, "slot_id", "=", slot.Id); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return Response{}, Error(err)
		}
	}

	return Response{}.SuccessWithData("Slot found.", resources.SlotDetails{}.New(slot)), nil
}

func (Slot) Store(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	createSlot := requests.CreateSlot{}
	requests.Read(request.Body, &createSlot)
	valid, errors := createSlot.Validate()
	if !valid {
		return Response{}.ValidationFailed("", errors), nil
	}

	slot := models.Slot{}
	slot.FromCreateRequest(createSlot)
	if err := models.Insert(db, &slot); err != nil {
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Slot created.", map[string]int{"id": slot.Id}), nil
}

func (Slot) Update(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	updateSlot := requests.UpdateSlot{Id: id}
	requests.Read(request.Body, &updateSlot)
	valid, validationErrors := updateSlot.Validate()
	if !valid {
		return Response{}.ValidationFailed("", validationErrors), nil
	}

	slot := models.Slot{}
	slot.FromUpdateRequest(updateSlot)
	if err := models.Update(db, slot); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Slot not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Slot updated."), nil
}

func (Slot) Destroy(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	if err := models.Delete(db, models.Slot{Id: id}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Slot not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Slot deleted."), nil
}
