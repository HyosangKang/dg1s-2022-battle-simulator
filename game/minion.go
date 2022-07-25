package game

type Minion struct {
  Hp float64
  Dps float64
  Spd float64
}

func NewMinion(min Minion) *Minion {
  return &Minion{
    Hp: min.Hp,
    Dps: min.Dps,
    Spd: min.Spd,
  }
}