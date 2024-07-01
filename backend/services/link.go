package services

import (
	"errors"
	"github.com/fadegor05/linkshortee/backend/database"
	"github.com/fadegor05/linkshortee/backend/models"
	"github.com/fadegor05/linkshortee/backend/types"
	"github.com/fadegor05/linkshortee/backend/utils"
	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid/v3"
	"log"
	"net/http"
)

func (h *Handler) handleCreateLink(w http.ResponseWriter, r *http.Request) {
	var payload types.LinkPostRequest
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	newLink := models.Link{
		Code:     shortuuid.New(),
		FullLink: payload.FullLink,
	}

	database.DB.Create(&newLink)

	response := types.LinkPostResponse{
		Code: newLink.Code,
	}

	err := utils.WriteJSON(w, http.StatusCreated, response)
	if err != nil {
		log.Panic(err)
	}
}

func (h *Handler) handleGetLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var code string
	if v, ok := vars["code"]; ok {
		code = v
	} else {
		utils.WriteError(w, http.StatusNotFound, errors.New("code not provided"))
	}

	var link models.Link
	database.DB.Where(&models.Link{Code: code}).First(&link)
	if link.Code == "" {
		utils.WriteError(w, http.StatusNotFound, errors.New("link not found"))
	}

	http.Redirect(w, r, link.FullLink, http.StatusPermanentRedirect)
}
