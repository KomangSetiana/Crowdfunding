package user

type UserFormater struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

func UserFormat(user User, token string) UserFormater {
	formatter := UserFormater{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}

	return formatter
}
