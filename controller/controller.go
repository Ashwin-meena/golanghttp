package controller

import (
	"encoding/json"
	"fmt"
	"go-dummy/database"
	"go-dummy/helper"
	"go-dummy/model"
	"net/http"
     "go-dummy/response"
	"github.com/gorilla/mux"
	"errors"
)

// CreateUser create a user in the postgres db
func CreateUser(w http.ResponseWriter, r *http.Request) {
    // set the header to content type x-www-form-urlencoded
    // Allow all origin to handle cors issue
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    fmt.Printf("heelele")
    // create an empty user of type models.User
    var user model.User_detail

    // decode the json request to user
    err := json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        //log.Fatalf("Unable to decode the request body.  %v", err)
		errMsg := fmt.Sprintf("Unable to decode the request body.  %v", err)
		response.ERROR(w, http.StatusBadRequest, errors.New(errMsg))
    }
	database.Instance.Create(&user)
	response.JSON(w, http.StatusCreated, user)
}

func GetAllUser(w http.ResponseWriter, r *http.Request){
	 w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
	 // create an empty user of type models.User
    var user []model.User_detail
    // get all the users in the db
    database.Instance.Find(&user)
    w.WriteHeader(http.StatusOK)
	fmt.Println(user)
    // send all the users as response
	response.JSON(w, http.StatusOK, user)
}
func GetUser(w http.ResponseWriter, r *http.Request){
	userId := mux.Vars(r)["id"]
	if checkIfuserExists(userId) == false {
		response.ERROR(w, http.StatusBadRequest, errors.New("User Not Found!"))
	}
	var user model.User_detail
	database.Instance.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	response.JSON(w, http.StatusOK, user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	if checkIfuserExists(userId) == false {
		response.ERROR(w, http.StatusBadRequest, errors.New("User Not Found!"))
	}
	var user model.User_detail
	database.Instance.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	response.JSON(w, http.StatusOK, user)
}
func Login(w http.ResponseWriter, r *http.Request){
	 w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
	var name =r.URL.Query().Get("name")
	var pass =r.URL.Query().Get("password")
	 // create an empty user of type models.User
    var user model.User_detail
    // get all the users in the db
    database.Instance.First(&user,"user_name = ? and password = ?",name,pass)
	if(user.ID !=0){
	token,err := helper.CreateToken(uint32(user.ID));
	if err != nil{
		response.ERROR(w, http.StatusBadRequest, errors.New("Unable to make token"))
		//log.Fatalf("Unable to make token.  %v", err)
	}
     w.Header().Set("Authorization",token)
    // send all the users as response
	response.JSON(w, http.StatusOK, "Success")
}else {
	response.ERROR(w, http.StatusBadRequest, errors.New("Name Or password is not correct"))
}
}
func checkIfuserExists(userId string) bool {
	var user model.User_detail
	database.Instance.First(&user, userId)
	if user.ID == 0 {
		return false
	}
	return true
}

