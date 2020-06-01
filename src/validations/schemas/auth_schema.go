package schemas

// Register schema
type Register struct {
	Name     string `validate:"required"`
	UserName string `validate:"required,min=3,max=40,alphanum"`
	Password string `validate:"required,min=6"`
	Email    string `validate:"omitempty,email"`
}

// Login schema
type Login struct {
	UserName string `validate:"required,min=3,max=40,alphanum"`
	Password string `validate:"required,min=6"`
}

// RefreshToken schema
type RefreshToken struct {
	RefreshToken string `validate:"required" json:"refresh_token"`
}
