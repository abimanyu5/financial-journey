package repository

import (
	"database/sql"
	"financial-journey/structs"
	"time"
)


func GetAllTransactions(db *sql.DB, id uint) ( results []structs.Transactions, err error) {
	sql := "SELECT * FROM transactions where user_id = $1"

	rows, err := db.Query(sql, id)
	if err != nil { return nil, err }

	defer rows.Close()

	for rows.Next() {
		var transactions = structs.Transactions{}
		err = rows.Scan(&transactions.ID, &transactions.MasterId, &transactions.GoalId, &transactions.UserId, &transactions.Amount, &transactions.Title, &transactions.Description, &transactions.CreatedAt, &transactions.UpdatedAt)
		if err != nil { return nil, err }
		
		results = append(results, transactions)
	}
	return
}

func InsertTransaction(db *sql.DB, transactions structs.Transactions) (err error) {
	sql := "INSERT INTO transactions (master_id,goal_id,user_id,amount,title,description,created_at,updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	
	errs := db.QueryRow(sql,transactions.MasterId, transactions.GoalId, transactions.UserId, transactions.Amount, transactions.Title, transactions.Description, time.Now(), time.Now())

	return errs.Err()
}

func UpdateTransaction(db *sql.DB, transactions structs.Transactions) (err error) {
	sql := "UPDATE transactions SET master_id=$1, goal_id=$2, user_id=$3, amount=$4, title=$5, description=$6, updated_at=$7 WHERE id=$8"
	
	errs := db.QueryRow(sql, transactions.MasterId, transactions.GoalId, transactions.UserId, transactions.Amount, transactions.Title, transactions.Description, time.Now(), transactions.ID)
	return errs.Err()
}

func DeleteTransaction(db *sql.DB, transactions structs.Transactions) (err error) {
	sql := "DELETE FROM transactions WHERE id = $1"
	
	errs := db.QueryRow(sql, transactions.ID)
	return errs.Err()
}
