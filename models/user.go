package models

import (
	"database/sql"
)

type (
	User struct {
		ID			int 	`json:"id"`
		Name		string	`json:"name" validate:"required"`
		Email		string	`json:"email" validate:"required,email"`
		Password	string	`json:"password" validate:"required"`
	}
)

func CreateUser(db *sql.DB, name string, email string, password string) (int64, error) {
	sql := "INSERT INTO users(name,email,password) VALUES(?,?,?)"
	
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name, email, password)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func ShowUserByEmail(db *sql.DB, email string) User  {
	rows, err := db.Query("SELECT * FROM users where email=?", email)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := User{}
	for rows.Next() {
		user := User{}	
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
        if err != nil {
            panic(err.Error())
		}
		
        result = user
	}

	return result
}