package repository

import (
	"database/sql"
	"fmt"

	"github.com/coddemn/get-access/internal/domain"
)

type ServiceRepository struct {
	db *sql.DB
}

func NewServiceRepo(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (r *ServiceRepository) GetById(id int, userId int) (domain.Service, error) {
	var service domain.Service

	err := r.db.QueryRow("SELECT * FROM services WHERE id = $1 and user_id = $2", id, userId).Scan(&service)
	if err != nil {
		if err == sql.ErrNoRows {
			return service, fmt.Errorf("GetServiceByID %d: unknown service", id)
		}
		return service, fmt.Errorf("GetServiceByID %d: %v", id, err)
	}

	return service, nil
}

func (r *ServiceRepository) GetByName(name string, userId int) (domain.Service, error) {
	var service domain.Service

	err := r.db.QueryRow("SELECT * FROM services WHERE name = $1 and user_id = $2", name, userId).Scan(&service.ID, &service.UserID, &service.Name, &service.Description, &service.Login, &service.Pass, &service.AddedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return service, fmt.Errorf("GetServiceByName %s: unknown service", name)
		}
		return service, fmt.Errorf("GetServiceByName %s: %v", name, err)
	}

	return service, nil
}

func (r *ServiceRepository) GetAllByUser(userId int) ([]domain.Service, error) {

	rows, err := r.db.Query("SELECT * FROM services WHERE user_id = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("GetAllServices: %v", err)
	}
	defer rows.Close()

	var services []domain.Service

	for rows.Next() {
		var service domain.Service

		err := rows.Scan(&service.ID, &service.UserID, &service.Name, &service.Description, &service.Login, &service.Pass, &service.AddedAt)
		if err != nil {
			return services, fmt.Errorf("GetAllServices: %v", err)
		}
		services = append(services, service)
	}

	if err = rows.Err(); err != nil {
		return services, fmt.Errorf("GetAllServices: %v", err)
	}

	return services, nil
}

func (r *ServiceRepository) Add(serviceData domain.Service) (int64, error) {

	row := r.db.QueryRow("INSERT INTO services (user_id, name, description, login, pass) VALUES ($1, $2, $3, $4, $5) RETURNING id", serviceData.UserID, serviceData.Name, serviceData.Description, serviceData.Login, serviceData.Pass)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("AddService: %v", err)
	}

	return id, nil
}
