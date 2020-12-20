package main

import "fmt"

type Terminator struct {
	On    bool
	Ammo  int
	Power int
}

func (t *Terminator) Shoot() bool {
	if !t.On {
		return false
	} else {
		if t.Ammo > 0 {
			t.Ammo--
			return true
		} else {
			return false
		}
	}
}

func (t *Terminator) RideBike() bool {
	if !t.On {
		return false
	} else {
		if t.Power > 0 {
			t.Power--
			return true
		} else {
			return false
		}
	}
}

func main() {
	term := new(Terminator)
	term.On = true
	term.Ammo = 10
	term.Power = 2
	for i := 0; i < 12; i++ {
		fmt.Printf("shoot %d: %t \n", i, term.Shoot())
		fmt.Printf("RideBike %d: %t \n", i, term.RideBike())
	}
}
