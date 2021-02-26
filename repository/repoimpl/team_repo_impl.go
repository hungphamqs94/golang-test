package repoimpl

import (
	"fmt"
	"database/sql"
	repo "go-postgres/repository"
	models "go-postgres/model"
)

type TeamRepoImpl struct {
	Db *sql.DB
}

func NewTeamRepo(db *sql.DB) repo.TeamRepo {
	return &TeamRepoImpl {
		Db : db,
	}
}

func (u *TeamRepoImpl) Select() ([]models.Team, error) {
	teams := make([]models.Team, 0)

	rows, err := u.Db.Query("SELECT * FROM team")
	if err != nil {
		return teams, err
	}
	defer rows.Close()
	for rows.Next() {
		team := models.Team{}

		err := rows.Scan(&team.ID, &team.Type, &team.Name,&team.GeoLocation)
		if err != nil {
			break
		}

		teams = append(teams, team)
	}

	err = rows.Err()
	if err != nil {
		return teams, err
	}

	return teams, nil
}

func (u *TeamRepoImpl) Insert(team models.Team) (error) {
	insertStatement := `
	INSERT INTO team (ID, Type, Name, GeoLocation)
	VALUES ($1, $2, $3, $4)`

	_, err := u.Db.Exec(insertStatement, team.ID, team.Type, team.Name, team.GeoLocation)	
	if err != nil {
		return err
	}

	fmt.Println("Record added: ", team)

	return nil
}
