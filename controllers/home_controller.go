package controllers

import (
	"fmt"
	"net/http"

	"github.com/ashunasar/go-jwt-auth-api/middleware"
	"github.com/ashunasar/go-jwt-auth-api/utils"
)

func HomeHndler(w http.ResponseWriter, r *http.Request) {
	id, ok := middleware.GetUserID(r)
	if !ok {
		utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(fmt.Errorf("id is not present")))
		return
	}

	utils.WriteJson(w, http.StatusOK, utils.GeneralResponse("uer id is "+id.String()))

}
