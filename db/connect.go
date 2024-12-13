package db

import "database/sql"

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "magento30"
	dbName   = "company_go"
)

func GetConnect() (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db, err
}
