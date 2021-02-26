package main

import (
	"fmt"
	"go-postgres/driver"
	"go-postgres/repository/repoimpl"
	models "go-postgres/model"
)

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	password = "123456"
	dbname =" test"
)

func main(){
	db  := driver.Connect(host, port, user, password, dbname)
	err := db.SQL.Ping()
    if err != nil {
		panic(err)
	}
	userRepo := repoimpl.NewUserRepo(db.SQL)
    hubRepo := repoimpl.NewHubRepo(db.SQL)
	teamRepo := repoimpl.NewTeamRepo(db.SQL)

	userOne := models.User{
		UserID: 1,
		Role: "DEV",
		Email: "hungphamit94@gmail.com",
		TeamType: "TEAMDEV",
	}

	userRepo.Insert(userOne)

	hubOne := models.Hub{
		ID: 1,
		Name: "Innovation",
		GeoLocation: "Quang Trung, Ho Chi Minh",
	}

	hubRepo.Insert(hubOne)

	teamOne := models.Team {
		ID: 1,
		Type: "DEV",
		Name: "TeamDEV",
		GeoLocation: "Quang Trung, Ho Chi Minh",
	}

	teamRepo.Insert(teamOne)
	
	users, _ := userRepo.Select()
	hubs, _ := hubRepo.Select()
	teams, _ := teamRepo.Select()
	for i:= range users {
		fmt.Println(users[i])
	}
	for i:= range hubs {
		fmt.Println(hubs[i])
	}
	for i:= range teams {
		fmt.Println(teams[i])
	}

	joinHubs, _ := hubRepo.SelectJoinHub()
	for i:= range joinHubs {
		fmt.Println(joinHubs[i])
	}
}