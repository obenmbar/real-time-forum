package forum



type Users struct {
	Nickname         string `json:"nickname"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Age              string `json:"age"`
	Gender           string `json:"gender"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Confirm_password string `json:"confirm-password"`
}
type Login struct {
	Nicknameoremail string `json:user`
	Password  string `json:password`
}
	