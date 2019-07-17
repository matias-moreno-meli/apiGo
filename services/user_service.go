package services

import (
	"../domains"
	"../utils"
)

func GetUserFromAPI(userID int64) (*domains.User, *utils.ApiError) {

	user := &domains.User{
		ID: userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}
