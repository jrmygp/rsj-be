package responses

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserLoginResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	UserRole string `json:"user_role"`
	Token    string `json:"token"`
}
