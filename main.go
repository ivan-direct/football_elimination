package main

import (
	"football_elimination/games"
	"football_elimination/teams"
)

func main() {
	// bills := teams.New()
	// bills.ID = 1

	db := teams.NewTeamService()

	// teams.Create(db, &packers)

	// teams.First(db, bills)
	// bills.TeamGreeting()

	// build team table
	// db.Migrator().DropTable(&seasons.Season{})
	// teams.AutoMigrate(db)
	// teams.Build(db)

	db.Migrator().DropTable(&games.Game{})
	games.AutoMigrate(db)
	games.BuildSeason(db)
	// teams.GroupByDivisional(db, "NFC")

}
