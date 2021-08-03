package apiservice

import (
	"crm-golang/models"
	apiresource "crm-golang/resources/api"
	"crm-golang/validator/user"
)

type UserService struct {
	User models.User
}

func (us *UserService) UserList() (interface{}, error) {
	result, err := us.User.UserList()

	if err != nil {
		return nil, err
	}

	var userResponses []apiresource.UserResponse
	users := result.([]models.User)

	for _, v := range users {
		userResponse := apiresource.UserResponse{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (us *UserService) CreateUser(userRequest user.CreateUserRequest) (interface{}, error) {

	us.User.FirstName = userRequest.FirstName
	us.User.LastName = userRequest.LastName

	err := us.User.CreateUser()

	if err != nil {
		return nil, err
	}

	userResponse := apiresource.UserResponse{
		ID:        us.User.ID,
		FirstName: us.User.FirstName,
		LastName:  us.User.LastName,
	}

	return userResponse, nil
}
