/*
Algorithms are tested against golang's native library for elliptic
curves in addition to our specialized libraries for hyperelliptic
curves.
*/

package attacks

import (
	"testing"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func TestPollardRhoOnElliptic(t *testing.T) {
	curve := elliptic.P521()

	priv, x, y, _ := elliptic.GenerateKey(curve, rand.Reader)
	fmt.Println(priv)
	fmt.Println(x)
	fmt.Println(y)
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

