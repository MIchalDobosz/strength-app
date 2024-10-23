package requests

type CreateSession struct {
	MicrocycleId  int    `db:"microcycle_id"`
	Name          string `db:"name"`
	PlannedDate   string `db:"planned_date"`
	PerformedDate string `db:"performed_date"`
	Completed     bool   `db:"completed"`
}

func (createSession CreateSession) Validate() (bool, map[string]string) {
	valid := true
	errors := map[string]string{}
	validateRequired("microcycle_id", createSession.MicrocycleId, &valid, errors)
	validateRequired("name", createSession.Name, &valid, errors)
	return valid, errors
}

type UpdateSession struct {
	Id int `json:"id"`
	CreateSession
}

func (updateSession UpdateSession) Validate() (bool, map[string]string) {
	valid, errors := updateSession.CreateSession.Validate()
	validateRequired("id", updateSession.Id, &valid, errors)
	return valid, errors
}
