package request

type RequestUpdate struct {
	FullName string `json:"fullName,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required"`
}
