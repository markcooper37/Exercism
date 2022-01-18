package tournament

import (
    "io"
    "bufio"
    "errors"
    "strings"
    "fmt"
)

func Tally(reader io.Reader, writer io.Writer) error {
    matchesPlayed := map[string]int{}
    wins := map[string]int{}
    draws := map[string]int{}
    losses := map[string]int{}
    points := map[string] int{}
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
	for key, value := range matchesPlayed {
        fmt.Println(key, value)
    }
    return nil
}
