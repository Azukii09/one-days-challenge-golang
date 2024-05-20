package robot

type DirA int

const (
	N DirA = iota
	E
	S
	W
)

func right(dir DirA) DirA {
	return (dir + 1) % 4
}

func left(dir DirA) DirA {
	return (dir + 3) % 4
}

func Right() {
	Step1Robot.DirA = right(Step1Robot.DirA)
}

func Left() {
	Step1Robot.DirA = left(Step1Robot.DirA)
}

func advance(dir DirA, x int, y int) (int, int) {
	if dir == N {
		y += 1
	} else if dir == E {
		x += 1
	} else if dir == S {
		y -= 1
	} else {
		x -= 1
	}
	return x, y
}

func Advance() {
	Step1Robot.X, Step1Robot.Y = advance(Step1Robot.DirA, Step1Robot.X, Step1Robot.Y)
}

func (d DirA) String() string {
	panic("Please implement the String function")
}

type Action Command

func StartRobot(command chan Command, action chan Action) {
	for c := range command {
		action <- Action(c)
	}
	close(action)
}

func inRectangle(p Pos, r Rect) bool {
	return p.Easting >= r.Min.Easting && p.Easting <= r.Max.Easting && p.Northing >= r.Min.Northing && p.Northing <= r.Max.Northing
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for a := range action {
		switch a {
		case 'R':
			robot.DirA = right(robot.DirA)
		case 'L':
			robot.DirA = left(robot.DirA)
		case 'A':
			var x, y = advance(robot.DirA, int(robot.Pos.Easting), int(robot.Pos.Northing))
			if inRectangle(Pos{RU(x), RU(y)}, extent) {
				robot.Pos.Easting = RU(x)
				robot.Pos.Northing = RU(y)
			}
		}
	}
	report <- robot
}

type Action3 struct {
	name   string
	action byte
}

const done = 7

func StartRobot3(name, script string, action chan Action3, log chan string) {
	for _, c := range script {
		action <- Action3{name, byte(c)}
	}
	action <- Action3{name, done}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	defer func() { rep <- robots }()

	names := map[string]bool{}
	occupiedPositions := map[Pos]bool{}
	robotIndexByName := map[string]int{}
	doneCount := 0

	for i, robot := range robots {
		if robot.Name == "" {
			log <- "Missing name"
			return
		}

		if names[robot.Name] {
			log <- "Duplicate name"
			return
		}
		names[robot.Name] = true

		if occupiedPositions[robot.Pos] {
			log <- "Duplicate position"
			return
		}
		occupiedPositions[robot.Pos] = true

		if !inRectangle(robot.Pos, extent) {
			log <- "Invalid position"
			return
		}

		robotIndexByName[robot.Name] = i
	}

	for a := range action {
		i, ok := robotIndexByName[a.name]
		robot := &robots[i]
		if !ok {
			log <- "Unknown robot"
			return
		}

		switch a.action {
		case 'R':
			robot.DirA = right(robot.DirA)
		case 'L':
			robot.DirA = left(robot.DirA)
		case 'A':
			x, y := advance(robot.DirA, int(robot.Pos.Easting), int(robot.Pos.Northing))
			newPos := Pos{RU(x), RU(y)}
			if occupiedPositions[newPos] {
				log <- "Bumped a robot"
			} else if !inRectangle(newPos, extent) {
				log <- "Bumped a wall"
			} else {
				occupiedPositions[robot.Pos] = false
				robot.Pos.Easting = newPos.Easting
				robot.Pos.Northing = newPos.Northing
				occupiedPositions[robot.Pos] = true
			}
		case done:
			doneCount += 1
			if doneCount == len(robots) {
				return
			}
		default:
			log <- "Invalid action"
			return
		}
	}
}
