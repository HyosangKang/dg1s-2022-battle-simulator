package game

type Team struct {
  Member []*Minion
}

func NewTeam(count int, min Minion) *Team {
  member := []*Minion{}
  for i := 0 ; i < count;i++ {
    minion := NewMinion(min)
    member = append(member, minion)
  }
  return &Team{
    Member: member,
  }
}