package helper

import "github.com/SuyunovJasurbek/CrudTask/models"

func UserCreateFeildCheck( user models.UserHandler)(models.Error, bool){
	if len(user.FullName) < 3 || len(user.FullName) > 20 {
		 return models.Error{Message: "Fullname must be between 3 and 20 characters"}, false
	}else if len(user.NickName) < 3 || len(user.NickName) > 20 {
		return models.Error{Message: "Nickname must be between 3 and 20 characters"}, false
	}else if len(user.Birthday) < 3 || len(user.Birthday) > 20 {
		return models.Error{Message: "Birthday must be between 3 and 20 characters"}, false
	}else if len(user.Location) < 5 || len(user.Location) > 32 {
		return models.Error{Message: "Location must be between 3 and 20 characters"}, false 
	}else if len(user.Photo) < 5 || len(user.Photo) > 35 {
		return models.Error{Message: "Photo must be between 3 and 20 characters"}, false 
	} 
	return models.Error{}, true
}