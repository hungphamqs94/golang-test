package repository

import (
	models "go-postgres/model"
)

type HubRepo interface {
	Select() ([]models.Hub, error)
	Insert(u models.Hub) (error)
	SelectJoinHub() ([]models.JoinHub, error)
}