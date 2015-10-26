package main

import (
	"./sort"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"log"
	"math/rand"
	"testing"
)

var (
	SAMPLES  int    = 16
	FILENAME string = "perf.png"
	result   [][]int
)

// Bench sorting algorithms for SAMPLES powers of 2,
// saving the results to FILENAME.
func main() {
	// To add more algorithms to the graphical bencher,
	// add them to names and fns.
	//
	// No other part of this program need be modified.
	names := []string{"Insertion", "Merge"}
	fns := []func([]int){sort.Insertion, sort.Merge}

	ns := make([][]int64, len(fns))
	X := make([]int64, SAMPLES)

	N := 2
	for pow := 0; pow < len(X); pow++ {
		X[pow] = int64(N)
		N *= 2
	}

	for i, fn := range fns {
		log.Printf("Benching %s", names[i])
		ns[i] = make([]int64, len(X))
		for x, N := range X {
			result := testing.Benchmark(getBenchmarker(fn, int(N)))
			nsPerOp := result.NsPerOp()
			log.Printf("\t%d\t%d", N, nsPerOp)
			ns[i][x] = nsPerOp
		}
	}

	save(FILENAME, names, ns, X)
}

// Get N unsorted lists of length l for benching.
func getLists(N, l int) [][]int {
	lists := make([][]int, N)
	for i := 0; i < N; i++ {
		lists[i] = make([]int, l)
		for j := 0; j < l; j++ {
			lists[i][j] = rand.Int()
		}
	}
	return lists
}

// Get Benchmark function for algorithm sorting lists of length l.
func getBenchmarker(algorithm func([]int), l int) func(b *testing.B) {
	return func(b *testing.B) {
		lists := getLists(b.N, l)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			algorithm(lists[i])
		}

		// Store the result to avoid compiler optimization
		result = lists
	}
}

// Save benched results.
func save(filename string, names []string, ns [][]int64, X []int64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "ns/sort"
	p.X.Label.Text = "N"
	p.Y.Label.Text = "ns"
	p.X.Scale = plot.LogScale{}
	p.Y.Scale = plot.LogScale{}

	data := make([]interface{}, len(names)*2)
	for i, name := range names {
		data[i*2] = name
		data[i*2+1] = points(X, ns[i])
	}

	if err := plotutil.AddLinePoints(p, data...); err != nil {
		panic(err)
	}

	if err := p.Save(10*vg.Inch, 8*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func points(X []int64, Y []int64) plotter.XYs {
	pts := make(plotter.XYs, len(Y))

	for i, _ := range Y {
		pts[i].X = float64(X[i])
		pts[i].Y = float64(Y[i])
	}

	return pts
}
