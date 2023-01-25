package user

type UserLogin struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}
