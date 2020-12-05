package model

import (
	"database/sql"
	"errors"

	"github.com/ismar/dsa/distrybuted_systems_api/utils"
)

// Values ...
type Values struct {
	ID       int    `db:"id" json:"id"`
	Distance int    `db:"distance" json:"distance"`
	Time     string `db:"time" json:"time"`
}

// WriteToDataBase ...
func (val *Values) WriteToDataBase() error {
	db := utils.SQLAcc.GetSQLDB()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO distances(distance, time) values (?, ?)", val.Distance, val.Time)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// ListValues ...
func ListValues() ([]Values, error) {
	db := utils.SQLAcc.GetSQLDB()
	var rows []Values
	query := "SELECT * FROM distances"

	err := db.Select(&rows, query)

	if err == sql.ErrNoRows {
		return rows, errors.New("Cant't get users")
	}

	if err != nil {
		return rows, err
	}

	return rows, nil
}
