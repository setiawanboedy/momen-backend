package users

type UserFormater struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user User, token string) UserFormater {
	formater := UserFormater{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}

	return formater
}
