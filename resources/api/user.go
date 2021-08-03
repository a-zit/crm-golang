package apiresource

type UserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name"`
}
