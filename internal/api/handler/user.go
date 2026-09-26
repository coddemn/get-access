package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/coddemn/get-access/internal/api/apperrors"
	"github.com/coddemn/get-access/internal/api/dto"
	"github.com/coddemn/get-access/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserRepo interface {
	GetByLogin(login, pass string) (domain.User, error)
	Add(userData domain.User) (int64, error)
}

type UserHandler struct {
	repo UserRepo
}

func NewUserHand(repo UserRepo) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (h *UserHandler) Auth(c *gin.Context) {
	var userData dto.LogIn

	if err := c.BindJSON(&userData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "no valid auth data"})
		return
	}

	user, err := h.repo.GetByLogin(userData.Login, userData.Pass)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrIncorrectAuth):
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": apperrors.ErrIncorrectAuth.Error()})
			return
		default:
			log.Printf("User auth error: %v\n", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"internal error": "user not authorized"})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, user)

}

func (h *UserHandler) Registrate(c *gin.Context) {

	var newUser dto.UserRegictrate

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "incorrect user data"})
		return
	}

	userData := domain.User{
		Name:  newUser.Name,
		Login: newUser.Login,
		Pass:  newUser.Password,
	}

	userId, err := h.repo.Add(userData)
	if err != nil {

		switch {
		case errors.Is(err, apperrors.ErrNameTaken):
			c.IndentedJSON(http.StatusConflict, gin.H{"error": apperrors.ErrNameTaken.Error()})
			return
		default:
			log.Printf("User creating error: %v\n", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"internal error": "user don`t registrated"})
			return
		}

	}

	c.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"id":      userId,
			"message": "new user is added",
		},
	)

}
