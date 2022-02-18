package teams

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

// An NFL team including record
type Team struct {
	gorm.Model
	Name       string
	Wins       int
	Loses      int
	Ties       int
	Division   string
	Conference string
}

// GO PACK GO!!!
func (t *Team) TeamGreeting() {
	fmt.Printf("Go %v!!!\n", t.Name)
}

// return a pointer to a new Team struct
func New() (t *Team) {
	return &Team{}
}

// initialize the Database
func NewTeamService() *gorm.DB {
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

// Create a Team record in the Database
func Create(db *gorm.DB, team *Team) {
	db.Create(team)
}

func FindDivisional(db *gorm.DB, conference, division string) []Team {
	teams := []Team{}
	db.Where("conference = ? and division = ?", conference, division).Find(&teams)
	return teams
}

func GroupByDivisional(db *gorm.DB, conference string) []Team {
	teams := []Team{}
	db.Model(&Team{}).Where("conference = ?", conference).Order("division, name").Find(&teams)
	return teams
}

// Select the first team from the Database
func First(db *gorm.DB, team *Team) error {
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

// run GORM AutoMigrate using Team struct
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Team{})
	if err != nil {
		return err
	}
	return nil
}
