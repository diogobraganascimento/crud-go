package model

type userDomain struct {
	id       string
	Email    string `json:"email"`
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}
