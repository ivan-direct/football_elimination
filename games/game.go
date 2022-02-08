package games

import (
	"football_elimination/teams"

	"gorm.io/gorm"
)

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
func BuildSeason(db *gorm.DB) {
	divisions := []string{"East", "North", "South", "West"}
	conferences := []string{"AFC", "NFC"}
	for _, conference := range conferences {
		for _, division := range divisions {
			interDivisionalGameScheduler(db, conference, division)
		}
		// TODO Schedule all conference home games
	}
}

// Create Games records for Divisional Games
// Rules: Six games against divisional opponents
// two games per team, one at home and one on the road.
func interDivisionalGameScheduler(db *gorm.DB, conference, division string) {
	//TODO use cur year as default and implement ability to set custom year
	var year uint = 2022
	division_teams := teams.FindDivisional(db, conference, division)
	for i, team := range division_teams {
		// TODO implement psuedorandom week generator
		// hard coding week as 1 or 12 for now...
		var week uint = 1
		if i%2 == 0 {
			week = 12
		}
		for _, opponent := range division_teams {
			if opponent.Name != team.Name {
				// schedule a home and away game
				game := Game{Week: week, HomeTeam: team.Name, AwayTeam: opponent.Name, Year: year}
				Create(db, &game)
			}
		}
	}
}

// TODO need logic to determine which divisons play each other as it rotates. Set 2022 accurately and use this as the origin to extrapolate for past/future years.
// Rules: Four games against teams from a division within its conference
// - two games at home and two on the road.
// func divisionalGameScheduler(){}
// Rules: Four games against teams from a division in the other conference
// — two games at home and two on the road.
// func nonConferenceGameScheduler(){}

// TODO need logic to determine last year's rankings. Set 2021 (last year) accurately and devise a means to record this.
// Rules: Two games against teams from the two remaining divisions in its own conference
// - one game at home and one on the road.
// - Matchups are based on division ranking from the previous season.
// func rankedDivisionalGameScheduler(){}
// Rules One game against a non-conference opponent from a division that the team is not scheduled to play.
// Matchups are based on division ranking from the previous season.
// func rankedNonConferenceGameScheduler(){}

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
