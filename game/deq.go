package game

type Deq struct {
  F []func([]float64) float64
  Init []float64
  T float64
  N int
}

func NewDeq(f []func([]float64) float64, init []float64, t float64, n int) *Deq {
  if len(f) != len(init) {
    panic("Invalid number of init vals!")
  }  
  return &Deq{
    F: f,
    Init: init,
    T: t,
    N: n,
  }
}

func (d *Deq) Solve() [][]float64 {
  l := len(d.F)
  x := [][]float64{}
  for j := 0 ; j < l; j++ {
    x = append(x, []float64{d.Init[j]})
  }
  h := d.T / float64(d.N)
  xx := make([]float64, len(d.F))
  copy(xx, d.Init)
  for i := 0 ; i < d.N; i++ {
    xxx := []float64{}
    for j := 0 ; j< l; j++ {
      next := xx[j] + h * d.F[j](xx)
      x[j] = append(x[j], next)
      xxx = append(xxx, next)
    }
    copy(xx, xxx)
  }
  return x
}

func ZeroAt(x [][]float64) int {
  l := len(x[0])
  c := len(x)
  for i := 0 ; i < l ; i++ {
    for j := 0 ; j < c; j ++ {
      if x[j][i] <= 0 {
        return i
      }
    }
  }
  return -1
}