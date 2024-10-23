package resources

import (
	"strength-app/models"
	"strings"
)

type Sessions []Session
type Session struct {
	Id            int    `json:"id"`
	MicrocycleId  int    `json:"microcycle_id"`
	Name          string `json:"name"`
	PlannedDate   string `json:"planned_date"`
	PerformedDate string `json:"performed_date"`
	Completed     bool   `json:"completed"`
}

func (Session) New(session models.Session) Session {
	return Session{
		Id:            session.Id,
		MicrocycleId:  session.MicrocycleId,
		Name:          session.Name,
		PlannedDate:   strings.Split(session.PlannedDate.String, "T")[0],
		PerformedDate: strings.Split(session.PerformedDate.String, "T")[0],
		Completed:     session.Completed,
	}
}

type SessionDetails struct {
	Session
	Slots Slots `json:"slots"`
}

func (SessionDetails) New(model models.Session) SessionDetails {
	details := SessionDetails{
		Session: Session{}.New(model),
	}
	Resources(model.Slots, &details.Slots, Slot{}.New)
	return details
}
