package storage

// User is list of user data
type UserInfo struct {
	IDUser    string `json:"id_user,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreateAt  string `json:"create_at,omitempty"`
	ChangeAt  string `json:"change_at,omitempty"`
	LastLogin string `json:"last_login,omitempty"`
	Level     string `json:"level,omitempty"`
}

type User struct {
	Info   UserInfo
	Salt   string
	Papper string
}
