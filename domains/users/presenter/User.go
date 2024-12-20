package presenter

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Owner bool   `json:"owner"`
	Phone string `json:"phone"`
	Meta  Meta
}

type Meta struct {
	Token        string `json:"token"`
	TokenType    string `json:"token_type"`
	TokenExpired int64  `json:"token_expires_in"`
}

func (u *User) SetMeta(token, tokenTipe string, tokenExpired int64) {
	u.Meta.Token = token
	u.Meta.TokenType = tokenTipe
	u.Meta.TokenExpired = tokenExpired
}
