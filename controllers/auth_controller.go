package controllers

import (
	"net/http"

	"github.com/ashunasar/go-jwt-auth-api/middleware"
	"github.com/ashunasar/go-jwt-auth-api/models"
	"github.com/ashunasar/go-jwt-auth-api/utils"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	body, ok := middleware.GetRequestBody[models.SignUpBody](r)
	if !ok {
		return
	}

	utils.WriteJson(w, http.StatusOK, utils.GeneralResponse(body))

}
