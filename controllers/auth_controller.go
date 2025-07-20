package controllers

import (
	"fmt"
	"net/http"

	"github.com/ashunasar/go-jwt-auth-api/database"
	"github.com/ashunasar/go-jwt-auth-api/middleware"
	"github.com/ashunasar/go-jwt-auth-api/models"
	"github.com/ashunasar/go-jwt-auth-api/utils"
	"github.com/google/uuid"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	body, ok := middleware.GetRequestBody[models.SignUpBody](r)
	if !ok {
		return
	}

	user := models.User{Email: body.Email, Name: body.Name, Password: body.Password}

	existingUserId, _, _ := database.GetUserByEmail(user.Email)

	if existingUserId != uuid.Nil {
		utils.WriteJson(w, http.StatusOK, utils.GeneralError(fmt.Errorf("user with email Id %s already exist ", user.Email)))
		return

	}

	password, err := utils.HashPassword(user.Password)

	if err != nil {
		return

	}
	user.Password = password

	id, err := database.CreateUser(user)

	if err != nil {

		utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
		return
	}

	accessToken, err := utils.SignAccessToken(id)

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
		return
	}

	refreshToken, err := utils.SignRefreshToken(id)

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
		return
	}

	utils.WriteJson(w, http.StatusOK, utils.GeneralResponse(map[string]any{
		"id":           id,
		"name":         user.Name,
		"email":        user.Email,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}))

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	body, ok := middleware.GetRequestBody[models.LoginBody](r)
	if !ok {
		return
	}

	userId, password, _ := database.GetUserByEmail(body.Email)

	if userId == uuid.Nil {
		utils.WriteJson(w, http.StatusOK, utils.GeneralError(fmt.Errorf("user with email Id %s is not registered", body.Email)))
		return

	}

	isCorrectPassword := utils.ComparePassword(password, body.Password)

	if !isCorrectPassword {
		utils.WriteJson(w, http.StatusOK, utils.GeneralError(fmt.Errorf("incorrect password")))
		return

	}

	accessToken, err := utils.SignAccessToken(userId)

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
		return
	}

	refreshToken, err := utils.SignRefreshToken(userId)

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
		return
	}

	utils.WriteJson(w, http.StatusOK, utils.GeneralResponse(map[string]any{
		"id":           userId,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}))

}
