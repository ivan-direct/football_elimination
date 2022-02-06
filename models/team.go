package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "nathandipiazza"
	dbname = "football_elimination_dev"
)

// wins, loses, ties, division, and conference
type Team struct {
	// gorm.Model
	Id         uint
	Name       string
	Wins       int
	Loses      int
	Ties       int
	Division   string
	Conference string
}

func newTeamService() *gorm.DB {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}

func first(db *gorm.DB, team *Team) error {
	err := db.First(team).Error
	switch err {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return err
	default:
		return err
	}
}
