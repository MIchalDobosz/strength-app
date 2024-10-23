package models

import (
	"database/sql"
	"fmt"
	"slices"
	"strength-app/utils"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Model interface {
	Table() string
	PrimaryKeyColumn() string
	PrimaryKeyValue() int
}

func SelectOne(db *sqlx.DB, model Model) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", model.Table(), model.PrimaryKeyColumn())
	return db.Get(model, query, model.PrimaryKeyValue())
}

func SelectAll[S ~[]M, M Model](db *sqlx.DB, models *S) error {
	return SelectAllWhere(db, models, "", "", nil)
}

func SelectAllWhere[S ~[]M, M Model](db *sqlx.DB, models *S, column string, operator string, value any) error {
	model := *new(M)
	query := fmt.Sprintf("SELECT * FROM %s", model.Table())

	args := []any{}
	if column != "" && operator != "" && value != nil {
		if strings.ToUpper(operator) == "IN" {
			var inQuery string
			var err error
			inQuery, args, err = sqlx.In("WHERE %s IN (?)", column, value)
			if err != nil {
				return err
			}
			query += " " + inQuery
		} else {
			query += fmt.Sprintf(" WHERE %s %s ?", column, operator)
			args = append(args, value)
		}
	}

	if err := db.Select(models, query, args...); err != nil {
		return err
	}
	if len(*models) == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func Insert(db *sqlx.DB, model Model) error {
	tags, err := utils.StructTags(model, "db")
	if err != nil {
		return err
	}
	pk := slices.Index(tags, model.PrimaryKeyColumn())
	if pk != -1 {
		tags = slices.Delete(tags, pk, pk)
	}
	columns := strings.Join(tags, ", ")
	binds := strings.Join(tags, ", :")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (:%s)", model.Table(), columns, binds)
	result, err := db.NamedExec(query, model)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	if err := utils.SetStructField(model, model.PrimaryKeyColumn(), id); err != nil {
		return err
	}
	return nil
}

func Update(db *sqlx.DB, model Model) error {
	tags, err := utils.StructTags(model, "db")
	if err != nil {
		return err
	}
	pk := slices.Index(tags, model.PrimaryKeyColumn())
	if pk != -1 {
		tags = slices.Delete(tags, pk, pk)
	}
	sets := ""
	for _, tag := range tags {
		sets += fmt.Sprintf("%s = :%s, ", tag, tag)
	}
	sets = strings.TrimSuffix(sets, ", ")

	query := fmt.Sprintf("Update %s SET %s WHERE %s = :%s", model.Table(), sets, model.PrimaryKeyColumn(), model.PrimaryKeyColumn())
	result, err := db.NamedExec(query, model)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func Delete(db *sqlx.DB, model Model) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", model.Table(), model.PrimaryKeyColumn())
	result, err := db.Exec(query, model.PrimaryKeyValue())
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func Ids[S ~[]M, M Model](models S) []int {
	ids := []int{}
	for _, model := range models {
		ids = append(ids, model.PrimaryKeyValue())
	}
	return ids
}
