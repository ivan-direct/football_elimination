package main

import (
	"football_elimination/games"
	"football_elimination/teams"
)

func main() {
	// bills := teams.New()
	// bills.ID = 1

	// db := teams.NewTeamService()

	// teams.Create(db, &packers)

	// teams.First(db, bills)
	// bills.TeamGreeting()

	// build team table
	// db.Migrator().DropTable(&seasons.Season{})
	// teams.AutoMigrate(db)
	// teams.Build(db)
	//TODO refactor teams.NewTeamService() to shared model since it is not just team specific
	gs := games.GameService{DB: teams.NewTeamService()}
	gs.DB.Migrator().DropTable(&games.Game{})
	gs.AutoMigrate()
	gs.BuildSeason()
	// teams.GroupByDivisional(db, "NFC")

}
