package user

type User struct {
	id        int    `json:"id"`
	firstName string `json:"first_name"`
	lastName  string `json:"last_name"`
}
