package request

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserPassword struct {
	Password string `json:"password"`
}
