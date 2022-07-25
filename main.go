package main

import (
	"main/game"
)

func main() {
  am := game.Minion{
    Hp: 10,
    Dps: 1,
    Spd: 1,
  }
  bm := game.Minion{
    Hp: 5,
    Dps: 2,
    Spd: 1,
  }
  a := game.NewTeam(10, am)
  b := game.NewTeam(5, bm)
  init := []float64{5, 10}
	s := game.NewSim(a, b, init)
  s.Run()
  s.Print()
}
