package domain

type (
	User struct {
		ID   int
		Name string
	}

	UserRepository interface {
		FindByName(name string) (*User, error)
	}
)
