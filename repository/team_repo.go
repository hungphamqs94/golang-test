package repository

import (
	models "go-postgres/model"
)

type TeamRepo interface {
	Select() ([]models.Team, error)
	Insert(u models.Team) (error)
}