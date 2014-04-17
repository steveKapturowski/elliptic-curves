/*
This library is intended ONLY for the safety verification of cryptographic
schemes based on the discrete logarithm problem, and is not in any way
intended to aid attackers in breaking cryptographic schemes that are used
in practice. Furthermore, we take the the view that any scheme utilizing
a set of parameters that CAN be broken by the below algorithms has no
business being used for any practical security purposes.
*/

package attacks

import (
	hyperelliptic "../../HyperellipticFuntime"
	"crypto/elliptic"
	"math/big"
)

func ellipticMixingFunction(zX, zY, pX, pY, alpha, beta *big.Int,
	curve *elliptic.CurveParams) (*big.Int, *big.Int) {
	params := curve.Params()

	cutoff_1 := big.NewInt(1).Div(params.P, big.NewInt(3))
	cutoff_2 := big.NewInt(1).Div(params.P, big.NewInt(3))
	cutoff_2.Mul(cutoff_2, big.NewInt(2))

	switch {
	case zY.Cmp(cutoff_1) == -1: // y < cutoff_1
		alpha.Add(alpha, big.NewInt(1))
		return curve.Add(zX, zY, params.Gx, params.Gy)
	case zY.Cmp(cutoff_2) == -1: // y < cutoff_2
		alpha.Mul(alpha, big.NewInt(2))
		beta.Mul(beta, big.NewInt(2))
		return curve.Double(zX, zY)
	default:
		beta.Add(beta, big.NewInt(1))
		return curve.Add(zX, zY, pX, pY)
	}
}

/*
Square root algorithms
*/

//TODO: allow swapping out of different mixing functions
func PollardRho(curve *elliptic.CurveParams,
	pX, pY big.Int) *big.Int {

	alpha, beta := big.NewInt(1), big.NewInt(0)
	gamma, delta := big.NewInt(1), big.NewInt(0)
	zX, zY := curve.Gx, curve.Gy
	z2X, z2Y := ellipticMixingFunction(zX, zY, pX, pY, gamma, delta, curve)

	for i := 0; i < 1000; i++ {
		zX, zY = ellipticMixingFunction(zX, zY, pX, pY, alpha, beta, curve)
		z2X, z2Y = ellipticMixingFunction(z2X, z2Y, pX, pY, gamma, delta, curve)
		z2X, z2Y = ellipticMixingFunction(z2X, z2Y, pX, pY, gamma, delta, curve)
		if zX.Cmp(z2X) == 0 && zY.Cmp(z2Y) == 0 {
			fmt.Println("Found collision in %s iterations", i)

			num := big.NewInt(1).Sub(alpha, gamma)
			denom := big.NewInt(1).Sub(delta, beta)
			num.Mod(num, curve.N)
			denom.Mod(denom, curve.N)
			fmt.Println("Num:", num)
			fmt.Println("Denom:", denom)

			//The inverse exists only if gcf(|E(F_q)|, denom) == 1
			//TODO: consider the general case
			inverse := denom.ModInverse(denom, curve.N)
			fmt.Println("Inverse:", inverse)
			num.Mul(num, inverse)
			num.Mod(num, curve.N)

			return num
			break
		}
	}
}

// Not a priority
func BabyStepGiantStep(curve *elliptic.CurveParams,
	pX, pY big.Int) *big.Int {
	return new(big.Int)
}

/*
Embedding in F_q by Weil Pairing
*/

/*
Index calculus algorithms
time complexity: exp( (sqrt(2)+O(1)) * sqrt(ln(p)lnln(p)) )
*/
