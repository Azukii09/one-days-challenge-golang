package day_25

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}
type Puzzle struct {
	Nationalities [5]string
	Colors        [5]string
	Pets          [5]string
	Drinks        [5]string
	Smokes        [5]string
	Constraints   []func(i int) (bool, func(int))
}

func NewPuzzle() *Puzzle {
	new := Puzzle{}
	new.Drinks[2] = "Milk"             // const 9
	new.Nationalities[0] = "Norwegian" // const 10
	new.Colors[1] = "Blue"             // const 15
	new.Constraints = []func(i int) (bool, func(int)){
		new.ConstraintTwoApply, new.ConstraintThreeApply, new.ConstraintFourApply,
		new.ConstraintFiveApply, new.ConstraintSevenApply, new.ConstraintEightApply,
		new.ConstraintThirteenApply, new.ConstraintFouteenApply,
		// new.onstraintElevenApply, new.ConstraintTwelveApply,
	}
	return &new
}
func (p *Puzzle) ConstraintTwoApply(i int) (bool, func(int)) {
	if p.Nationalities[i] != "" || p.Colors[i] != "" {
		return false, nil
	}
	p.Nationalities[i] = "Englishman"
	p.Colors[i] = "Red"
	return true, p.ConstraintTwoRollback
}
func (p *Puzzle) ConstraintTwoRollback(i int) {
	p.Nationalities[i] = ""
	p.Colors[i] = ""
}
func (p *Puzzle) ConstraintThreeApply(i int) (bool, func(int)) {
	if p.Nationalities[i] != "" || p.Pets[i] != "" {
		return false, nil
	}
	p.Nationalities[i] = "Spaniard"
	p.Pets[i] = "Dog"
	return true, p.ConstraintThreeRollback
}
func (p *Puzzle) ConstraintThreeRollback(i int) {
	p.Nationalities[i] = ""
	p.Pets[i] = ""
}
func (p *Puzzle) ConstraintFourApply(i int) (bool, func(int)) {
	if i == 0 || p.Colors[i] != "" || p.Drinks[i] != "" || p.Colors[i-1] != "" {
		return false, nil
	}
	p.Colors[i] = "Green"
	p.Colors[i-1] = "Ivory"
	p.Drinks[i] = "Coffee"
	return true, p.ConstraintFourRollback
}
func (p *Puzzle) ConstraintFourRollback(i int) {
	p.Colors[i] = ""
	p.Colors[i-1] = ""
	p.Drinks[i] = ""
}
func (p *Puzzle) ConstraintFiveApply(i int) (bool, func(int)) {
	if p.Nationalities[i] != "" || p.Drinks[i] != "" {
		return false, nil
	}
	p.Nationalities[i] = "Ukrainian"
	p.Drinks[i] = "Tea"
	return true, p.ConstraintFiveRollback
}
func (p *Puzzle) ConstraintFiveRollback(i int) {
	p.Nationalities[i] = ""
	p.Drinks[i] = ""
}
func (p *Puzzle) ConstraintSevenApply(i int) (bool, func(int)) {
	if p.Pets[i] != "" || p.Smokes[i] != "" {
		return false, nil
	}
	p.Pets[i] = "Snails"
	p.Smokes[i] = "Old Gold"
	return true, p.ConstraintSevenRollback
}
func (p *Puzzle) ConstraintSevenRollback(i int) {
	p.Pets[i] = ""
	p.Smokes[i] = ""
}
func (p *Puzzle) ConstraintEightApply(i int) (bool, func(int)) {
	if p.Colors[i] != "" || p.Smokes[i] != "" {
		return false, nil
	}
	p.Colors[i] = "Yellow"
	p.Smokes[i] = "Kools"
	return true, p.ConstraintEightRollback
}
func (p *Puzzle) ConstraintEightRollback(i int) {
	p.Colors[i] = ""
	p.Smokes[i] = ""
}
func (p *Puzzle) ConstraintThirteenApply(i int) (bool, func(int)) {
	if p.Smokes[i] != "" || p.Drinks[i] != "" {
		return false, nil
	}
	p.Smokes[i] = "Lucky Strike"
	p.Drinks[i] = "Orange juice"
	return true, p.ConstraintThirteenRollback
}
func (p *Puzzle) ConstraintThirteenRollback(i int) {
	p.Smokes[i] = ""
	p.Drinks[i] = ""
}
func (p *Puzzle) ConstraintFouteenApply(i int) (bool, func(int)) {
	if p.Nationalities[i] != "" || p.Smokes[i] != "" {
		return false, nil
	}
	p.Nationalities[i] = "Japanese"
	p.Smokes[i] = "Parliaments"
	return true, p.ConstraintFouteenRollback
}
func (p *Puzzle) ConstraintFouteenRollback(i int) {
	p.Nationalities[i] = ""
	p.Smokes[i] = ""
}
func (p *Puzzle) GetGiraffeOwner() (bool, string) {
	var KoolsSmoker int
	var ChesterSmoker int
	for i := 0; i < 5; i++ {
		if p.Smokes[i] == "Kools" {
			KoolsSmoker = i
		} else if p.Smokes[i] == "" {
			ChesterSmoker = i
		}
	}
	var KoolsSmokerGoodNeighs []int
	var ChesterSmokerGoodNeighs []int
	UniqueGoodNeighs := make(map[int]bool)
	if KoolsSmoker > 0 && p.Pets[KoolsSmoker-1] == "" {
		KoolsSmokerGoodNeighs = append(KoolsSmokerGoodNeighs, KoolsSmoker-1)
		UniqueGoodNeighs[KoolsSmoker-1] = true
	} else if KoolsSmoker < 4 && p.Pets[KoolsSmoker+1] == "" {
		KoolsSmokerGoodNeighs = append(KoolsSmokerGoodNeighs, KoolsSmoker+1)
		UniqueGoodNeighs[KoolsSmoker+1] = true
	}
	if ChesterSmoker > 0 && p.Pets[ChesterSmoker-1] == "" {
		ChesterSmokerGoodNeighs = append(ChesterSmokerGoodNeighs, ChesterSmoker-1)
		UniqueGoodNeighs[ChesterSmoker-1] = true
	} else if ChesterSmoker < 4 && p.Pets[ChesterSmoker+1] == "" {
		ChesterSmokerGoodNeighs = append(ChesterSmokerGoodNeighs, ChesterSmoker+1)
		UniqueGoodNeighs[ChesterSmoker+1] = true
	}
	if len(UniqueGoodNeighs) == 2 && len(KoolsSmokerGoodNeighs) == 1 {
		for i := 0; i < 5; i++ {
			if i == KoolsSmokerGoodNeighs[0] || i == ChesterSmokerGoodNeighs[0] || p.Pets[i] != "" {
				continue
			}
			return true, p.Nationalities[i]
		}
	} else if len(UniqueGoodNeighs) == 3 {
		if len(KoolsSmokerGoodNeighs) == 2 {
			return true, p.Nationalities[KoolsSmokerGoodNeighs[0]]
		}
		return true, p.Nationalities[ChesterSmokerGoodNeighs[0]]
	}
	return false, ""
}
func (p *Puzzle) GetWaterDrinker() string {
	for i := 0; i < 5; i++ {
		if p.Drinks[i] == "" {
			return p.Nationalities[i]
		}
	}
	return ""
}

var usedConstraints = make(map[int]bool)
var solution Solution

func SolvePuzzle() Solution {
	p := NewPuzzle()
	dfs(p, 0)
	return solution
}
func dfs(p *Puzzle, level int) {
	if level == 8 {
		if success, giraffeOwner := p.GetGiraffeOwner(); success {
			waterDrinker := p.GetWaterDrinker()
			solution = Solution{waterDrinker, giraffeOwner}
		}
		return
	}
	for i := 0; i < 5; i++ {
		for ci, c := range p.Constraints {
			if solution != (Solution{}) {
				return
			}
			if usedConstraints[ci] {
				continue
			}
			if success, rollback := c(i); success {
				usedConstraints[ci] = true
				dfs(p, level+1)
				rollback(i)
				delete(usedConstraints, ci)
			}
		}
	}
}
