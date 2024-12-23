package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(t1 string, s1 int, t2 string, s2 int) {
	if _, ok := l.Teams[t1]; !ok {
		// team 1 is not in league
		return
	}
	if _, ok := l.Teams[t2]; !ok {
		// team 2 is not in league
		return
	}
	if s1 == s2 {
		// tie; both teams get a point
		l.Wins[t1]++
		l.Wins[t2]++
	}
	if s1 > s2 {
		// team 1 gets a win
		l.Wins[t1]++
	}
	if s1 < s2 {
		// team 2 gets a win
		l.Wins[t2]++
	}
}

func (l League) Ranking() []string {
	teams := make([]string, 0, len(l.Teams))

	for t := range l.Teams {
		teams = append(teams, t)
	}

	sort.Slice(teams, func(i, j int) bool {
		return l.Wins[teams[i]] > l.Wins[teams[j]]
	})

	return teams
}

func NewTeam(name string) Team {
	return Team{
		Name:    name,
		Players: []string{"Player 1", "Player 2", "Player 3", "Player 4"},
	}
}

func NewLeague(name string, teamNames []string) League {
	teams := make(map[string]Team, len(teamNames))

	for _, teamName := range teamNames {
		teams[teamName] = NewTeam(teamName)
	}

	return League{
		Name:  name,
		Teams: teams,
		Wins:  map[string]int{},
	}
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	ranking := r.Ranking()

	for _, v := range ranking {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}

}

func main() {
	// establish league
	league := NewLeague("Bush League", []string{"Canada", "France", "Australia"})

	// play matches
	league.MatchResult("Canada", 3, "France", 2)
	league.MatchResult("Canada", 1, "Australia", 0)
	league.MatchResult("France", 0, "Australia", 4)

	// final results
	RankPrinter(league, os.Stdout)

}
