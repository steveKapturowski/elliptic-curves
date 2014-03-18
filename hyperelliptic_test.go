package hyperelliptic

import(
	"testing"
	"math/big"
	"fmt"
)

func TestDoubleAndAdd(t *testing.T){
	x := big.NewInt(107)
	y := big.NewInt(56)
	m := new(big.Int)
	m.SetString("101100", 2)
	p0 := &Point{x, y}

	curve := new(HyperellipticCurve)

	curve.DoubleAndAdd(p0, m)
	fmt.Println(*p0)
}