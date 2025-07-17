package api

import (
	"errors"
	"log"
	"regexp"

	"github.com/muhammad21236/femProject/internal/store"
)

type registerUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

type UserHandler struct {
	userStore store.UserStore
	logger    *log.Logger
}

func NewUserHandler(userStore store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		logger:    logger,
	}
}

func (h *UserHandler) validateRegisterRequest(req *registerUserRequest) error {
	if req.Username == "" {
		return errors.New("USERNAME IS REQUIRED FIELD")
	}
	if len(req.Username) > 50 {
		return errors.New("USERNAME CANNOT BE GREATER THAN 50 CHARACTERS")
	}
	if req.Email == "" {
		return errors.New("EMAIL IS REQUIRED FIELD")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return errors.New("INVALID EMAIL FORMAT")
	}
	if req.Password == "" {
		return errors.New("PASSWORD IS REQUIRED FIELD")
	}
	return nil
}
