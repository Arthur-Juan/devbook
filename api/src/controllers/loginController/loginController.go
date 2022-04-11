package loginController

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/services/passwordServices"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewUserRepository(db)
	userExists, err := userRepository.SearchUserByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = passwordServices.VerifyPassword(userExists.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GenerateToken(userExists.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}
	w.Write([]byte(token))

}
