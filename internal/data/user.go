package data

type UserData struct {
	*Data
}

func NewUserData(data *Data) *UserData {
	return &UserData{
		Data: data,
	}
}
