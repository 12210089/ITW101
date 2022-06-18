package userdomain

import (
	"prj101/datasource/postgres"
)

type User struct {
	UserName     string
	Email        string
	UserPassword string
}

const (
	queryInsertUser = "INSERT INTO UserData(UserName, Email, UserPassword) VALUES($1,$2, $3)"
	querygetuser    = "SELECT Email, UserPassword from userdata where Email = $1"
)

func (user *User) UserRegSave() error {
	stmt, err := postgres.Client.Prepare(queryInsertUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, saveErr := stmt.Exec(user.UserName, user.Email, user.UserPassword)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

// login
func (user *User) GetUserDetailsOFEmail() error {
	stmt, err := postgres.Client.Prepare(querygetuser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email)
	if getErr := result.Scan(&user.Email, &user.UserPassword); getErr != nil {
		return getErr
	}
	return nil
}
