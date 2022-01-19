package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

func Tally(reader io.Reader, writer io.Writer) error {
	matchesPlayed := map[string]int{}
	wins := map[string]int{}
	draws := map[string]int{}
	losses := map[string]int{}
	points := map[string]int{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 || scanner.Text()[0] == '#' {
			continue
		}
		result := strings.Split(scanner.Text(), ";")
		if len(result) != 3 {
			return errors.New("input is not in correct format")
		} else if result[2] != "win" && result[2] != "draw" && result[2] != "loss" {
			return errors.New("input is not in correct format")
		}
		matchesPlayed[result[0]]++
		matchesPlayed[result[1]]++
		if result[2] == "win" {
			wins[result[0]]++
			points[result[0]] += 3
			losses[result[1]]++
		} else if result[2] == "draw" {
			draws[result[0]]++
			points[result[0]]++
			draws[result[1]]++
			points[result[1]]++
		} else if result[2] == "loss" {
			losses[result[0]]++
			wins[result[1]]++
			points[result[1]] += 3
		}
	}
	teamList := []string{}
	for team := range matchesPlayed {
		teamList = append(teamList, team)
	}
	sort.Slice(teamList, func(i, j int) bool { return points[teamList[i]] > points[teamList[j]] || (points[teamList[i]] >= points[teamList[j]] && teamList[i] < teamList[j])})
    resultWriter := bufio.NewWriter(writer)
	header := fmt.Sprintf("%-30v | MP |  W |  D |  L |  P\n", "Team")
    resultWriter.WriteString(header)
    for _, team := range teamList {
        teamResults := fmt.Sprintf("%-30v | %2v | %2v | %2v | %2v | %2v\n", team, matchesPlayed[team], wins[team], draws[team], losses[team], points[team])
        resultWriter.WriteString(teamResults)
    }
    resultWriter.Flush()
	return nil
}
