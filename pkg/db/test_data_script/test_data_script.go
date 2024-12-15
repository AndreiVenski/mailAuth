package test_data_script

import (
	"database/sql"
	"io/ioutil"
)

func ExecuteSQLFile(db *sql.DB, filepath string) error {
	sqlContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlContent))
	if err != nil {
		return err
	}

	return nil
}
