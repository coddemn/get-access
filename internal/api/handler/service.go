package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/coddemn/get-access/internal/api/dto"
	"github.com/coddemn/get-access/internal/domain"
	"github.com/gin-gonic/gin"
)

type ServiceRepo interface {
	GetById(id int, userId int) (domain.Service, error)
	GetByName(name string, userId int) (domain.Service, error)
	GetAllByUser(userId int) ([]domain.Service, error)
	Add(serviceData domain.Service) (int64, error)
}

type ServiceHandler struct {
	repo ServiceRepo
}

func NewServiceHand(repo ServiceRepo) *ServiceHandler {
	return &ServiceHandler{
		repo: repo,
	}
}

func (h *ServiceHandler) OneService(c *gin.Context) {
	var req dto.ServiceReq

	req.ServName = c.Param("name")

	// TODO: get user_id from jwt-token or session
	var err error
	req.UserId, err = strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "no valid data"})
		return
	}

	service, err := h.repo.GetByName(req.ServName, req.UserId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "unknown service"})
		return
	}

	c.IndentedJSON(http.StatusOK, service)
}

func (h *ServiceHandler) AllServices(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "no valid data"})
		return
	}

	services, err := h.repo.GetAllByUser(userId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "unknown user or list is empty"})
		return
	}

	c.IndentedJSON(http.StatusOK, services)

}

func (h *ServiceHandler) NewService(c *gin.Context) {
	var newService dto.ServiceAdd

	if err := c.BindJSON(&newService); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "no valid data"})
		return
	}

	serviceData := domain.Service{
		Name:        newService.Name,
		Description: newService.Description,
		Login:       newService.Login,
		Pass:        newService.Pass,
		UserID:      newService.UserID,
	}

	serviceId, err := h.repo.Add(serviceData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"internal error": "service don`t added"})
		log.Printf("Service adding error: %v\n", err.Error())
		return
	}

	c.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"id":      serviceId,
			"message": "new service is added",
		},
	)
}
