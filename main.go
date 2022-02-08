package main

import (
	"football_elimination/teams"
)

func main() {
	// packers := teams.Team{
	// 	Name:       "Packers",
	// 	Division:   "North",
	// 	Conference: "NFC",
	// }
	bills := teams.New()
	bills.ID = 1

	db := teams.NewTeamService()

	// teams.Create(db, &packers)

	teams.First(db, bills)
	bills.TeamGreeting()

	// build team table
	// db.Migrator().DropTable(&seasons.Season{})
	// teams.AutoMigrate(db)
	// teams.Build(db)

	// games.AutoMigrate(db)
	// games.DivisionalGameScheduler(db, "NFC", "North")

}
