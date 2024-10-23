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

type Session struct{}

func (Session) Index(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	sessions := models.Sessions{}
	if err := models.SelectAll(db, &sessions); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Sessions not found "), nil
		}
		return Response{}, Error(err)
	}

	resource := resources.Sessions{}
	resources.Resources(sessions, &resource, resources.Session{}.New)

	return Response{}.SuccessWithData("Sessions found.", resource), nil
}

func (Session) Show(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	session := models.Session{Id: id}
	if err := models.SelectOne(db, &session); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Session not found."), nil
		}
		return Response{}, Error(err)
	}
	if err := session.Slots.LoadWithExercisesBySessionId(db, session.Id); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return Response{}, Error(err)
		}
	}

	return Response{}.SuccessWithData("Session found.", resources.SessionDetails{}.New(session)), nil
}

func (Session) Store(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	createSession := requests.CreateSession{}
	requests.Read(request.Body, &createSession)
	valid, errors := createSession.Validate()
	if !valid {
		return Response{}.ValidationFailed("", errors), nil
	}

	session := models.Session{}
	session.FromCreateRequest(createSession)
	if err := models.Insert(db, &session); err != nil {
		return Response{}, Error(err)
	}

	return Response{}.SuccessWithData("Session created.", map[string]int{"id": session.Id}), nil
}

func (Session) Update(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	updateSession := requests.UpdateSession{Id: id}
	requests.Read(request.Body, &updateSession)
	valid, validationErrors := updateSession.Validate()
	if !valid {
		return Response{}.ValidationFailed("", validationErrors), nil
	}

	session := models.Session{}
	session.FromUpdateRequest(updateSession)
	if err := models.Update(db, session); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Session not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Session updated."), nil
}

func (Session) Destroy(request *http.Request, vars map[string]string, db *sqlx.DB) (Response, error) {
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Response{}, Error(err)
	}

	if err := models.Delete(db, models.Session{Id: id}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}.NotFound("Session not found."), nil
		}
		return Response{}, Error(err)
	}

	return Response{}.Success("Session deleted."), nil
}
