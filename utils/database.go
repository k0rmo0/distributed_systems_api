package utils

import (

	//Needs to be imported, or it will cause an error .
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//DBAccess ...
type DBAccess struct {
	SQLDB *sqlx.DB
}

// SQLAcc ...
var SQLAcc DBAccess

//GetSQLDB ...
func GetSQLDB() error {
	dbDriver := "mysql"
	dbUser := "dsa"
	dbPass := "dsa2020"
	dbName := "supersonic"
	dbAddress := "localhost"
	dbPort := "3306"
	db, err := sqlx.Connect(dbDriver, dbUser+":"+dbPass+"@"+"tcp("+dbAddress+":"+dbPort+")/"+dbName)

	if err != nil {
		return err
	}

	SQLAcc.SQLDB = db

	return err
}

//GetSQLDB method ...
func (a DBAccess) GetSQLDB() *sqlx.DB {
	return a.SQLDB
}
