package game

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type Sim struct {
	A, B *Team
	Init []float64
}

func NewSim(a, b *Team, init []float64) *Sim {
	return &Sim{
		A:    a,
		B:    b,
		Init: init,
	}
}

func (s *Sim) Run() {
	n := 3
	p := plot.New()
	var x0 float64 = 0
	for i := 0; i < n; i++ {
		var t float64 = 1
		f := []func([]float64) float64{
			func(x []float64) float64 {
				return -s.B.Member[0].Dps / s.A.Member[0].Hp * x[1]
			},
			func(x []float64) float64 {
				return -s.A.Member[0].Dps / s.B.Member[0].Hp * x[0]
			}}
		var sol [][]float64
		var idx int
		var d *Deq
		for {
			d = NewDeq(f, s.Init, t, 100)
			sol = d.Solve()
			idx = ZeroAt(sol)
			if idx >= 0 {
				break
			}
			t += 1
		}

		h := d.T / float64(d.N)
		for i := 0; i < 2; i++ {
			var pts plotter.XYs
			for j := 0; j < idx; j++ {
				pts = append(pts, plotter.XY{
					X: x0 + float64(j) * h,
					Y: sol[i][j]})
				line, _ := plotter.NewLine(pts)
				p.Add(line)
			}
		}
		for i := 0; i < 2; i++ {
			s.Init[i] = sol[i][idx-1] + 10
		}
		x0 += h * float64(idx)
	}
	p.Save(400, 400, "graph.png")
}

func (s *Sim) Print() {

}
