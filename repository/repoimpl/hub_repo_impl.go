package repoimpl

import (
	"fmt"
	"database/sql"
	repo "go-postgres/repository"
	models "go-postgres/model"
)

type HubRepoImpl struct {
	Db *sql.DB
}


func NewHubRepo(db *sql.DB) repo.HubRepo {
	return &HubRepoImpl {
		Db : db,
	}
}

func (u *HubRepoImpl) Select() ([]models.Hub, error) {
	hubs := make([]models.Hub, 0)

	rows, err := u.Db.Query("SELECT * FROM hub")
	if err != nil {
		return hubs, err
	}
	defer rows.Close()
	for rows.Next() {
		hub := models.Hub{}

		err := rows.Scan(&hub.ID, &hub.Name, &hub.GeoLocation)
		if err != nil {
			break
		}

		hubs = append(hubs, hub)
	}

	err = rows.Err()
	if err != nil {
		return hubs, err
	}

	return hubs, nil
}

func (u *HubRepoImpl) SelectJoinHub() ([]models.JoinHub, error) {
	joinHubs := make([]models.JoinHub, 0)

	rows, err := u.Db.Query("Select hub.id, hub.name, hub.geolocation, users.userid, users.role, users.teamtype, users.email, team.name FROM hub, users, team where hub.GeoLocation=team.GeoLocation and users.TeamType = team.Type")
	if err != nil {
		return joinHubs, err
	}
	defer rows.Close()
	for rows.Next() {
		joinHub := models.JoinHub{}
		err := rows.Scan(&joinHub.HubID, &joinHub.Name, &joinHub.GeoLocation, &joinHub.UserID, &joinHub.RoleUser, &joinHub.TeamType, &joinHub.EmailUser, &joinHub.TeamName)
		if err != nil {
			break
		}

		joinHubs = append(joinHubs, joinHub)
	}
	err = rows.Err()
	if err != nil {
		return joinHubs, err
	}

	return joinHubs, nil
}

func (u *HubRepoImpl) Insert(hub models.Hub) (error) {
	insertStatement := `
	INSERT INTO hub (ID, Name, GeoLocation)
	VALUES ($1, $2, $3)`

	_, err := u.Db.Exec(insertStatement, hub.ID, hub.Name, hub.GeoLocation)	
	if err != nil {
		return err
	}

	fmt.Println("Record added: ", hub)

	return nil
}
