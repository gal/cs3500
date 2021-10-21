package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/Strum355/log"
	"github.com/gal/timber/models"
	"github.com/gal/timber/utils"
	"github.com/go-chi/chi"
)

type UserHandler struct {
	UserInfo models.UserStore
	UserAuth models.AuthStore
}

func NewUserHandler(uStore models.UserStore, aStore models.AuthStore) *UserHandler {
	return &UserHandler{uStore, aStore}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondJSON(w, nil, "error", "invalid id", http.StatusBadRequest)
		log.WithContext(r.Context()).WithError(err).
			Info("failed to convert string to int")
		return
	}
	if !utils.HasAccess(r, id) {
		utils.RespondJSON(w, nil, "error",
			"unauthorized for requested content", http.StatusUnauthorized,
		)
		log.WithContext(r.Context()).Info(fmt.Sprintf("unauthorized for User %d", id))
		return
	}

	user := &models.User{ID: id}

	utils.RespondJSON(w, user, "success", "", http.StatusOK)
	log.WithContext(r.Context()).Info("Served get request for user")
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var loginDetails *models.LoginDetails
	var u *models.User

	if err := json.NewDecoder(r.Body).Decode(
		&loginDetails,
	); err != nil {
		utils.RespondJSON(w, nil, "err",
			"invalid request", http.StatusBadRequest,
		)

		log.WithContext(r.Context()).Info("invalid login request")
		return
	}

	if loginDetails.Email == "" && loginDetails.Username == "" ||
		loginDetails.Password == "" {
		utils.RespondJSON(w, nil, "err",
			"invalid request", http.StatusBadRequest,
		)

		log.WithContext(r.Context()).Info("invalid login request")
		return
	}

	u = &models.User{
		Username: loginDetails.Username,
		Email:    loginDetails.Email,
	}

	if err := h.UserHandler.UserInfo.GetByEmail(u.Email); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := h.UserHandler.UserInfo.Create(u); err != nil {
				utils.RespondJSON(w, nil, "error",
					"error creating user", http.StatusInternalServerError,
				)

				log.WithContext(r.Context()).
					WithError(err).Info("error creating user")

				return
			}
			// TODO make sure to delete user object if userauth fails
			hashed, err := utils.HashPassword([]byte(loginDetails.Password))
			if err != nil {
				utils.RespondJSON(w, nil, "error",
					"error creating user", http.StatusInternalServerError,
				)

				log.WithContext(r.Context()).
					WithError(err).Info("error creating password hash")

				return
			}

			if u.ID == 0 {
				utils.RespondJSON(w, nil, "error",
					"error creating user", http.StatusInternalServerError,
				)

				log.WithContext(r.Context()).
					Info("error creating user")

				return
			}

			userAuth := &models.UserAuth{
				ID:      u.ID,
				Email:   u.Email,
				Hash:    hashed,
				Enabled: false,
			}

			if err = h.UserHandler.UserAuth.Create(userAuth); err != nil {
				utils.RespondJSON(w, nil, "error",
					"error creating user", http.StatusInternalServerError,
				)

				log.WithContext(r.Context()).WithError(err).
					Info("error creating userAuth")
				// TODO delete user object
				return
			}

			utils.RespondJSON(w, u, "success", "", http.StatusCreated)
		} else {
			utils.RespondJSON(w, nil, "error",
				"error creating user", http.StatusInternalServerError,
			)

			log.WithContext(r.Context()).
				WithError(err).Info("error creating user")

			return
		}
	}

	utils.RespondJSON(w, nil, "error",
		"user with email already exists", http.StatusConflict,
	)
	log.WithContext(r.Context()).
		Info("user with email already exists")
}

func (h *Handler) PatchUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondJSON(w, nil, "error", "invalid id", http.StatusBadRequest)
		log.WithContext(r.Context()).WithError(err).
			Info("failed to convert string to in")
		return
	}

	if !utils.HasAccess(r, id) {
		utils.RespondJSON(w, nil, "error",
			"unauthorized for requested content", http.StatusUnauthorized,
		)
		log.WithContext(r.Context()).Info(fmt.Sprintf("unauthorized for User %d", id))
		return
	}

	var u *models.User

	if err := json.NewDecoder(r.Body).Decode(
		&u,
	); err != nil {
		utils.RespondJSON(w, nil, "err",
			"invalid request", http.StatusBadRequest,
		)

		log.WithContext(r.Context()).WithError(err).Info("invalid PATCH request")
		return
	}

	u.ID = id

	if err := h.UserHandler.UserInfo.Patch(u); err != nil {
		utils.RespondJSON(w, nil, "error",
			"error updating user", http.StatusInternalServerError,
		)

		log.WithContext(r.Context()).WithError(err).Info(
			fmt.Sprintf("error updating user with id %d", id),
		)
		return
	}

	utils.RespondJSON(w, u, "success", "updated user", http.StatusOK)
	log.WithContext(r.Context()).Info(fmt.Sprintf("updated user with id %d", id))
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.RespondJSON(w, nil, "error", "invalid id", http.StatusBadRequest)
		log.WithContext(r.Context()).WithError(err).
			Info("failed to convert string to in")
		return
	}

	if !utils.HasAccess(r, id) {
		utils.RespondJSON(w, nil, "error",
			"unauthorized for requested content", http.StatusUnauthorized,
		)
		log.WithContext(r.Context()).Info(fmt.Sprintf("unauthorized for User %d", id))
		return
	}

	user := &models.User{ID: id}
	if err := h.UserHandler.UserInfo.Delete(user); err == nil {
		userAuth := &models.UserAuth{ID: id}
		if err = h.UserHandler.UserAuth.Delete(userAuth); err == nil {
			utils.RespondJSON(w, nil, "success",
				"deleted user", http.StatusOK,
			)
			log.WithContext(r.Context()).Info(
				fmt.Sprintf("Deleted user with id %d", id),
			)
			return
		}
		utils.RespondJSON(w, nil, "error", "error deleting user", http.StatusInternalServerError)
		log.WithContext(r.Context()).WithError(err).Info("error deleting user")

		return
	}

	utils.RespondJSON(w, nil, "error", "error deleting user", http.StatusInternalServerError)
	log.WithContext(r.Context()).Info("error deleting userauth")
}
