package teams

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateTeam(db *gorm.DB, teams []string, conference string, division string) {
	team := Team{}
	for _, name := range teams {
		fmt.Println("Creating...", name)
		team = Team{
			Name:       name,
			Division:   division,
			Conference: conference,
			Wins:       0,
			Loses:      0,
			Ties:       0,
		}
		Create(db, &team)
	}
}

// Populate Database with all 32 NFL teams
// Last updated 2-2022
func Build(db *gorm.DB) {
	// AFC //////////////////////////////////////////////////
	conference := "AFC"
	division := "East"
	teams := []string{"Buffalo Bills",
		"Miami Dolphins",
		"New England Patriots",
		"New York Jets"}
	CreateTeam(db, teams, conference, division)

	division = "North"
	teams = []string{"Baltimore Ravens",
		"Cincinnati Bengals",
		"Cleveland Browns",
		"Pittsburgh Steelers"}
	CreateTeam(db, teams, conference, division)

	division = "South"
	teams = []string{"Houston Texans",
		"Indianapolis Colts",
		"Jacksonville Jaguars",
		"Tennessee Titans"}
	CreateTeam(db, teams, conference, division)

	division = "West"
	teams = []string{"Denver Broncos",
		"Kansas City Chiefs",
		"Las Vegas Raiders",
		"Los Angeles Chargers"}
	CreateTeam(db, teams, conference, division)
	// NFC //////////////////////////////////////////////////
	conference = "NFC"
	division = "East"
	teams = []string{"Dallas Cowboys",
		"New York Giants",
		"Philadelphia Eagles",
		"Washington Commanders"}
	CreateTeam(db, teams, conference, division)

	division = "North"
	teams = []string{"Chicago Bears",
		"Detroit Lions",
		"Green Bay Packers",
		"Minnesota Vikings"}
	CreateTeam(db, teams, conference, division)

	division = "South"
	teams = []string{"Atlanta Falcons",
		"Carolina Panthers",
		"New Orleans Saints",
		"Tampa Bay Buccaneers"}
	CreateTeam(db, teams, conference, division)

	division = "West"
	teams = []string{"Arizona Cardinals",
		"Los Angeles Rams",
		"San Francisco 49ers",
		"Seattle Seahawks"}
	CreateTeam(db, teams, conference, division)

}
