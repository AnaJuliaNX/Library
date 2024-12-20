package presenter

type Users struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Owner bool   `json:"owner"`
	Phone string `json:"phone"`
}
