/*
Algorithms are tested against golang's native library for elliptic
curves in addition to our specialized libraries for hyperelliptic
curves.
*/

package attacks

import (
	"crypto/elliptic"
	//"crypto/rand"
	"fmt"
	"math/big"
	"testing"
)

func TestPollardRhoOnElliptic(t *testing.T) {
	curve := new(elliptic.CurveParams)
	curve.B = big.NewInt(89)
	curve.N = big.NewInt(58)
	curve.P = big.NewInt(173)
	curve.Gx = big.NewInt(10)
	curve.Gy = big.NewInt(59)
	curve.BitSize = 8

	priv := []byte{3}

	pX, pY := curve.ScalarBaseMult(priv)
	fmt.Println("Public key - X:", pX)
	fmt.Println("Public key - Y:", pY)
	fmt.Println("Private key:", priv)

	iX, iY := curve.Gx, curve.Gy
	for i := 2; i < 1000; i++ {
		iX, iY = curve.Add(iX, iY, curve.Gx, curve.Gy)
		if iX.Cmp(curve.Gx) == 0 && iY.Cmp(curve.Gy) == 0 {
			fmt.Println("Order:", i)
			break
		}
	}

	result := PollardRho(curve, pX, pY)
	fmt.Println("m:", result)
}

func TestBabyStepGiantStepOnElliptic(t *testing.T) {

}

/*
Benchmark comparisons against times gathered from implementations
in other languages applied to the same elliptic curves
*/
func BenchmarkPollardRhoOnElliptic(b *testing.B) {

}

func BenchmarkBabyStepGiantStepOnElliptic(b *testing.B) {

}
