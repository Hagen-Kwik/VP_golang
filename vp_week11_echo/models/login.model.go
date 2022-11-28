package models

import (
	"database/sql"
	"fmt"
	"vp_week11_echo/db"
	"vp_week11_echo/helpers"
)

type User struct {
	id       int    `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
	email    string `json:"email"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"
	err := con.QueryRow(sqlStatement, username).Scan(&obj.id, &obj.username, &pwd, &obj.email)

	if err == sql.ErrNoRows {
		fmt.Print("Username Not Found!")
		return false, err
	}

	if err != nil {
		fmt.Print("Query Error!")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)

	if !match {
		fmt.Print("Hash and password doesn't match!")
		return false, err
	}

	return true, nil
}
