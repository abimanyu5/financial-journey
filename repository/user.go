package repository

import (
	"database/sql"
	"errors"
	"financial-journey/structs"
	"fmt"
	"time"
)

func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users (name, password, email, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	errs := db.QueryRow(sql, user.Username, user.Password, user.Email, user.Role, time.Now(), time.Now())

	return errs.Err()
}

func GetUserByName(db *sql.DB, user structs.User) (results []structs.User) {
	sql := "SELECT * FROM users WHERE name = $1"
	fmt.Println("test user",user)
	rows, err := db.Query(sql, user.Username)
	if err != nil {
		fmt.Println("error",err)
		return results
	}

	defer rows.Close()
	fmt.Println("result1",rows)

	for rows.Next() {
		var user = structs.User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return results
		}

		results = append(results, user)
		fmt.Println(results)
	}
	return
}
func GetUserByUsername(db *sql.DB, username string) (*structs.User, error) {
	user := &structs.User{}
	row := db.QueryRow("SELECT id, name, password, role FROM users WHERE name = $1", username)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}