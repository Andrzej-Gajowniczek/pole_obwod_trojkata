package main

import (
	"fmt"
	"math"
)

type trójkąt struct {
	a float64
	b float64
	c float64
	o float64
	p float64
}

func (t trójkąt) policz() (float64, float64) {
	t.o = t.a + t.b + t.c

	t.p = math.Sqrt(t.o / 2 * (t.o/2 - t.a) * (t.o/2 - t.b) * (t.o/2 - t.c))
	return t.p, t.o
}

func (t *trójkąt) pole2() (float64, float64) {
	t.o = t.a + t.b + t.c

	t.p = math.Sqrt(t.o / 2 * (t.o/2 - t.a) * (t.o/2 - t.b) * (t.o/2 - t.c))
	return t.p, t.o
}

type kwadrat struct {
	a float64
	o float64
	p float64
}

type prostokąt struct {
	a float64
	b float64
	o float64
	p float64
}
type koło struct {
	r float64
	p float64
	o float64
}

func main() {

	var T trójkąt
	T = trójkąt{
		a: 2,
		b: 2,
		c: 2,
	}

	pol, obw := T.policz()
	fmt.Println("obw:", obw, "pol:", pol)
	T.pole2()
	fmt.Println("Pole trójkąta o bokach:", T.a, T.b, T.c, "wynosi w jednostkach kwadratowych:", T.p, "a obwód w jednostkach liniowych:", T.o)

}
