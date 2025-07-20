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

	existingUserId, _ := database.GetUseByEmail(user.Email)

	if existingUserId != uuid.Nil {
		utils.WriteJson(w, http.StatusOK, utils.GeneralError(fmt.Errorf("user with email Id %s already exist ", user.Email)))
		return

	}

	id, err := database.CreateUser(user)
	if err != nil {

		utils.WriteJson(w, http.StatusOK, utils.GeneralError(err))
		return
	}

	utils.WriteJson(w, http.StatusOK, utils.GeneralResponse(map[string]any{
		"id":       id,
		"name":     body.Name,
		"email":    body.Email,
		"password": body.Password,
	}))

}
