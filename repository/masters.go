package repository

import (
	"database/sql"
	"financial-journey/structs"
	"time"
)


func GetAllMasters(db *sql.DB, id uint) (err error, results []structs.Masters) {
	sql := "SELECT * FROM masters where user_id = $1"

	rows, err := db.Query(sql, id)
	if err != nil { return err, nil }

	defer rows.Close()

	for rows.Next() {
		var masters = structs.Masters{}
		err = rows.Scan(&masters.ID, &masters.Name, &masters.Description, &masters.UserId, &masters.IncomeType, &masters.CreatedAt, &masters.UpdatedAt)
		if err != nil { return err, nil }
		
		results = append(results, masters)
	}
	return
}

func InsertMaster(db *sql.DB, masters structs.Masters) (err error) {
	sql := "INSERT INTO masters ( name, description, user_id, income_type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	errs := db.QueryRow(sql, masters.Name, masters.Description, masters.UserId, masters.IncomeType, time.Now(), time.Now())

	return errs.Err()
}

func UpdateMaster(db *sql.DB, masters structs.Masters) (err error) {
	sql := "UPDATE masters SET name=$1, description = $2, income_type = $3, updated_at=$4 WHERE id=$5"
	
	errs := db.QueryRow(sql, masters.Name, masters.Description, masters.IncomeType, time.Now(), masters.ID)
	return errs.Err()
}

func DeleteMaster(db *sql.DB, masters structs.Masters) (err error) {
	sql := "DELETE FROM masters WHERE id = $1"
	
	errs := db.QueryRow(sql, masters.ID)
	return errs.Err()
}

func GetMasterById(db *sql.DB, masters structs.Masters) (err error, results []structs.Masters) {
	sql := "SELECT * FROM masters WHERE id = $1"

	rows, err := db.Query(sql, masters.ID)
	if err != nil { return err, nil }

	defer rows.Close()

	for rows.Next() {
		var masters = structs.Masters{}
		err = rows.Scan(&masters.ID, &masters.Name, &masters.Description, &masters.UserId, &masters.IncomeType, &masters.CreatedAt, &masters.UpdatedAt)
		if err != nil { return err, nil }
		
		results = append(results, masters)
	}
	return
}
