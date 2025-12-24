package user

type Repository interface {
	GetUserByID(id string) (*User, error)
}
