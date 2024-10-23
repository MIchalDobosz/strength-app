package models

import (
	"database/sql"
	"strength-app/requests"
)

type Sessions []Session
type Session struct {
	Id            int            `db:"id"`
	MicrocycleId  int            `db:"microcycle_id"`
	Name          string         `db:"name"`
	PlannedDate   sql.NullString `db:"planned_date"`
	PerformedDate sql.NullString `db:"performed_date"`
	Completed     bool           `db:"completed"`
	Microcycle    Microcycle
	Slots         Slots
}

func (Session) Table() string {
	return "sessions"
}

func (Session) PrimaryKeyColumn() string {
	return "id"
}

func (session Session) PrimaryKeyValue() int {
	return session.Id
}

func (session *Session) FromCreateRequest(createSession requests.CreateSession) {
	session.MicrocycleId = createSession.MicrocycleId
	session.Name = createSession.Name
}

func (session *Session) FromUpdateRequest(updateSession requests.UpdateSession) {
	session.Id = updateSession.Id
	session.MicrocycleId = updateSession.MicrocycleId
	session.Name = updateSession.Name
}
