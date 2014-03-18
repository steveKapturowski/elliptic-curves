/*
Diffie-Hellman public-key cryptography implemented via the abelian group
structure on Jacobians of hyperelliptic (and elliptic) curves
Such a curve has the general form: C = y^2 - h(x)y - f(x) = 0
*/

package hyperelliptic

import (
	"math/big"
	"crypto/rand"
	//"fmt"
)

type HyperellipticCurve struct {
	H Polynomial // deg(H) < g
	F Polynomial // deg(F) = 2g+1
}

type Jacobian struct {
}

type Polynomial struct {
	Coefficients []big.Int
}

func (p *Polynomial) GetDegree() int {
	return len(p.Coefficients)
}

type Point struct {
	X *big.Int
	Y *big.Int
}

func (curve *HyperellipticCurve) IsNonsingular() bool {
	return true // sure, why not?
}

func (curve *HyperellipticCurve) IsElliptic() bool {
	return curve.GetGenus() == 1
}

func (curve *HyperellipticCurve) GetGenus() int {
	return (curve.F.GetDegree()-1)/2
}

func (curve *HyperellipticCurve) GetCharacteristic() int {
	return 0 // Want to allow curves defined over any finite field
}

func (curve *HyperellipticCurve) GetJacobian() *Jacobian {
	return new(Jacobian) // TODO: make this do something real
}

func (curve *HyperellipticCurve) Double(P *Point) *Point {
	x, y := P.X, P.Y
	return &Point{x, y} // TODO: make this do something real
}

func (curve *HyperellipticCurve) Add(P, Q *Point) *Point {
	x, y := P.X, P.Y
	return &Point{x, y} // TODO: make this do something real
}

//Returns m*P in O(log(m)) time and O(1) space
func (curve *HyperellipticCurve) DoubleAndAdd(P *Point, m *big.Int) *Point{
	kQ := *P
	curve.Double(&kQ)
	// TODO: use P.Bits() to boost performance 
	for i := m.BitLen()-2; i > 0; i-- {
		if m.Bit(i) == 1 {
			curve.Add(&kQ, P)
		}
		curve.Double(&kQ)
	}
	if m.Bit(0) == 1 {
		curve.Add(&kQ, P)
	}
	return &kQ
}

func GenerateBValue() *big.Int {
	max := new(big.Int)
	max.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)

	// TODO: make this satisfy an expanded set of properties beyond just primality
	b, _ := rand.Int(rand.Reader, max)
	if b.ProbablyPrime(1000) {
		return b
	}else {
		return GenerateBValue()
	}
}




