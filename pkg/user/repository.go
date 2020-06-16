package user

//Repository is an interface to separate impl of different db func
type Repository interface {
	FindByEmail(email string) (*User, error)
	FindByID(id uint64) (*User, error)
	Register(user *User) (*User, error)
	DoesEmailExist(email string) bool
}
