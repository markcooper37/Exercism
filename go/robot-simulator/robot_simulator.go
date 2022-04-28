package robot

import "fmt"

const (
	N Dir = iota
	E
	S
	W
)

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

func Advance() {
	if Step1Robot.Dir == N {
		Step1Robot.Y++
	} else if Step1Robot.Dir == E {
		Step1Robot.X++
	} else if Step1Robot.Dir == S {
		Step1Robot.Y--
	} else if Step1Robot.Dir == W {
		Step1Robot.X--
	}
}

func (d Dir) String() string {
	if d == N {
		return "N"
	} else if d == E {
		return "E"
	} else if d == S {
		return "S"
	} else {
		return "W"
	}
}

type Action byte

func StartRobot(command chan Command, action chan Action) {
	for toDo := range command {
		action <- Action(toDo)
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for toDo := range action {
		if toDo == 'R' {
			robot.Dir = (robot.Dir + 1) % 4
		} else if toDo == 'L' {
			robot.Dir = (robot.Dir + 3) % 4
		} else if toDo == 'A' {
			if robot.Dir == N && robot.Pos.Northing < extent.Max.Northing {
				robot.Pos.Northing++
			} else if robot.Dir == E && robot.Pos.Easting < extent.Max.Easting {
				robot.Pos.Easting++
			} else if robot.Dir == S && robot.Pos.Northing > extent.Min.Northing {
				robot.Pos.Northing--
			} else if robot.Dir == W && robot.Pos.Easting > extent.Min.Easting {
				robot.Pos.Easting--
			}
		}
	}
	report <- robot
	close(report)
}

type Action3 struct {
	name   string
	script string
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	for _, command := range script {
		if command != 'A' && command != 'R' && command != 'L' {
			log <- "invalid script"
			action <- Action3{name: name, script: ""}
			return
		}
	}
	action <- Action3{name: name, script: script}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	for index, robot := range robots {
		if robot.Name == "" {
			log <- "robot has no name"
		}
		if robot.Northing < extent.Min.Northing || robot.Northing > extent.Max.Northing || robot.Easting < extent.Min.Easting || robot.Easting > extent.Max.Easting {
			log <- "robot placed outside room"
		}
		for i := index + 1; i < len(robots); i++ {
			if robot.Name == robots[i].Name {
				log <- fmt.Sprintf("duplicate name: %s", robot.Name)
			}
			if robot.Pos == robots[i].Pos {
				log <- "robots placed in same position"
			}
		}
	}
	var toDo Action3
	for range robots {
		toDo = <-action
		isRobot := false
		for robotIndex, robot := range robots {
			if robot.Name == toDo.name {
				isRobot = true
				for _, command := range toDo.script {
					if command == 'R' {
						robot.Dir = (robot.Dir + 1) % 4
					} else if command == 'L' {
						robot.Dir = (robot.Dir + 3) % 4
					} else if command == 'A' {
						if IsBlockedByRobot(robot.Pos, robot.Dir, robots) {
							log <- "running into robot"
							continue
						}
						var moved bool
						robot, moved = HitWallOrMove(robot, extent)
						if !moved {
							log <- "running into wall"
						}
					}
					robots[robotIndex] = robot
				}
			}
		}
		if !isRobot {
			log <- fmt.Sprintf("no robot with name %s exists", toDo.name)
		}
	}
	rep <- robots
	close(rep)
}

func IsBlockedByRobot(position Pos, direction Dir, robots []Step3Robot) bool {
	newPos := position
		switch direction {
		case N:
			newPos.Northing++
		case E:
			newPos.Easting++
		case S:
			newPos.Northing--
		case W:
			newPos.Easting--
		}
	for _, robot := range robots {
		if robot.Pos == newPos {
			return true
		}
	}
	return false
}

func HitWallOrMove(robot Step3Robot, extent Rect) (Step3Robot, bool) {
	switch robot.Dir {
	case N:
		if robot.Pos.Northing >= extent.Max.Northing {
			return robot, false
		} else {
			robot.Pos.Northing++
		}
	case E:
		if robot.Pos.Easting >= extent.Max.Easting {
			return robot, false
		} else {
			robot.Pos.Easting++
		}
	case S:
		if robot.Pos.Northing <= extent.Min.Northing {
			return robot, false
		} else {
			robot.Pos.Northing--
		}
	case W:
		if robot.Pos.Easting <= extent.Min.Easting {
			return robot, false
		} else {
			robot.Pos.Easting--
		}
	}
	return robot, true
}
