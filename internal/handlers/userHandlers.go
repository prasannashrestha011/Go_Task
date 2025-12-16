package handlers

import (
	"encoding/json"
	"main/internal/models"
	"main/internal/services"
	"main/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserService interface {
	GET_USER(w http.ResponseWriter, r *http.Request)
	GET_ALL_USER(w http.ResponseWriter, r *http.Request)
	POST(w http.ResponseWriter, r *http.Request)
	PUT(w http.ResponseWriter, r *http.Request)
	DELETE(w http.ResponseWriter, r *http.Request)
}
type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserService {
	return &userHandler{userService:userService}

}



//handlers

func (u *userHandler) GET_ALL_USER(w http.ResponseWriter, r *http.Request) {
	var users []*models.UserResponseDTO
	users,err:=u.userService.GetUsers()
	if err!=nil{
		http.Error(w,"User list is empty",400)
		return 
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(users)
}

func (u *userHandler) GET_USER(w http.ResponseWriter, r *http.Request) {
	id:=chi.URLParam(r,"id")
	user,err:=u.userService.GetUserByID(uuid.MustParse(id))
	if err!=nil{
		http.Error(w,err.Error(),404)
		return 
	}
	
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func (u *userHandler) DELETE(w http.ResponseWriter, r *http.Request) {
	id:=chi.URLParam(r,"id")
	err:=u.userService.DeleteUser(uuid.MustParse(id))
	if err!=nil{
		err:=&utils.CustomError{
			Message: "User not found",
			Status: false,
		}
		http.Error(w,err.Error(),404)
		return
	}
	response:=map[string]string{
		"message":"user deleted with ID: "+id,
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(&response)
	
}


func (u *userHandler) POST(w http.ResponseWriter, r *http.Request) {
	var new_user models.UserCreateDTO 
	if err:=json.NewDecoder(r.Body).Decode(&new_user);err!=nil{
		http.Error(w, err.Error(), 400)
		return
	}
	user_details,err:=u.userService.RegisterUser(&new_user)
	if err!=nil{
		err:=&utils.CustomError{
			Message:"Email address  exists already",
			Status: false,
		}
		http.Error(w,err.Error(),500)
		return
	}
	w.Header().Set("Content-Type","applicatoin/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user_details)
}

func (u *userHandler) PUT(w http.ResponseWriter, r *http.Request) {
	var userDetails models.UserUpdateDTO
	id,err:=uuid.Parse(chi.URLParam(r,"id"))
	if err != nil{
		http.Error(w,"Invalid user id",400)
		return
	}
	if err:=json.NewDecoder(r.Body).Decode(&userDetails);err!=nil{
		http.Error(w,"Invalid user details",400)
		return
	}
	 updated_details,err:=u.userService.UpdateUser(id,&userDetails);
	 if err!=nil{
		http.Error(w,"User Update error: "+err.Error(),500)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(updated_details)
}

