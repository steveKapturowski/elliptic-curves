package main

import (
	"math/big"
	"crypto/rand"
	"crypto/elliptic"
	"fmt"
	//"time"
)

func main(){
	b := GenerateBValue()
	fmt.Println(b)

	curve := elliptic.P521()

	priv, x, y, _ := elliptic.GenerateKey(curve, rand.Reader)
	fmt.Println(priv)
	fmt.Println(x)
	fmt.Println(y)

	p, _ := rand.Prime(rand.Reader, 512)
	fmt.Println(p)

}

type HyperellipticCurve struct {
	P Polynomial 
}

type Jacobian struct {
}

type Polynomial struct {
	Coefficients []big.Int
}

type Point struct {
	X big.Int
	Y big.Int
}

func IsNonsingular(curve *HyperellipticCurve) bool {
	return true // sure, why not?
}

func GetJacobian(curve *HyperellipticCurve) *Jacobian {
	return new(Jacobian) // TODO: make this do something real
}

func DoubleAndAdd() {
}

func GenerateBValue() *big.Int{
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




