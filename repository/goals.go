package repository

import (
	"database/sql"
	"errors"
	"financial-journey/structs"
	"time"
)


func GetAllGoals(db *sql.DB, id uint) ( results []structs.Goals, err error) {
	sql := "SELECT * FROM goals where id_user = $1"

	rows, err := db.Query(sql, id)
	if err != nil { return nil, err }

	defer rows.Close()

	for rows.Next() {
		var goals = structs.Goals{}
		err = rows.Scan(&goals.ID, &goals.UserId, &goals.Amount, &goals.AmountGoal, &goals.Name, &goals.Description, &goals.CreatedAt, &goals.UpdatedAt)
		if err != nil { return nil, err }
		
		results = append(results, goals)
	}
	return
}

func InsertGoals(db *sql.DB, goals structs.Goals) (err error) {
	sql := "INSERT INTO goals ( id_user, amount, amount_goal, name, description, created_at ,updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	errs := db.QueryRow(sql, goals.UserId, goals.Amount, goals.AmountGoal, goals.Name, goals.Description, time.Now(), time.Now())

	return errs.Err()
}

func UpdateGoals(db *sql.DB, goals structs.Goals) (err error) {
	sql := "UPDATE masters SET amount = $1, amount_goal=$1, name=$3, description=$4, updated_at=$5 WHERE id=$6"
	
	errs := db.QueryRow(sql, goals.Amount, goals.AmountGoal, goals.Name, goals.Description, time.Now(), goals.ID)
	return errs.Err()
}

func DeleteGoals(db *sql.DB, goals structs.Goals) (err error) {
	sql := "DELETE FROM goals WHERE id = $1"
	
	errs := db.QueryRow(sql, goals.ID)
	return errs.Err()
}

func GetGoalsById(db *sql.DB, goals structs.Goals) (results []structs.Goals, err error) {
	sql := "SELECT * FROM goals WHERE id = $1"

	rows, err := db.Query(sql, goals.ID)
	if err != nil { return nil, err }

	defer rows.Close()

	for rows.Next() {
		var goals = structs.Goals{}
		err = rows.Scan(&goals.ID, &goals.UserId, &goals.Amount, &goals.AmountGoal, &goals.Name, &goals.Description, &goals.CreatedAt, &goals.UpdatedAt)
		if err != nil { return nil, err }
		
		results = append(results, goals)
	}
	return
}

func CheckGoals(db *sql.DB, id int) (*structs.Goals, error) {
	var goal structs.Goals

	row := db.QueryRow("SELECT id, id_user, amount, amount_goal, name, description, created_at, updated_at FROM goals WHERE id = $1", id)
	err := row.Scan(&goal.ID, &goal.UserId, &goal.Amount, &goal.AmountGoal, &goal.Name, &goal.Description, &goal.CreatedAt, &goal.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil 
		}
		return nil, err 
	}

	return &goal, nil
}

func UpdateGoal(db *sql.DB, goal *structs.Goals) error {
	_, err := db.Exec("UPDATE goals SET id_user = $1, amount = $2, amount_goal = $3, name = $4, description = $5, updated_at = NOW() WHERE id = $6",
		goal.UserId, goal.Amount, goal.AmountGoal, goal.Name, goal.Description, goal.ID)
	if err != nil {
		return err
	}

	return nil
}