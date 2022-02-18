package games

import (
	"fmt"
	"football_elimination/teams"
	"math/rand"
	"time"

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
	// TODO toggle home/away games to keep them even
	nonConferenceGameScheduler(db)
	divisions := []string{"East", "North", "South", "West"}
	conferences := []string{"AFC", "NFC"}
	for _, conference := range conferences {
		// TODO toggle home/away games to keep them even
		divisionalGameScheduler(db, conference)
		for _, division := range divisions {
			interDivisionalGameScheduler(db, conference, division)
		}
	}
}

// Create Games records for Divisional Games
// Rules: Six games against divisional opponents
// two games per team, one at home and one on the road.
func interDivisionalGameScheduler(db *gorm.DB, conference, division string) {
	//TODO use cur year as default and implement ability to set custom year
	var year uint = 2022
	// TODO implement psuedorandom week generator
	divisionWeeks := []uint{1, 8, 9, 15, 17, 18}
	currentSchedule := make(map[string][]uint)

	divisionTeams := teams.FindDivisional(db, conference, division)
	for _, team := range divisionTeams {
		// create team bucket for storing scheduled weeks
		if _, ok := currentSchedule[team.Name]; !ok {
			currentSchedule[team.Name] = make([]uint, 6)
		}
		for _, opponent := range divisionTeams {
			// don't schedule a game against the same team
			if opponent.Name != team.Name {
				// create opponent bucket for storing scheduled weeks
				if _, ok := currentSchedule[opponent.Name]; !ok {
					currentSchedule[opponent.Name] = make([]uint, 6)
				}
				// find next available week for this game
				for i := 0; i < len(divisionWeeks); i++ {
					//Check to see if week is available
					week := divisionWeeks[i]
					teamSchedule := currentSchedule[team.Name]
					oppenentSchedule := currentSchedule[opponent.Name]
					// if not included in map, add to map and Create game
					if weekAvailable(week, teamSchedule, oppenentSchedule) {
						fmt.Println()
						// schedule a home and away game
						game := Game{Week: week, HomeTeam: team.Name, AwayTeam: opponent.Name, Year: year}
						Create(db, &game)
						// add to map to prevent scheduling conflicts
						currentSchedule[team.Name] = append(teamSchedule, week)
						currentSchedule[opponent.Name] = append(oppenentSchedule, week)
						break // found a week that will work
					}
				}
			}
		}
	}
}

// TODO need logic to determine which divisons play each other as it rotates.
// - Set 2022 accurately and use this as the origin to extrapolate for past/future years.
// Rules: Four games against teams from a division within its conference
// - two games at home and two on the road.
func divisionalGameScheduler(db *gorm.DB, conference string) {
	divisionTeams := make(map[string][]teams.Team)
	conferenceWeeks := []uint{2, 6, 7, 13}

	conferenceTeams := teams.GroupByDivisional(db, conference)
	divisionTeams["East"] = conferenceTeams[0:4]
	divisionTeams["North"] = conferenceTeams[4:8]
	scheduleByDivision(db, divisionTeams["East"], divisionTeams["North"], conferenceWeeks)
	// GroupByDivisional() needs to be called again
	conferenceTeams = teams.GroupByDivisional(db, conference)
	divisionTeams["South"] = conferenceTeams[8:12]
	divisionTeams["West"] = conferenceTeams[12:16]
	scheduleByDivision(db, divisionTeams["South"], divisionTeams["West"], conferenceWeeks)
}

// Stub Code
func shuffleDivision(teams []teams.Team) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(teams), func(i, j int) { teams[i], teams[j] = teams[j], teams[i] })
}

// Rules: Four games against teams from a division in the other conference
// — two games at home and two on the road.
func nonConferenceGameScheduler(db *gorm.DB) {
	//TODO replace hardcoded AFC East vs. NFC North, AFC South vs. NFC West
	afcDivisionOne := teams.FindDivisional(db, "AFC", "East")
	nfcDivisionOne := teams.FindDivisional(db, "NFC", "North")
	afcDivisionTwo := teams.FindDivisional(db, "AFC", "South")
	nfcDivisionTwo := teams.FindDivisional(db, "NFC", "West")
	nonConferenceWeeks := []uint{3, 4, 5, 14}

	scheduleByDivision(db, afcDivisionOne, nfcDivisionOne, nonConferenceWeeks)
	scheduleByDivision(db, afcDivisionTwo, nfcDivisionTwo, nonConferenceWeeks)
}

// schedule games for 2 divisions given an array of week values
func scheduleByDivision(db *gorm.DB, divisionOne, divisionTwo []teams.Team, weeks []uint) {
	for _, teamOne := range divisionOne {
		for j, teamTwo := range divisionTwo {
			week := weeks[j]
			game := Game{Week: week, HomeTeam: teamOne.Name, AwayTeam: teamTwo.Name, Year: 2022}
			Create(db, &game)
		}
		// must rotate divisionTwo to avoid scheduling inner team on the same week > 1
		divisionTwo = append(divisionTwo[1:], divisionTwo[0])
	}
}

// check both teams schedule to see if the given week has
// been scheduled yet. returns true if week hasn't been yet.
func weekAvailable(week uint, teamWeeks, opponentWeeks []uint) bool {
	for _, teamWeek := range teamWeeks {
		if week == teamWeek {
			return false
		}
	}
	for _, teamWeek := range opponentWeeks {
		if week == teamWeek {
			return false
		}
	}
	return true
}

// TODO need logic to determine last year's rankings. Set 2021 (last year) accurately and devise a means to record this.
// Rules: Two games against teams from the two remaining divisions in its own conference
// - one game at home and one on the road.
// - Matchups are based on division ranking from the previous season.
// func rankedDivisionalGameScheduler(){}
// for division, teams := range divisionTeams {
// 	shuffleDivision(teams) //TODO replace me with last years rankings
// 	fmt.Printf("Division:%s # # # # # # # # # #\n", division)
// 	for _, team := range teams {
// 		fmt.Printf("Team: %s\n", team.Name)
// 	}
// }
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
