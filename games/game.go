package games

import "gorm.io/gorm"

// Game model representing a scheduled game in an NFL season
type Game struct {
	gorm.Model
	Week          uint
	HomeTeam      string
	AwayTeam      string
	WinningTeamID uint
	Tie           bool
	Year          uint
}

// Creates games for all 32 teams for 17 weeks ////
// Honors current (as of 2022) NFL scheduling rules
//TODO func BuildSeason()

// Create a Game record in the Database
func Create(db *gorm.DB, game *Game) {
	db.Create(game)
}

// run GORM AutoMigrate using Game struct
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Game{})
	if err != nil {
		return err
	}
	return nil
}
