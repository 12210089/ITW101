package userservices

import "prj101/model/userdomain"

func UserRegCreate(user userdomain.User) error {
	if err := user.UserRegSave(); err != nil {
		return err
	}
	return nil

}

// login
func ReadUserEmail(email string) (*userdomain.User, error) {
	result := &userdomain.User{Email: email}
	if err := result.GetUserDetailsOFEmail(); err != nil {
		return nil, err
	}
	return result, nil
}
