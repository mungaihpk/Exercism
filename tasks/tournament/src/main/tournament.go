package main

import (
	"fmt"
	"io"
	"strings"
)

type Team struct {
	name   string
	played int
	won    int
	drawn  int
	lost   int
	points int
}

func AddTeam(teams []Team, team Team, opponent Team) {

}

func (t Team) ToString() string {
	return fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", t.name, t.played, t.won, t.drawn, t.lost, t.points)
}

func (t *Team) UpdateScore(result string) {

	switch result {
	case "win":
		t.won = t.won + 1
		t.points += 3
		t.played += 1

	case "draw":
		t.drawn += 1
		t.points += 1
		t.played += 1

	case "loss":
		t.lost += 1
		t.played += 1

	}
}

func SortTeams(teams []Team) []Team {
	sorted_list := []Team{}

	for len(teams) > 0 {
		if len(teams) == 1 {
			sorted_list = append(sorted_list, teams[0])
			return sorted_list
		}

		var name string
		var points int

		team := teams[0]
		name = team.name
		points = team.points

		index := 0

		for i := 1; i < len(teams); i++ {
			new_team := teams[i]
			new_points := new_team.points
			new_name := new_team.name

			if new_points > points {
				index = i
			} else if new_points == points {
				if strings.Compare(new_name, name) == 1 {
					index = i
				}
			}
		}

		sorted_list = append(sorted_list, teams[index])
		teams = append(teams[:index], teams[index+1:]...)
	}
	return sorted_list
}

// Function to tally the totals
func Tally(reader io.Reader, writer io.Writer) error {
	//panic("Please implement the Tally function")
	var teams []Team

	bytes, err := io.ReadAll(reader)

	if err != nil {
		return err
	}

	slice := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	for _, line := range slice {
		entry := strings.Split(line, ";")

		name := strings.TrimSpace(entry[0])

		if strings.HasPrefix(name, "#") {
			continue
		} else if name == "" {
			continue
		} else if len(entry) > 0 && len(entry) < 3 {
			return fmt.Errorf("Invalid input values")
		}

		opponent := strings.TrimSpace(entry[1])
		result := strings.TrimSpace(entry[2])

		if result != "win" && result != "draw" && result != "loss" {
			return fmt.Errorf("Invalid input")
		}

		team_found := false
		opponent_found := false

		for i := range teams {
			team := &teams[i]
			if team.name == name {
				team.UpdateScore(result)
				team_found = true
				break
			}
		}

		if !team_found {
			team := Team{name: name}
			team_pointer := &team
			team_pointer.UpdateScore(result)
			teams = append(teams, team)
		}

		for i := range teams {
			team := &teams[i]
			if team.name == opponent {
				if result == "win" {
					team.UpdateScore("loss")
				} else if result == "loss" {
					team.UpdateScore("win")
				} else {
					team.UpdateScore("draw")
				}

				opponent_found = true
				break
			}
		}

		if !opponent_found {
			team := Team{name: opponent}
			team_pointer := &team

			if result == "win" {
				team_pointer.UpdateScore("loss")
			} else if result == "loss" {
				team_pointer.UpdateScore("win")
			} else {
				team_pointer.UpdateScore("draw")
			}

			teams = append(teams, team)
		}
	}

	_, err = io.WriteString(writer, fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P"))
	teams = SortTeams(teams)
	for _, team := range teams {
		_, err = io.WriteString(writer, team.ToString())
	}
	return err
}
