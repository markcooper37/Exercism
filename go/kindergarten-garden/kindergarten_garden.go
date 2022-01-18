package kindergarten

import (
    "strings"
    "errors"
    "regexp"
)

type Garden struct {
    childPlants map[string][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
    plants := map[byte]string{'V': "violets", 'R': "radishes", 'C': "clover",
                                'G': "grass"}
    newGarden := &Garden{childPlants: map[string][]string{}}
    re := regexp.MustCompile(`\n[A-Z]+\n[A-Z]+`)
    correctformat := re.MatchString(diagram)
    if !correctformat {
        return newGarden, errors.New("incorrect format")
    }
    newDiagram := strings.Split(diagram, "\n")
    if len(newDiagram[1]) != len(newDiagram[2]) {
        return newGarden, errors.New("mismatched rows")
    }
	if len(newDiagram[1]) != 2 * len(children) {
        return newGarden, errors.New("rows of incorrect length")
    }
    for i := 0; i < len(children); i++ {
        child := children[i]
        if _, ok := newGarden.childPlants[child]; ok {
            return newGarden, errors.New("duplicate name")
        }
    	index := 0
        for j := 0; j < len(children); j++ {
            if child > children[j] {
                index += 1
            }
        }
        plant1, ok1 := plants[newDiagram[1][2 * index]]
        plant2, ok2 := plants[newDiagram[1][2 * index + 1]]
        plant3, ok3 := plants[newDiagram[2][2 * index]]
        plant4, ok4 := plants[newDiagram[2][2 * index + 1]]
        if !ok1 || !ok2 || !ok3 || !ok4 {
            return newGarden, errors.New("plant does not exist")
        }
    	childPlants := []string{plant1, plant2, plant3, plant4}
        newGarden.childPlants[child] = childPlants
    }
	return newGarden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
    if _, ok := g.childPlants[child]; !ok {
        return nil, false
    }
	return g.childPlants[child], true
}
